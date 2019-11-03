package quickSort

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	list := &List{Values: []int{2, 1, 5, 4, 3}}
	expected := &List{Values: []int{1, 2, 3, 4, 5}}

	sorted := list.Sort()

	if !reflect.DeepEqual(sorted, expected) {
		t.Errorf("Expected %v, but got %v", expected, sorted)
	}
}
