/*
	This software is licensed under the MIT License.
	The full license agreement can be found in the LICENSE file.
*/

package ass

func Panics(t Test, action func(), msg ...any) (res Result) {
	res = failure(t)

	defer func() {
		err := recover()
		if err != nil {
			res = success(t)
			return
		}

		t.Helper()
		printMsg(msg...)
		t.Errorf("expected action to result in a panic but got none")
	}()

	action()

	return
}
