package main

import (
	"testing"
)

func TestCounter(t *testing.T) {
	counter := Counter{}
	if c := counter.Read(); c != 0 {
		t.Fatalf("counter should be 0, but was %d", c)
	}
	counter.Increment()
	if c := counter.Read(); c != 1 {
		t.Fatalf("counter should be 1, but was %d", c)
	}
	counter.Increment()
	if c := counter.ReadAndReset(); c != 2 {
		t.Fatalf("counter should have been 2, but was %d", c)
	}
	if c := counter.Read(); c != 0 {
		t.Fatalf("counter should be have been 0 after read/reset, but was %d", c)
	}
}