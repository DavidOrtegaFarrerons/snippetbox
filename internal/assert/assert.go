package assert

import "testing"

func Equal[T comparable](t *testing.T, got, want T) {
	//This is used so if an error is shown, the displayed file and line will
	//The one that called this function, not the line were Errorf is in this file
	t.Helper()

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
