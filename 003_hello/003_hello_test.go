package main

import "testing"

func TestMain(t *testing.T) {
	s := GetHello()
	if s != "hello" {
		t.Error("FALL")
	}

}
