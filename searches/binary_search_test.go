package main

import "testing"

func Test_binary_search(t *testing.T) {
	searchValue := 3
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	position := binary_search(list, searchValue)
	if position != 2 {
		t.Errorf("expected 2, got %d", position)
	}
}

func Test_binary_search_missing_value(t *testing.T) {
	searchValue := 11
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	position := binary_search(list, searchValue)
	if position != -1 {
		t.Errorf("expected -1, got %d", position)
	}
}
