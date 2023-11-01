/*
	This software is licensed under the MIT License.
	The full license agreement can be found in the LICENSE file.
*/

package resp_test

import (
	"net/http"
	"testing"

	"github.com/go-lean/fun/ass"
	"github.com/go-lean/fun/resp"
)

func TestOptionsWithContentType(t *testing.T) {
	opts := resp.WithContentType("flag/plain")

	response := resp.New(http.StatusOK, "baba", "baba/plain", opts)

	ass.Equal[string](t, "flag/plain", response.Type, "wrong content type")
}

func TestOptionsWithHeader(t *testing.T) {
	header := http.Header{
		http.CanonicalHeaderKey("x-baba-is-you"): []string{"baba", "is", "you"},
	}
	opts := resp.WithHeaders(header)

	response := resp.New(http.StatusOK, "success", "text/plain", opts)

	headerBaba := response.Header.Get("x-baba-is-you")
	ass.Equal[string](t, "baba", headerBaba, "wrong header value")
}
