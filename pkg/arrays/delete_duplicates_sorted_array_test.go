package arrays

import (
	"reflect"
	"testing"
)

func TestDeleteDuplicates(t *testing.T) {
	input := []int{1, 1, 1, 3, 4, 4, 5, 6, 7, 7, 8}
	res := DeleteDuplicates(input)
	expected := []int{1, 3, 4, 5, 6, 7, 8}
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("got: %v, expected: %v", res, expected)
	}
}

func TestDeleteDuplicatesInPlace(t *testing.T) {
	input := []int{1, 1, 1, 3, 4, 4, 5, 6, 7, 7, 8}
	res := DeleteDuplicatesInPlace(input)
	expected := []int{1, 3, 4, 5, 6, 7, 8}
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("got: %v, expected: %v", res, expected)
	}
}
