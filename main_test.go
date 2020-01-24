package main

import "testing"

func TestMain(t *testing.T) {
	want := "This is main test"

	if got := Main(); got != want {
		t.Errorf("Got %q want %q", got, want)
	}
}
