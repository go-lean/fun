/*
	This software is licensed under the MIT License.
	The full license agreement can be found in the LICENSE file.
*/

package ass_test

import (
	"testing"

	"github.com/go-lean/fun/ass"
)

func TestEqual_True(t *testing.T) {
	fake := newFake()

	ass.Equal(fake, "baba", "baba")

	assertFalse(t, fake.failed)
	assertFalse(t, fake.terminated)
}

func TestEqual_False_NotTerminated(t *testing.T) {
	fake := newFake()

	ass.Equal(fake, "baba", "flag")

	assertTrue(t, fake.failed)
	assertFalse(t, fake.terminated)
}

func TestEqual_False_Terminated(t *testing.T) {
	fake := newFake()

	ass.Equal(fake, "baba", "flag").Required()

	assertTrue(t, fake.failed)
	assertTrue(t, fake.terminated)
}

func TestEqual_True_Required(t *testing.T) {
	fake := newFake()

	ass.Equal(fake, "baba", "baba").Required()

	assertFalse(t, fake.failed)
	assertFalse(t, fake.terminated)
}
