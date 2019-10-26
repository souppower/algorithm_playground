package binarySearch

import "testing"

func Test_binary_search(t *testing.T) {
	searchValue := 3
	list := &List{Values: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}}
	position := list.Search(searchValue)

	if position != 2 {
		t.Errorf("expected 2, got %d", position)
	}
}

func Test_binary_search_missing_value(t *testing.T) {
	searchValue := 11
	list := &List{Values: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}}
	position := list.Search(searchValue)
	if position != -1 {
		t.Errorf("expected -1, got %d", position)
	}
}
