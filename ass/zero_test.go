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

func TestZero(t *testing.T) {
	fake := newFake()

	ass.Zero(fake, "")
	assertFalse(t, fake.failed)

	ass.Zero(fake, "baba")
	assertTrue(t, fake.failed)
}
