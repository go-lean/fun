/*
	Copyright (c) 2023 go-lean

	This software is licensed under the MIT License.
	The full license agreement can be found in the LICENSE file.
*/

package ass

func True[T ~bool](t Test, value T, msg ...any) Result {
	if value {
		return success(t)
	}

	t.Helper()
	printMsg(msg...)

	t.Errorf("expected 'true' but got 'false'", value)

	return failure(t)
}

func False[T ~bool](t Test, value T, msg ...any) Result {
	if !value {
		return success(t)
	}

	t.Helper()
	printMsg(msg...)

	t.Errorf("expected 'false' but got 'true'", value)

	return failure(t)
}
