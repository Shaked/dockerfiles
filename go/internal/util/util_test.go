package util

import (
	"testing"
)

func TestHello(t *testing.T) {
	expected := "Hello, world!"
	got := Hello()

	if got != expected {
		t.Errorf("Hello() = %s; want %s", got, expected)
	}
}
