package main

import (
	"fmt"
	"testing"
)

func Test_sum(t *testing.T) {
	expected := 15
	result := sum(1, 2, 3, 4, 5)

	if result != expected {
		(*t).Error("Expected:", expected, "Result:", result)
	}
}

func Test_multiplication(t *testing.T) {
	expected := 120
	result := multiplication(1, 2, 3, 4, 5)

	if result != expected {
		(*t).Error("Expected:", expected, "Result:", result)
	}
}

func ExampleSum() {
	fmt.Println(sum(1, 2, 3, 4, 5))
	// Output:
	// 15
}

func ExampleMultiplication() {
	fmt.Println(multiplication(1, 2, 3, 4, 5))
	// Output:
	// 120
}
