package util

import "testing"

func TestSignature(t *testing.T) {
	// abc sig
	abc := "a9993e364706816aba3e25717850c26c9cd0d89d"
	if abc != Signature("a", "b", "c") {
		t.Error("test Signature Error")
	}
}
