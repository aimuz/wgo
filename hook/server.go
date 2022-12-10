package hook

import (
	"crypto/sha1" // #nosec
	"encoding/hex"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strings"
)

// GetTokenFunc ...
type GetTokenFunc func() (string, error)

// MessageHandle ...
type MessageHandle func(msg *Message) (Messager, error)

// Server implements http.Handler. It validates incoming WeChat Public Platform webhooks and
// then dispatches them to the appropriate plugins.
type Server struct {
	token  string
	appID  string
	aesKey string
	//messageHandle MessageHandle
}

// NewServer not implemented
func NewServer(appid, token, aesKey string) *Server {
	return &Server{
		appID:  appid,
		token:  token,
		aesKey: aesKey,
	}
}

// ServeHTTP implements an http.Handler that answers callback requests.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// this is the request for verification
	if r.Method == http.MethodGet {
		DefaultEchoHandle(s.token, w, r)
		return
	}

	//query := r.URL.Query()
	//openid := query.Get("openid")
	//
	//msg, encrypted, err := ValidateWebhook(r, s.token)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusBadRequest)
	//	return
	//}
	//
	//// DecryptMsg
	//if encrypted {
	//
	//}
	//
	//reply, err := s.messageHandle(msg)
	//if err != nil {
	//	return
	//}
	//
	//replyMsg := &Message{
	//	ToUserName:   msg.FromUserName,
	//	FromUserName: msg.ToUserName,
	//	CreateTime:   time.Now().Unix(),
	//}
	//switch t := reply.(type) {
	//case *TextMessage:
	//	replyMsg.MsgType = "text"
	//	replyMsg.Content = t.Content
	//case *ImageMessage:
	//case *VoiceMessage:
	//case *VideoMessage:
	//case *MusicMessage:
	//case *ArticleMessage:
	//}
	//
	//err = xml.NewEncoder(w).Encode(replyMsg)
	//if err != nil {
	//	return
	//}
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
func ValidateWebhook(r *http.Request, token string) (payload []byte, encrypted bool, err error) {
	fail := func(err error) ([]byte, bool, error) { return nil, false, err }

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
	msg := &Message{}
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
	return payload, encrypted, nil
}

// SHA1Sign Computing signatures using sha1
func SHA1Sign(args ...string) string {
	sort.Strings(args)

	h := sha1.New() // #nosec
	h.Write([]byte(strings.Join(args, "")))
	hashcode := hex.EncodeToString(h.Sum(nil))
	return hashcode
}
