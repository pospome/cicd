package main

import "testing"

func TestNewMessage(t *testing.T) {
	if NewMessage() != "Hello, World" {
		t.Errorf("error")
	}
}
