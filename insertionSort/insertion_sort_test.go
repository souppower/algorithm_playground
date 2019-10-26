package insertionSort

import (
	"testing"
)

func Test_insertion_sort(t *testing.T) {
	unsorted_list := []int{10, 2, 3, 1, 5, 6}

	result := insertion_sort(unsorted_list)
	expected := []int{1, 2, 3, 5, 6, 10}

	slices_equal := slice_equals(result, expected)

	if !slices_equal {
		t.Errorf("list was not sorted. expected: %v, but got %v", expected, result)
	}
}

func slice_equals(first []int, second []int) bool {
	for i, v := range first {
		if v != second[i] {
			return false
		}
	}

	return true
}
