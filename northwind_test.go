package main

import "testing"

func TestNorthWind(t *testing.T) {
	// TODO
	got := 5 + 5
    want := 10

    if got != want {
        t.Errorf("got %q, wanted %q", got, want)
    }

}