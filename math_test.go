package main

import "testing"

func TestAdd(t *testing.T) {
	got := Add(100, 110)
	want := 210

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestMultiply(t *testing.T) {
	got := Multiply(45, 1)
	want := 450

	if want != got {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
