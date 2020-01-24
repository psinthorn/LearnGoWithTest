package hello

import "testing"

func TestAddDepend(t *testing.T) {
	want := "Hello, world."

	if got := AddDepend(); got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
