package main

import (
	"testing"
)

func TestHelloGo(t *testing.T) {
	want := "Hello Go"
	
	got := hello()

	if want != got {
		t.Fatalf("Want: %s, got %s", want, got)
	}
}