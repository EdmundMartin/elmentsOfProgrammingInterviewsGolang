package stacks_queues

import "testing"

func TestIsWellFormed(t *testing.T) {
	wellFormed := "{}()[]"
	if IsWellFormed(wellFormed) != true {
		t.Errorf("expected to be well formed")
	}
	poorlyFormed := "{)"
	if IsWellFormed(poorlyFormed) != false {
		t.Errorf("expected to be poorly formed")
	}
	badlyFormed := "{(})"
	if IsWellFormed(badlyFormed) != false {
		t.Errorf("expected to be poorly formed")
	}
}