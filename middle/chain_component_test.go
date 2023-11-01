/*
	This software is licensed under the MIT License.
	The full license agreement can be found in the LICENSE file.
*/

package middle_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-lean/fun/ass"
	"github.com/go-lean/fun/middle"
	"github.com/go-lean/fun/resp"
)

func TestComponent_NewComponent(t *testing.T) {
	firstStep := middle.Step(func(r *http.Request, next middle.Handler) resp.Result {
		response := next(r)
		response.Payload = fmt.Sprintf("first%v", response.Payload)

		return response
	})

	secondStep := middle.Step(func(r *http.Request, next middle.Handler) resp.Result {
		response := next(r)
		response.Payload = fmt.Sprintf("second%v", response.Payload)

		return response
	})

	c := middle.NewChainComponent(firstStep, secondStep)

	handler := c.Chain().Build(func(r *http.Request) resp.Result {
		return resp.New(http.StatusOK, "handler", "text/plain")
	})

	r := httptest.NewRequest("", "/", nil)

	response := handler(r)

	ass.Equal[string](t, "firstsecondhandler", response.Payload.(string))
}

func TestComponent_Use(t *testing.T) {
	firstStep := middle.Step(func(r *http.Request, next middle.Handler) resp.Result {
		response := next(r)
		response.Payload = fmt.Sprintf("first%v", response.Payload)

		return response
	})

	secondStep := middle.Step(func(r *http.Request, next middle.Handler) resp.Result {
		response := next(r)
		response.Payload = fmt.Sprintf("second%v", response.Payload)

		return response
	})

	component := middle.NewChainComponent(firstStep).Use(secondStep)

	handler := component.Chain().Build(func(r *http.Request) resp.Result {
		return resp.New(http.StatusOK, "handler", "text/plain")
	})

	r := httptest.NewRequest("", "/", nil)

	response := handler(r)

	ass.Equal[string](t, "firstsecondhandler", response.Payload.(string))
}

func TestComponent_Clone(t *testing.T) {
	firstStep := middle.Step(func(r *http.Request, next middle.Handler) resp.Result {
		response := next(r)
		response.Payload = fmt.Sprintf("first%v", response.Payload)

		return response
	})

	secondStep := middle.Step(func(r *http.Request, next middle.Handler) resp.Result {
		response := next(r)
		response.Payload = fmt.Sprintf("second%v", response.Payload)

		return response
	})

	component := middle.NewChainComponent(firstStep, secondStep)
	cloned := component.CloneChain()

	component.Use(func(r *http.Request, next middle.Handler) resp.Result {
		return resp.New(200, "must not be called", "text/plain")
	})

	handler := cloned.Build(func(r *http.Request) resp.Result {
		return resp.New(http.StatusOK, "handler", "text/plain")
	})

	r := httptest.NewRequest("", "/", nil)

	response := handler(r)

	ass.Equal[string](t, "firstsecondhandler", response.Payload.(string))
}

func TestComponent_UseChain(t *testing.T) {
	firstStep := middle.Step(func(r *http.Request, next middle.Handler) resp.Result {
		response := next(r)
		response.Payload = fmt.Sprintf("first%v", response.Payload)

		return response
	})

	secondStep := middle.Step(func(r *http.Request, next middle.Handler) resp.Result {
		response := next(r)
		response.Payload = fmt.Sprintf("second%v", response.Payload)

		return response
	})

	component := middle.NewChainComponent(firstStep)
	component.UseChain(middle.New(secondStep))

	handler := component.Chain().Build(func(r *http.Request) resp.Result {
		return resp.New(http.StatusOK, "handler", "text/plain")
	})

	r := httptest.NewRequest("", "/", nil)

	response := handler(r)

	ass.Equal[string](t, "firstsecondhandler", response.Payload.(string))
}
