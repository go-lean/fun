/*
	This software is licensed under the MIT License.
	The full license agreement can be found in the LICENSE file.
*/

package ass

func Equal[T comparable](t Test, expected, actual T, msg ...any) Result {
	if expected == actual {
		return success(t)
	}

	t.Helper()
	printMsg(msg...)

	t.Errorf("expected %#v got %#v", expected, actual)

	return failure(t)
}
