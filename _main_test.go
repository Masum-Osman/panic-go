package main

import "testing"

func TestMul(t *testing.T) {
	got := Multiply(100, 110)
	want := 110

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
