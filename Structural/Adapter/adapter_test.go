package adapter

import (
	"testing"
)

const expect = "Adaptee's method"

func TestAdapter(t *testing.T) {
	adaptee := NewAdaptee()
	target := NewAdapter(adaptee)
	res := target.Request()
	if res != expect {
		t.Fatalf("expect %s, but get %s\n", expect, res)
	}
}