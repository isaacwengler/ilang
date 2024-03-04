package test

import "testing"

/**
 * Shared testing utils
 */

func assertPanic(t *testing.T, f func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	f()
}

func assertStringEquals(t *testing.T, actual string, expected string) {
	if actual != expected {
		t.Errorf("Expected: '%s' but got: '%s'", expected, actual)
	}
}
