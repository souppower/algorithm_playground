package insertionSort

import (
	"testing"
)

func Test_insertion_sort(t *testing.T) {
	list := &List{Values: []int{10, 2, 3, 1, 5, 6}}
	expected := &List{Values: []int{1, 2, 3, 5, 6, 10}}

	list.Sort()

	isEqualList := isEqualList(list, expected)

	if !isEqualList {
		t.Errorf("list was not sorted. expected: %v, but got %v", expected, list)
	}
}

func isEqualList(first, second *List) bool {
	for i, v := range (*first).Values {
		if v != (*second).Values[i] {
			return false
		}
	}

	return true
}
