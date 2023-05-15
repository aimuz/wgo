package rest

import (
	"errors"
	"net/http"
	"net/url"
	"testing"
)

func TestRequest_Method(t *testing.T) {
	t.Run(http.MethodPost, func(t *testing.T) {
		r := (&Request{base: &url.URL{}}).Post()
		if r.verb != http.MethodPost {
			t.Fatalf("Post expected %s, got %s", http.MethodPost, r.verb)
		}
	})

	t.Run(http.MethodPut, func(t *testing.T) {
		r := (&Request{base: &url.URL{}}).Put()
		if r.verb != http.MethodPut {
			t.Fatalf("Put expected %s, got %s", http.MethodPut, r.verb)
		}
	})

	t.Run(http.MethodPatch, func(t *testing.T) {
		r := (&Request{base: &url.URL{}}).Patch()
		if r.verb != http.MethodPatch {
			t.Fatalf("Patch expected %s, got %s", http.MethodPatch, r.verb)
		}
	})

	t.Run(http.MethodGet, func(t *testing.T) {
		r := (&Request{base: &url.URL{}}).Get()
		if r.verb != http.MethodGet {
			t.Fatalf("Get expected %s, got %s", http.MethodGet, r.verb)
		}
	})

	t.Run(http.MethodDelete, func(t *testing.T) {
		r := (&Request{base: &url.URL{}}).Delete()
		if r.verb != http.MethodDelete {
			t.Fatalf("Delete expected %s, got %s", http.MethodDelete, r.verb)
		}
	})
}

func TestRequest_AbsPath(t *testing.T) {
	t.Run("Preserves Trailing Slash", func(t *testing.T) {
		r := (&Request{base: &url.URL{}}).AbsPath("/foo/")
		if s := r.URL().String(); s != "/foo/" {
			t.Errorf("trailing slash should be preserved: %s", s)
		}
	})

	t.Run("Path Joins", func(t *testing.T) {
		r := (&Request{base: &url.URL{}}).AbsPath("foo/bar", "baz")
		if s := r.URL().String(); s != "foo/bar/baz" {
			t.Errorf("trailing slash should be preserved: %s", s)
		}
	})

	t.Run("err not nil", func(t *testing.T) {
		r := (&Request{base: &url.URL{}, err: errors.New("foo")}).AbsPath("/foo/")
		if s := r.URL().String(); s != "" {
			t.Errorf("trailing slash should be preserved: %s", s)
		}
	})
}

func TestResult_Into(t *testing.T) {
	// Test success case
	successBody := []byte(`{"name": "Alice", "age": 30}`)
	successResult := Result{body: successBody}
	var successVal struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	err := successResult.Into(&successVal)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if successVal.Name != "Alice" || successVal.Age != 30 {
		t.Errorf("Unexpected value: %#v", successVal)
	}

	// Test error case
	errorBody := []byte(`{"errcode": 10001, "errmsg": "Unknown error"}`)
	errorResult := Result{body: errorBody}
	var errorVal struct {
		ErrCode int    `json:"errcode"`
		ErrMsg  string `json:"errmsg"`
	}
	err = errorResult.Into(&errorVal)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if errorVal.ErrCode != 10001 || errorVal.ErrMsg != "Unknown error" {
		t.Errorf("Unexpected value: %#v", errorVal)
	}

	// Test error case with Result.err
	someErr := errors.New("some error")
	errorResultWithErr := Result{body: errorBody, err: someErr}
	err = errorResultWithErr.Into(&errorVal)
	if err != someErr {
		t.Errorf("Unexpected error: %v", err)
	}
}
