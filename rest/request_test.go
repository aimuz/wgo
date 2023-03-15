package rest

import (
	"errors"
	"testing"
)

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
