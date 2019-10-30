package fibonacci

import (
	"reflect"
	"testing"
)

func TestFibonacci(t *testing.T) {
	result := fibonacci(30)
	expected := 832040

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestFibonacci2(t *testing.T) {
	result := fibonacci2(30)
	expected := 832040

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func BenchmarkFibonacci(b *testing.B) {
	fibonacci(40)
}

func BenchmarkFibonacci2(b *testing.B) {
	fibonacci2(40)
}