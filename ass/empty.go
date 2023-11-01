/*
	This software is licensed under the MIT License.
	The full license agreement can be found in the LICENSE file.
*/

package ass

func EmptyString[T ~string](t Test, value T, msg ...any) Result {
	if value == "" {
		return success(t)
	}

	t.Helper()
	printMsg(msg...)

	t.Errorf("expected empty string got %#v", value)

	return failure(t)
}
