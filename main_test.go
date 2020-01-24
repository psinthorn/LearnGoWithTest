package main

import "testing"

func TestHello(t *testing.T) {
	want := "Hello go"

	if got := Hello(); got != want {
		t.Errorf("Got %q want %q", got, want)
	}
}
