package knapsack

import (
	"testing"
)

func TestKnapsack(t *testing.T) {
	weights := []int{10, 20, 30}
	values := []int{60, 100, 120}
	capacity := 50
	expected := 220

	res := knapsack(len(values), weights, values, capacity)

	if res != expected {
		t.Errorf("Expected %d, but got %d", expected, res)
	}
}
