// Copyright 2023 The WGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
			expectErr:  &Error{Code: 123, Msg: "Oops"},
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
