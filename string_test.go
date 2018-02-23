package slice

import (
	"testing"
)

func TestHasString(t *testing.T) {
	a := []string{
		"asdf",
		"qwer",
	}

	if HasString(a, "qwerqw") {
		t.Error("expected false")
	}

	if !HasString(a, "asdf") {
		t.Error("expected true")
	}
}
