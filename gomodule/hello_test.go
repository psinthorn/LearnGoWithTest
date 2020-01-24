package hello

import "testing"

func TestHello(t *testing.T) {
	want := "Hello Golang"
	if got := Hello(); got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
