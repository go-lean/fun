/*
	Copyright (c) 2023 go-lean

	This software is licensed under the MIT License.
	The full license agreement can be found in the LICENSE file.
*/

package resp

import (
	"net/http"
)

type (
	Result struct {
		Payload any
		Code    int
		Type    string
		Header  http.Header
	}
)

func New(statusCode int, data any, contentType string, opts ...Opts) Result {
	resp := Result{
		Code:    statusCode,
		Payload: data,
		Type:    contentType,
	}

	for _, o := range opts {
		o(&resp)
	}

	return resp
}

func (r Result) IsSuccessful() bool {
	return r.Code > 199 && r.Code < 400
}
