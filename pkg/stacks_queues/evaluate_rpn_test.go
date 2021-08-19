package stacks_queues

import "testing"

func TestEvaluateRPN(t *testing.T) {
	input := "3,4,+,2,*,1,+"
	expected := 15
	result := EvaluateRPN(input)
	if result != expected {
		t.Errorf("got: %d, expected: %d", result, expected)
	}
}
