package testutil

import "testing"

func CheckUnexpectedError(t *testing.T, err error) {
	if err == nil {
		return
	}
	t.Fatalf("failed, but no error was expected: %v", err)
}
