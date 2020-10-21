package main

import (
	"testing"
)

func TestFun(t *testing.T) {
	secret := "1123"
	guess := "0111"
	out := getHint(secret, guess)
	want := "1A1B"
	if out != want {
		t.Errorf("got %v, want %v", out, want)
	}
}
