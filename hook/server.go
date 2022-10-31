package hook

import (
	"net/http"
)

// Server implements http.Handler. It validates incoming WeChat Public Platform webhooks and
// then dispatches them to the appropriate plugins.
type Server struct {
}

// NewServer not implemented
func NewServer() *Server {
	return &Server{}
}

func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var eventType string
	b, err := s.demuxEvent(eventType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = b
}

func (s Server) demuxEvent(eventType string) ([]byte, error) {
	panic("not implemented")
}
