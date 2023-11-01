/*
	This software is licensed under the MIT License.
	The full license agreement can be found in the LICENSE file.
*/

package ass_test

import (
	"testing"

	"github.com/go-lean/fun/ass"
)

func TestPanic(t *testing.T) {
	fake := newFake()

	ass.Panics(fake, func() { panic("kaboom") })
	assertFalse(t, fake.failed)

	ass.Panics(fake, func() {})
	assertTrue(t, fake.failed)
}
