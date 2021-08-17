package linked_lists

import (
	"testing"
)

func TestRemoveKthNode(t *testing.T) {
	res := SingleLinkedListFromSlice([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	removed := RemoveKthNode(res, 5)
	if removed.String() != "0, 1, 2, 3, 4, 5, 7, 8, 9, 10, " {
		t.Errorf("got unexpected output")
	}
}
