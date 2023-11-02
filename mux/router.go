/*
	This software is licensed under the MIT License.
	The full license agreement can be found in the LICENSE file.
*/

package mux

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

type (
	Router struct {
		getPaths     []Route
		postPaths    []Route
		putPaths     []Route
		patchPaths   []Route
		deletePaths  []Route
		tracePaths   []Route
		headPaths    []Route
		optionsPaths []Route

		NotFoundHandler http.HandlerFunc
	}

	Route struct {
		tokens    []string
		tokensLen int
		handler   routeHandler
	}

	routeHandler struct {
		handler http.HandlerFunc
		params  map[int]string
	}
)

const errUnknownMethodFmt = "unknown method: %q"

var (
	keyRouteParams = struct{}{}

	_notFoundHandlerDefault = func(w http.ResponseWriter, _ *http.Request) {

		w.WriteHeader(http.StatusNotFound)
		payload := http.StatusText(http.StatusNotFound)

		_, _ = w.Write([]byte(payload))
	}
)

func NewRouter() *Router {

	router := &Router{
		NotFoundHandler: _notFoundHandlerDefault,
	}

	return router
}

func (r *Router) Register(method, path string, handler http.HandlerFunc) {

	params := make(map[int]string)
	route := Route{
		handler: routeHandler{
			handler: handler,
			params:  params,
		},
	}

	if path == "/" {
		route.tokens = []string{"/"}
		route.tokensLen = 1
		r.registerRoute(method, route)

		return
	}

	path = strings.Trim(path, "/")
	route.tokens = strings.Split(path, "/")
	route.tokensLen = len(route.tokens)

	for i, token := range route.tokens {
		if !strings.HasPrefix(token, ":") {
			continue
		}

		params[i] = token[1:]
	}

	r.registerRoute(method, route)
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	var tokens []string
	if req.URL.Path == "/" {
		tokens = []string{"/"}
	} else {
		path := strings.Trim(req.URL.Path, "/")
		tokens = strings.Split(path, "/")
	}

	routeHandler, ok := r.matchHandler(req.Method, tokens)
	if !ok {
		r.NotFoundHandler(w, req)
		return
	}

	if len(routeHandler.params) > 0 {
		req = r.loadParams(req, tokens, routeHandler.params)
	}

	routeHandler.handler(w, req)
}

func SetParams(r *http.Request, vars map[string]string) *http.Request {

	paramsContext := context.WithValue(r.Context(), keyRouteParams, vars)
	return r.WithContext(paramsContext)
}

func ParamsFor(r *http.Request) map[string]string {

	params := r.Context().Value(keyRouteParams)
	if params == nil {
		return make(map[string]string)
	}

	return params.(map[string]string)
}

func (r *Router) registerRoute(method string, route Route) {

	switch method {
	case http.MethodGet:
		r.getPaths = append(r.getPaths, route)
	case http.MethodPost:
		r.postPaths = append(r.postPaths, route)
	case http.MethodPatch:
		r.patchPaths = append(r.patchPaths, route)
	case http.MethodPut:
		r.putPaths = append(r.putPaths, route)
	case http.MethodDelete:
		r.deletePaths = append(r.deletePaths, route)
	case http.MethodHead:
		r.headPaths = append(r.headPaths, route)
	case http.MethodTrace:
		r.tracePaths = append(r.tracePaths, route)
	case http.MethodOptions:
		r.optionsPaths = append(r.optionsPaths, route)
	default:
		panic(fmt.Sprintf(errUnknownMethodFmt, method))
	}
}

func (r *Router) loadParams(req *http.Request, tokens []string, params map[int]string) *http.Request {

	vars := make(map[string]string)
	for k, v := range params {
		vars[v] = tokens[k]
	}

	return SetParams(req, vars)
}

func (r *Router) matchHandler(method string, tokens []string) (*routeHandler, bool) {

	paths := r.getPathsFor(method)
	maxMatchCount := 0

	var handler *routeHandler

	for i := 0; i < len(paths); i++ {
		path := paths[i]

		if path.tokensLen != len(tokens) {
			continue
		}

		matchScore := evaluateHandler(path.tokens, tokens)

		if matchScore == path.tokensLen {
			return &path.handler, true
		}

		if matchScore > maxMatchCount {
			maxMatchCount = matchScore
			handler = &path.handler
		}
	}

	if handler == nil {
		return nil, false
	}

	return handler, true
}

func evaluateHandler(pathTokens, requestTokens []string) int {

	matchScore := 0
	for i, token := range pathTokens {
		if token == requestTokens[i] {
			matchScore++
			continue
		}

		if strings.HasPrefix(token, ":") {
			continue
		}

		return 0
	}

	return matchScore
}

func (r *Router) getPathsFor(method string) []Route {

	switch method {
	case http.MethodGet:
		return r.getPaths
	case http.MethodPost:
		return r.postPaths
	case http.MethodPatch:
		return r.patchPaths
	case http.MethodPut:
		return r.putPaths
	case http.MethodDelete:
		return r.deletePaths
	case http.MethodHead:
		return r.headPaths
	case http.MethodTrace:
		return r.tracePaths
	case http.MethodOptions:
		return r.optionsPaths
	}

	panic(fmt.Sprintf(errUnknownMethodFmt, method))
}
