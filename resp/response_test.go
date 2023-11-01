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

func TestNew_Getters(t *testing.T) {
	response := resp.New(http.StatusOK, "baba", "baba/plain")

	ass.Equal[string](t, "baba", response.Payload.(string), "wrong payload")
	ass.Equal[string](t, "baba/plain", response.Type, "wrong content type")
	ass.Equal[int](t, http.StatusOK, response.Code, "wrong status code")
}

func TestIsSuccessful(t *testing.T) {
	response := resp.New(200, "baba", "baba/plain")
	ass.True(t, response.IsSuccessful(), "response has to be successful")

	response = resp.New(399, "baba", "baba/plain")
	ass.True(t, response.IsSuccessful(), "response has to be successful")

	response = resp.New(199, "baba", "baba/plain")
	ass.False(t, response.IsSuccessful(), "response has to be unsuccessful")

	response = resp.New(400, "baba", "baba/plain")
	ass.False(t, response.IsSuccessful(), "response has to be unsuccessful")
}
