package wgo

import (
	"net/http"
	"time"
)

// Reference: golang.org/x/oauth2
// More targeted changes compared to oauth2 to avoid introducing too many dependencies

// A TokenSource is anything that can return a token.
type TokenSource interface {
	// Token returns a token or an error.
	// Token must be safe for concurrent use by multiple goroutines.
	// The returned Token must not be modified.
	Token() (*Token, error)
}

// Token represents the credentials used to authorize
// the requests to access protected resources on the OAuth 2.0
// provider's backend.
//
// Most users of this package should not access fields of Token
// directly. They're exported mostly for use by related packages
// implementing derivative OAuth2 flows.
type Token struct {
	// AccessToken is the token that authorizes and authenticates
	// the requests.
	AccessToken string `json:"access_token"`

	// Expiry is the optional expiration time of the access token.
	//
	// If zero, TokenSource implementations will reuse the same
	// token forever and RefreshToken or equivalent
	// mechanisms for that TokenSource will not be used.
	Expiry time.Time `json:"expiry,omitempty"`
}

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
