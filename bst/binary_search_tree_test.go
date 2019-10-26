package bst

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSearch(t *testing.T) {
	tree := &Node{Value: 5, Left: &Node{Value: 3, Right: &Node{Value: 4}}}

	if ok := tree.Search(4); !ok {
		t.Errorf("Expected 4 to be found.")
	}
}

func TestSearchMissingValue(t *testing.T) {
	tree := &Node{Value: 5, Left: &Node{Value: 3}}

	if ok := tree.Search(4); ok {
		t.Errorf("Found missing value '4' somehow")
	}
}

func TestInsertLeft(t *testing.T) {
	tree := &Node{Value: 5, Left: &Node{Value: 3, Right: &Node{Value: 4}}}

	tree.Insert(1)

	expected := &Node{Value: 5, Left: &Node{Value: 3, Left: &Node{Value: 1}, Right: &Node{Value: 4}}}

	if !reflect.DeepEqual(tree, expected) {
		t.Errorf("Expected %v, but got %v", expected, tree)
	}
}

func TestInsertRight(t *testing.T) {
	tree := &Node{Value: 5, Left: &Node{Value: 3, Right: &Node{Value: 4}}}

	tree.Insert(6)

	expected := &Node{Value: 5, Left: &Node{Value: 3, Right: &Node{Value: 4}}, Right: &Node{Value: 6}}

	if !reflect.DeepEqual(tree, expected) {
		t.Errorf("Expected %v, but got %v", expected, tree)
	}
}

func TestDeleteLeft(t *testing.T) {
	tree := &Node{Value: 5, Left: &Node{Value: 3, Right: &Node{Value: 4}}}
	res := tree.Delete(3)
	expected := &Node{Value: 5, Left: &Node{Value: 4}}

	fmt.Print(res)

	if !reflect.DeepEqual(tree, expected) {
		t.Errorf("Expected %v, but got %v", expected, tree)
	}
}

func TestDeleteRight(t *testing.T) {
	tree := &Node{Value: 5, Right: &Node{Value: 6, Right: &Node{Value: 7}}}
	tree.Delete(6)
	expected := &Node{Value: 5, Right: &Node{Value: 7}}

	if !reflect.DeepEqual(tree, expected) {
		t.Errorf("Expected %v, but got %v", expected, tree)
	}
}
