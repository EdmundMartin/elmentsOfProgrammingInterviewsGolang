package linked_lists

import (
	"testing"
)

func TestReverseSublist(t *testing.T) {
	res := SingleLinkedListFromSlice([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	res = ReverseSublist(res, 3, 8)
	if res.String() != "0, 1, 2, 8, 7, 6, 5, 4, 3, 9, 10, " {
		t.Errorf("did not get expected result")
	}
}