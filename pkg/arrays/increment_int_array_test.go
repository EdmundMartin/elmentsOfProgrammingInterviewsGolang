package arrays

import (
	"reflect"
	"testing"
)

func TestPlusOne(t *testing.T) {
	arr := PlusOne([]int{1, 2, 9})
	expected := []int{1, 3, 0}
	isEqual := reflect.DeepEqual(arr, expected)
	if !isEqual {
		t.Errorf("expected: %v, got: %v", expected, arr)
	}
}
