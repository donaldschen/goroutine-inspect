package main

import "testing"

func Test_Dedupe(t *testing.T) {
	d, err := load("samples/stack2.txt")
	if err != nil {
		t.Fatal(err)
	}

	d.Dedupe()
}
