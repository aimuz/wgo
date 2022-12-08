package hook

import (
	// #nosec
	"crypto/sha1"
	"encoding/hex"
	"encoding/xml"
	"errors"
	"net/http"
	"sort"
	"strings"
	"time"
)

// GetTokenFunc ...
type GetTokenFunc func(appid string) (string, error)

// MessageHandle ...
type MessageHandle func(appid string, msgType string, msg *Message) (Messager, error)

// Server implements http.Handler. It validates incoming WeChat Public Platform webhooks and
// then dispatches them to the appropriate plugins.
type Server struct {
	getToken      GetTokenFunc
	handleMessage MessageHandle
}

// NewServer not implemented
func NewServer() *Server {
	return &Server{}
}

// ServeHTTP ...
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	appid := r.URL.Query().Get("appid")
	token, err := s.getToken(appid)
	if err != nil {
		// TODO
		return
	}

	msg, _, err := ValidateWebhook(r, token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	echostr := r.URL.Query().Get("echostr")
	if len(echostr) > 0 && msg.MsgType == "Validate" {
		_, _ = w.Write([]byte(echostr))
		return
	}

	reply, err := s.handleMessage(appid, msg.MsgType, msg)
	if err != nil {
		return
	}

	replyMsg := &Message{
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

	//if encrypted {
	// TODO Support Encrypt Msg
	//}

	err = xml.NewEncoder(w).Encode(replyMsg)
	if err != nil {
		return
	}
}

// SignatureError signature mismatch
var SignatureError = errors.New("signature mismatch")

// ValidateWebhook ...
func ValidateWebhook(r *http.Request, token string) (msg *Message, encrypted bool, err error) {
	query := r.URL.Query()
	signature := query.Get("signature")
	timestamp := query.Get("timestamp")
	nonce := query.Get("nonce")

	msg = &Message{
		MsgType: "Validate",
	}

	if r.Method == http.MethodPost {
		decoder := xml.NewDecoder(r.Body)
		err = decoder.Decode(msg)
		if err != nil {
			return
		}
		defer r.Body.Close() // nolint: errcheck
	}

	if len(msg.Encrypt) > 0 {
		encrypted = true
	}

	// TODO Support Encrypt Msg
	args := []string{timestamp, nonce, token, msg.Encrypt}
	sort.Strings(args)

	// #nosec
	h := sha1.New()
	h.Write([]byte(strings.Join(args, "")))
	hashcode := hex.EncodeToString(h.Sum(nil))
	if signature != hashcode {
		return msg, encrypted, SignatureError
	}
	return msg, encrypted, nil
}
