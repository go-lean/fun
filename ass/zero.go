/*
	Copyright (c) 2023 go-lean

	This software is licensed under the MIT License.
	The full license agreement can be found in the LICENSE file.
*/

package ass

func Zero[T comparable](t Test, value T, msg ...any) Result {
	var zero T
	if value == zero {
		return success(t)
	}

	t.Helper()
	printMsg(msg...)

	t.Errorf("expected empty string got %#v", value)

	return failure(t)
}
