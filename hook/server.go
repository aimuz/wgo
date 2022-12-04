package hook

import (
	// #nosec
	"crypto/sha1"
	"encoding/hex"
	"encoding/xml"
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

	msgType, msg, _, ok := ValidateWebhook(w, r, token)
	if !ok {
		return
	}

	echostr := r.URL.Query().Get("echostr")
	if len(echostr) > 0 && r.Method == http.MethodGet {
		_, _ = w.Write([]byte(echostr))
		return
	}

	reply, err := s.handleMessage(appid, msgType, msg)
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

// ValidateWebhook ...
func ValidateWebhook(w http.ResponseWriter, r *http.Request, token string) (string, *Message, bool, bool) {
	query := r.URL.Query()
	signature := query.Get("signature")
	timestamp := query.Get("timestamp")
	nonce := query.Get("nonce")

	msg := Message{}
	decoder := xml.NewDecoder(r.Body)
	err := decoder.Decode(&msg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return "", nil, false, false
	}
	defer r.Body.Close() // nolint: errcheck

	encrypted := false
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
		http.Error(w, "signature mismatch", http.StatusBadRequest)
		return "", nil, encrypted, false
	}
	return msg.MsgType, &msg, encrypted, true
}
