package main

import (
	"testing"
)

func TestFragger(t *testing.T) {
	a := "Gladys"
	b := "Gladys"

	if a != b {
		t.Fatalf(`%q != %v`, a, b)
	}
}
