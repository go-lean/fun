/*
	This software is licensed under the MIT License.
	The full license agreement can be found in the LICENSE file.
*/

package ass

import (
	"fmt"
)

type Test interface {
	Errorf(format string, args ...interface{})
	FailNow()
	Helper()
}

type Result struct {
	passing bool
	test    Test
}

func (r Result) Required() {
	if r.passing {
		return
	}

	r.test.FailNow()
}

func failure(t Test) Result {
	return Result{
		passing: false,
		test:    t,
	}
}

func success(t Test) Result {
	return Result{
		passing: true,
		test:    t,
	}
}

func printMsg(msg ...any) {
	fmt.Print("\n\t")
	fmt.Print(msg...)
	fmt.Printf("\n\n")
}
