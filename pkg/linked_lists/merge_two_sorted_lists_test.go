package linked_lists

import (
	"testing"
)

func TestMergeTwoSortedLists(t *testing.T) {
	first := SingleLinkedListFromSlice([]int{1, 2, 4, 5, 8, 11, 13, 17, 23, 34})
	second := SingleLinkedListFromSlice([]int{10, 14, 18, 35, 46})

	result := MergeTwoSortedLists(first, second)
	if result.String() != "1, 2, 4, 5, 8, 10, 11, 13, 14, 17, 18, 23, 34, 35, 46, " {
		t.Errorf("Did not get expected output")
	}
}