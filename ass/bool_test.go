/*
	Copyright (c) 2023 go-lean

	This software is licensed under the MIT License.
	The full license agreement can be found in the LICENSE file.
*/

package ass_test

import (
	"testing"

	"github.com/go-lean/fun/ass"
)

func TestTrue(t *testing.T) {
	fake := newFake()

	ass.True(fake, true)

	assertFalse(t, fake.failed)

	ass.True(fake, false)
	assertTrue(t, fake.failed)
}

func TestFalse(t *testing.T) {
	fake := newFake()

	ass.False(fake, false)

	assertFalse(t, fake.failed)

	ass.False(fake, true)
	assertTrue(t, fake.failed)
}
