package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"
)

// Request ...
type Request struct {
	// base is the root URL for all invocations of the client
	base *url.URL

	hc *http.Client
	// generic components accessible via method setters
	verb       string
	pathPrefix string
	params     url.Values
	headers    http.Header

	// output
	err  error
	body io.Reader
}

// NewRequest returns a new request object.
func NewRequest(base *url.URL, hc *http.Client) *Request {
	return &Request{
		base: base,
		hc:   hc,
	}
}

// Post sends a POST request
func (r *Request) Post() *Request {
	return r.Verb(http.MethodPost)
}

// Put sends a PUT request
func (r *Request) Put() *Request {
	return r.Verb(http.MethodPut)
}

// Patch sends a PATCH request
func (r *Request) Patch() *Request {
	return r.Verb(http.MethodPatch)
}

// Get sends a GET request
func (r *Request) Get() *Request {
	return r.Verb(http.MethodGet)
}

// Delete sends a DELETE request
func (r *Request) Delete() *Request {
	return r.Verb(http.MethodDelete)
}

// Verb sets the verb this request will use.
func (r *Request) Verb(verb string) *Request {
	r.verb = verb
	return r
}

// AbsPath overwrites an existing path with the segments provided. Trailing slashes are preserved
// when a single segment is passed.
func (r *Request) AbsPath(segments ...string) *Request {
	if r.err != nil {
		return r
	}
	r.pathPrefix = path.Join(r.base.Path, path.Join(segments...))
	if len(segments) == 1 && (len(r.base.Path) > 1 || len(segments[0]) > 1) && strings.HasSuffix(segments[0], "/") {
		// preserve any trailing slashes for legacy behavior
		r.pathPrefix += "/"
	}
	return r
}

// RequestURI overwrites existing path and parameters with the value of the provided server relative
// URI.
func (r *Request) RequestURI(uri string) *Request {
	if r.err != nil {
		return r
	}
	locator, err := url.Parse(uri)
	if err != nil {
		r.err = err
		return r
	}
	r.pathPrefix = locator.Path
	if len(locator.Query()) > 0 {
		if r.params == nil {
			r.params = make(url.Values)
		}
		for k, v := range locator.Query() {
			r.params[k] = v
		}
	}
	return r
}

// Param creates a query parameter with the given string value.
func (r *Request) Param(paramName, s string) *Request {
	if r.err != nil {
		return r
	}
	return r.setParam(paramName, s)
}

func (r *Request) setParam(paramName, value string) *Request {
	if r.params == nil {
		r.params = make(url.Values)
	}
	r.params[paramName] = append(r.params[paramName], value)
	return r
}

// SetHeader sets the request header
func (r *Request) SetHeader(key string, values ...string) *Request {
	if r.headers == nil {
		r.headers = http.Header{}
	}
	r.headers.Del(key)
	for _, value := range values {
		r.headers.Add(key, value)
	}
	return r
}

// Body ...
func (r *Request) Body(val interface{}) *Request {
	switch v := val.(type) {
	case string:
		r.body = strings.NewReader(v)
	case url.Values:
		r.body = strings.NewReader(v.Encode())
	case io.Reader:
		r.body = v
	default:
		var buf bytes.Buffer
		r.err = json.NewEncoder(&buf).Encode(val)
		r.body = &buf
	}
	return r
}

// URL returns the current working URL.
func (r *Request) URL() *url.URL {
	p := r.pathPrefix

	finalURL := &url.URL{}
	if r.base != nil {
		*finalURL = *r.base
	}
	finalURL.Path = p

	query := url.Values{}
	for key, values := range r.params {
		for _, value := range values {
			query.Add(key, value)
		}
	}

	finalURL.RawQuery = query.Encode()
	return finalURL
}

func (r *Request) newHTTPRequest(ctx context.Context) (*http.Request, error) {
	url := r.URL().String()
	req, err := http.NewRequest(r.verb, url, r.body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	req.Header = r.headers
	return req, nil
}

// Do sends an request and returns Result
func (r *Request) Do(ctx context.Context) Result {
	if r.err != nil {
		return Result{
			err: r.err,
		}
	}

	client := r.hc
	if client == nil {
		client = http.DefaultClient
	}

	req, err := r.newHTTPRequest(ctx)
	if err != nil {
		return Result{
			err: err,
		}
	}

	resp, err := client.Do(req) // nolint:bodyclose
	if err != nil {
		return Result{
			err: err,
		}
	}
	// closed
	defer readAndCloseResponseBody(resp)

	result := r.transformResponse(req, resp)
	return result
}

func readAndCloseResponseBody(resp *http.Response) {
	if resp == nil {
		return
	}

	// Ensure the response body is fully read and closed
	// before we reconnect, so that we reuse the same TCP
	// connection.
	const maxBodySlurpSize = 2 << 10
	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.ContentLength <= maxBodySlurpSize {
		_, _ = io.Copy(io.Discard, &io.LimitedReader{R: resp.Body, N: maxBodySlurpSize})
	}
}

func (r *Request) transformResponse(_ *http.Request, resp *http.Response) Result {
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return Result{err: err}
	}

	return Result{
		body:       b,
		statusCode: resp.StatusCode,
	}
}

// Result is the result of transforming a request
type Result struct {
	body       []byte
	err        error
	statusCode int
}

// Into ...
func (r Result) Into(val interface{}) error {
	if r.err != nil {
		return r.err
	}

	err := json.Unmarshal(r.body, val)
	return err
}
