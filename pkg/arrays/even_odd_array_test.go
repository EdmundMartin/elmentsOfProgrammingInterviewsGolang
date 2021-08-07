package arrays

import (
	"reflect"
	"testing"
)

func TestEvenOdd(t *testing.T) {
	inputArray := []int{1, 20, 12, 77, 21, 24, 38, 11}
	EvenOdd(inputArray)
	isEqual := reflect.DeepEqual(inputArray, []int{38, 20, 12, 24, 21, 77, 11, 1})
	if !isEqual {
		t.Errorf("expected evenOdd array, got: %v", inputArray)
	}
}
