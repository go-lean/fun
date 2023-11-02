/*
	This software is licensed under the MIT License.
	The full license agreement can be found in the LICENSE file.
*/

package mux_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-lean/fun/ass"
	"github.com/go-lean/fun/mux"
)

type route struct {
	path    string
	method  string
	handler http.HandlerFunc
}

var (
	dudHandler    http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("wrong")) }
	targetHandler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("baba")) }
)

func TestRouter_Match(t *testing.T) {

	tc := []struct {
		name        string
		routes      []route
		path        string
		shouldMatch bool
	}{
		{
			"single root path - should match",
			[]route{
				{
					path:    "/",
					method:  http.MethodGet,
					handler: targetHandler,
				},
			},
			"/",
			true,
		},
		{
			"root path after other - should match",
			[]route{
				{
					path:    "/other",
					method:  http.MethodGet,
					handler: dudHandler,
				},
				{
					path:    "/",
					method:  http.MethodGet,
					handler: targetHandler,
				},
			},
			"/",
			true,
		},
		{
			"root path before other - should match",
			[]route{
				{
					path:    "/",
					method:  http.MethodGet,
					handler: targetHandler,
				},
				{
					path:    "/other",
					method:  http.MethodGet,
					handler: dudHandler,
				},
			},
			"/",
			true,
		},
		{
			"no paths - should not match",
			[]route{},
			"/baba",
			false,
		},
		{
			"second path different len - should match",
			[]route{
				{
					path:    "/flag",
					method:  http.MethodGet,
					handler: dudHandler,
				},
				{
					path:    "/flag/baba",
					method:  http.MethodGet,
					handler: targetHandler,
				},
			},
			"/flag/baba",
			true,
		},
		{
			"after path with parameter - should match",
			[]route{
				{
					path:    "/flag/:id",
					method:  http.MethodGet,
					handler: dudHandler,
				},
				{
					path:    "/flag/baba",
					method:  http.MethodGet,
					handler: targetHandler,
				},
			},
			"/flag/baba",
			true,
		},
		{
			"path with parameter - should match",
			[]route{
				{
					path:    "/flag/:id",
					method:  http.MethodGet,
					handler: targetHandler,
				},
				{
					path:    "/flag/win",
					method:  http.MethodGet,
					handler: dudHandler,
				},
			},
			"/flag/baba",
			true,
		},
		{
			"after path with multiple parameters - should match",
			[]route{
				{
					path:    "/flag/:id/then/:anotherId",
					method:  http.MethodGet,
					handler: dudHandler,
				},
				{
					path:    "/flag/win/then/winAgain",
					method:  http.MethodGet,
					handler: targetHandler,
				},
			},
			"/flag/win/then/winAgain",
			true,
		},
		{
			"match first parameter and last exact - should match",
			[]route{
				{
					path:    "/flag/:id/then/:anotherId",
					method:  http.MethodGet,
					handler: targetHandler,
				},
				{
					path:    "/flag/win/then/winAgain",
					method:  http.MethodGet,
					handler: dudHandler,
				},
			},
			"/flag/baba/then/winAgain",
			true,
		},
		{
			"match parameter but not match exact - should not match",
			[]route{
				{
					path:    "/flag/:id/then",
					method:  http.MethodGet,
					handler: targetHandler,
				},
				{
					path:    "/flag/win/then",
					method:  http.MethodGet,
					handler: dudHandler,
				},
			},
			"/flag/baba/melt",
			false,
		},
	}

	methods := []string{
		http.MethodGet,
		http.MethodPost,
		http.MethodPatch,
		http.MethodPut,
		http.MethodDelete,
		http.MethodHead,
		http.MethodTrace,
		http.MethodOptions,
	}

	for _, c := range tc {
		for _, method := range methods {
			t.Run(c.name+"_"+method, func(t *testing.T) {
				router := mux.NewRouter()

				for _, route := range c.routes {
					router.Register(method, route.path, route.handler)
				}

				w := httptest.NewRecorder()
				r := httptest.NewRequest(method, c.path, nil)

				router.ServeHTTP(w, r)

				if c.shouldMatch {
					ass.Equal(t, http.StatusOK, w.Code, "wrong status code")
					ass.Equal(t, "baba", w.Body.String())
					return
				}

				ass.Equal(t, http.StatusNotFound, w.Code, "wrong status code")
				ass.Equal(t, http.StatusText(http.StatusNotFound), w.Body.String())
			})
		}
	}
}

func TestHandle_InvalidMethod_Panics(t *testing.T) {

	router := mux.NewRouter()
	action := func() {
		router.Register("baba", "/", func(w http.ResponseWriter, r *http.Request) {})
	}

	ass.Panics(t, action, "failed to panic on invalid menthod handle")
}

func TestRouter_ServeHTTP_InvalidMethod_Panics(t *testing.T) {

	router := mux.NewRouter()
	action := func() {
		w := httptest.NewRecorder()
		r := httptest.NewRequest("baba", "/", nil)

		router.ServeHTTP(w, r)
	}

	ass.Panics(t, action, "failed to panic on invalid method serve")
}

func TestRouter_ServeHTTP_RouteParams(t *testing.T) {

	router := mux.NewRouter()
	called := false

	handler := func(w http.ResponseWriter, r *http.Request) {
		params := mux.ParamsFor(r)

		id := params["id"]
		name := params["name"]

		ass.Equal[string](t, "baba", id, "wrong id")
		ass.Equal[string](t, "is-you", name, "wrong name")

		w.WriteHeader(http.StatusOK)
		called = true
	}

	router.Register(http.MethodGet, "/some/:id/and/some/:name", handler)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/some/baba/and/some/is-you", nil)

	router.ServeHTTP(w, r)

	ass.True(t, called, "handler was not called")
}

func TestRouteParams_NotSet_Default(t *testing.T) {

	r := httptest.NewRequest("", "/", nil)

	params := mux.ParamsFor(r)
	nonExistant := params["baba"]

	ass.EmptyString(t, nonExistant, "unexpected parameter in route params")
}

func TestRouter_RouteWithQueryParams(t *testing.T) {

	router := mux.NewRouter()

	targetHandler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusTeapot)
		name := r.URL.Query().Get("name")

		_, _ = w.Write([]byte(name))
	}

	router.Register(http.MethodGet, "/some/path", targetHandler)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/some/path?name=baba", nil)

	router.ServeHTTP(w, r)

	ass.Equal(t, http.StatusTeapot, w.Code, "wrong status code")
	ass.Equal(t, "baba", w.Body.String(), "wrong response payload")
}

func TestRouter_RouteWithIDAndQueryParams(t *testing.T) {

	router := mux.NewRouter()

	targetHandler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusTeapot)
		name := r.URL.Query().Get("name")

		_, _ = w.Write([]byte(name))
	}

	router.Register(http.MethodGet, "/some/path/with/:id", targetHandler)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/some/path/with/someId?name=baba", nil)

	router.ServeHTTP(w, r)

	ass.Equal(t, http.StatusTeapot, w.Code, "wrong status code")
	ass.Equal(t, "baba", w.Body.String(), "wrong response payload")
}

func TestRouter_RootWithQueryParams(t *testing.T) {

	router := mux.NewRouter()

	targetHandler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusTeapot)
		name := r.URL.Query().Get("name")

		_, _ = w.Write([]byte(name))
	}

	router.Register(http.MethodGet, "/", targetHandler)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/?name=baba", nil)

	router.ServeHTTP(w, r)

	ass.Equal(t, http.StatusTeapot, w.Code, "wrong status code")
	ass.Equal(t, "baba", w.Body.String(), "wrong response payload")
}
