package wgo

import (
	"net/http"
)

// RoundTripper is a RoundTripper that implements the http.RoundTripper interface
type RoundTripper struct {
	next        http.RoundTripper
	TokenSource TokenSource
}

// RoundTrip Add the AccessToken parameter to the request
func (r *RoundTripper) RoundTrip(request *http.Request) (*http.Response, error) {
	if r.next == nil {
		r.next = http.DefaultTransport
	}

	if r.TokenSource == nil {
		resp, err := r.next.RoundTrip(request)
		return resp, err
	}
	tok, err := r.TokenSource.Token()
	if err != nil {
		return nil, err
	}
	val := request.URL.Query()
	val.Set("access_token", tok.AccessToken)
	request.URL.RawQuery = val.Encode()
	resp, err := r.next.RoundTrip(request)
	return resp, err
}
