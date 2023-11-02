/*
	Copyright (c) 2023 go-lean

	This software is licensed under the MIT License.
	The full license agreement can be found in the LICENSE file.
*/

package ass_test

import "testing"

type fakeTest struct {
	failed     bool
	terminated bool
}

func newFake() *fakeTest {
	return &fakeTest{}
}

func (t *fakeTest) Errorf(fm string, args ...interface{}) {
	t.failed = true
}

func (t *fakeTest) Helper() {
	// stargaze
}

func (t *fakeTest) FailNow() {
	t.terminated = true
}

func assertTrue(t *testing.T, value bool) {
	if value {
		return
	}

	t.Helper()
	t.Errorf("value is not true")
}

func assertFalse(t *testing.T, value bool) {
	if value == false {
		return
	}

	t.Helper()
	t.Errorf("value is not false")
}
