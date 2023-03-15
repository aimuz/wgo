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
	"encoding/json"
	"fmt"
)

// Error ...
type Error struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

// Error implements the error interface
func (e Error) Error() string {
	return fmt.Sprintf("errCode: %d, errMsg: %s", e.ErrCode, e.ErrMsg)
}

// NewError ...
func NewError(errCode int, errMsg string) error {
	return &Error{
		ErrCode: errCode,
		ErrMsg:  errMsg,
	}
}

// JSONValidator 是一个结构体，它用于在解析 JSON 数据之前验证 JSON 是否符合要求。
//
//	主要的作用是检查 JSON 数据中是否包含名为 errcode 和 errmsg 的字段，
//	如果存在 errcode 字段，则返回一个包含 errcode 和 errmsg 的错误。
//	否则，JSONValidator 将解析 JSON 数据。
type JSONValidator struct {
	val interface{}
}

// UnmarshalJSON 是 JSONValidator 的方法，它用于在解析 JSON 数据之前验证 JSON 是否符合要求。
//
//	它首先从 JSON 数据中获取 errcode 和 errmsg 字段，并检查 errcode 是否大于 0。
//	如果 errcode 大于 0，则 UnmarshalJSON 方法会返回一个包含 errcode 和 errmsg 的错误。
//	如果 errcode 不大于 0，则 UnmarshalJSON 方法将使用 encoding/json 包的 Unmarshal 函数解析 JSON 数据到目标结构体或值中。
func (r *JSONValidator) UnmarshalJSON(body []byte) error {
	var e Error
	if err := json.Unmarshal(body, &e); err == nil {
		if e.ErrCode > 0 {
			return &e
		}
	}
	return json.Unmarshal(body, r.val)
}

// NewJSONValidator 是一个函数，它创建并返回一个 JSONValidator 的实例。
// 它接受一个参数 val，它是要解析 JSON 数据的目标结构体或值。
func NewJSONValidator(val interface{}) *JSONValidator {
	return &JSONValidator{val: val}
}
