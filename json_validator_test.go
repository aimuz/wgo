package wgo

import (
	"reflect"
	"testing"
)

func BenchmarkJSONValidator(b *testing.B) {
	type Example struct {
		Foo string `json:"foo"`
	}
	data := []byte(`{"foo": "bar"}`)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var example Example
		jv := NewJSONValidator(&example)
		_ = jv.UnmarshalJSON(data)
	}
}

func TestJSONValidator_UnmarshalJSON(t *testing.T) {
	type Example struct {
		Foo string `json:"foo"`
	}
	tests := []struct {
		name       string
		jsonData   []byte
		expectData *Example
		expectErr  error
	}{
		{
			name:       "No error",
			jsonData:   []byte(`{"foo": "bar"}`),
			expectData: &Example{Foo: "bar"},
			expectErr:  nil,
		},
		{
			name:       "With error",
			jsonData:   []byte(`{"errcode": 123, "errmsg": "Oops"}`),
			expectData: &Example{},
			expectErr:  &Error{ErrCode: 123, ErrMsg: "Oops"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var data Example
			jv := NewJSONValidator(&data)
			err := jv.UnmarshalJSON(tt.jsonData)
			if !reflect.DeepEqual(err, tt.expectErr) {
				t.Errorf("Expect error %v, but got %v", tt.expectErr, err)
			}

			if !reflect.DeepEqual(&data, tt.expectData) {
				t.Errorf("Expect data %v, but got %v", tt.expectData, &data)
			}
		})
	}
}
