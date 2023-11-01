/*
	This software is licensed under the MIT License.
	The full license agreement can be found in the LICENSE file.
*/

package resp

import "net/http"

type Opts func(r *Result)

var (
	WithHeaders = func(header http.Header) Opts {
		return func(r *Result) {
			r.Header = header
		}
	}

	WithContentType = func(contentType string) Opts {
		return func(r *Result) {
			r.Type = contentType
		}
	}
)
