package webhook

import (
	"context"
	"crypto/sha1" // #nosec
	"encoding/hex"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strings"
	"time"
)

// GetTokenFunc ...
type GetTokenFunc func() (string, error)

// MessageHandle ...
type MessageHandle func(ctx context.Context, msg *PlainMessage) (Messager, error)

// Server implements http.Handler. It validates incoming WeChat Public Platform webhooks and
// then dispatches them to the appropriate plugins.
type Server struct {
	token         string
	appID         string
	aesKey        string
	crypto        WXBizMsgCrypto
	messageHandle MessageHandle
}

// Option ...
type Option func(*Server)

// WithPlainMode ...
func WithPlainMode(token string) Option {
	return func(server *Server) {
		server.token = token
	}
}

// WithSafeMode ...
func WithSafeMode(token, encodingAESKey string) Option {
	return func(server *Server) {
		server.token = token
		server.aesKey = encodingAESKey
	}
}

// NewServer not implemented
func NewServer(appid string, opts ...Option) (*Server, error) {
	s := &Server{
		appID: appid,
	}
	for _, opt := range opts {
		opt(s)
	}
	return s, s.init()
}

func (s *Server) init() error {
	if len(s.aesKey) > 0 {
		c, err := NewWXBizMsgCrypto(s.aesKey, s.appID, s.token)
		if err != nil {
			return err
		}
		s.crypto = c
	}
	return nil
}

type ctxKey struct{}

// ServeHTTP implements an http.Handler that answers callback requests.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// this is the request for verification
	if r.Method == http.MethodGet {
		DefaultEchoHandle(s.token, w, r)
		return
	}

	encryptMsg, payload, encrypted, err := ValidateWebhook(r, s.token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var msg = encryptMsg.PlainMessage
	if encrypted {
		err = s.crypto.DecryptMessage(payload, &msg)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	query := r.URL.Query()

	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	ctx = context.WithValue(ctx, ctxKey{}, query)
	reply, err := s.messageHandle(ctx, &msg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	replyMsg := &PlainMessage{
		ToUserName:   msg.FromUserName,
		FromUserName: msg.ToUserName,
		CreateTime:   time.Now().Unix(),
	}
	switch t := reply.(type) {
	case *TextMessage:
		replyMsg.MsgType = "text"
		replyMsg.Content = t.Content
	case *ImageMessage:
	case *VoiceMessage:
	case *VideoMessage:
	case *MusicMessage:
	case *ArticleMessage:
	}

	replyPayload, err := xml.Marshal(replyMsg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if encrypted {
		replyPayload, err = s.crypto.EncryptMessage(replyPayload, time.Now().Unix(), "")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	_, _ = fmt.Fprint(w, replyPayload)
}

// DefaultEchoHandle This is the default Echo Handle, which will be used when WeChat Public sends an authentication request
func DefaultEchoHandle(token string, w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	signature := query.Get("signature")
	timestamp := query.Get("timestamp")
	nonce := query.Get("nonce")
	hashcode := SHA1Sign(timestamp, nonce, token)
	if signature != hashcode {
		http.Error(w, ErrSignatureMismatch.Error(), http.StatusBadRequest)
		return
	}
	echo := r.URL.Query().Get("echostr")
	_, _ = fmt.Fprint(w, echo)
}

// ErrSignatureMismatch signature mismatch
var ErrSignatureMismatch = errors.New("signature mismatch")

// ErrEmptyBody ...
var ErrEmptyBody = errors.New("empty body")

// ValidateWebhook ...
func ValidateWebhook(r *http.Request, token string) (msg *EncryptMessage, payload []byte, encrypted bool, err error) {
	fail := func(err error) (*EncryptMessage, []byte, bool, error) { return nil, nil, false, err }

	query := r.URL.Query()
	signature := query.Get("signature")
	timestamp := query.Get("timestamp")
	nonce := query.Get("nonce")

	defer r.Body.Close() // nolint: errcheck
	payload, err = io.ReadAll(r.Body)
	if err != nil {
		return fail(err)
	}
	if len(payload) == 0 {
		return fail(ErrEmptyBody)
	}
	msg = &EncryptMessage{}
	err = xml.Unmarshal(payload, msg)
	if err != nil {
		return fail(err)
	}

	if len(msg.Encrypt) > 0 {
		encrypted = true
		payload = []byte(msg.Encrypt)
	}

	hashcode := SHA1Sign(timestamp, nonce, token, msg.Encrypt)
	if signature != hashcode {
		return fail(ErrSignatureMismatch)
	}
	return msg, payload, encrypted, nil
}

// SHA1Sign Computing signatures using sha1
func SHA1Sign(args ...string) string {
	sort.Strings(args)

	h := sha1.New() // #nosec
	h.Write([]byte(strings.Join(args, "")))
	hashcode := hex.EncodeToString(h.Sum(nil))
	return hashcode
}
