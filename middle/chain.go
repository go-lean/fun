/*
	This software is licensed under the MIT License.
	The full license agreement can be found in the LICENSE file.
*/

package middle

import (
	"net/http"

	"github.com/go-lean/fun/resp"
)

type (
	Chain struct {
		steps []Step
	}

	Step    func(r *http.Request, next Handler) resp.Result
	Handler func(r *http.Request) resp.Result
)

func New(steps ...Step) *Chain {
	return &Chain{
		steps: steps,
	}
}

func (c *Chain) Add(steps ...Step) *Chain {
	c.steps = append(c.steps, steps...)
	return c
}

func (c *Chain) Merge(other *Chain) *Chain {
	c.steps = append(c.steps, other.steps...)
	return c
}

func (c *Chain) Clone() *Chain {
	return New(c.steps...)
}

func (c *Chain) Build(lastHandler Handler) Handler {

	for i := len(c.steps) - 1; i > -1; i-- {
		lastHandler = wrap(c.steps[i], lastHandler)
	}

	return lastHandler
}

func wrap(step Step, next Handler) Handler {

	return func(r *http.Request) resp.Result {
		return step(r, next)
	}
}
