package main

import "testing"

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

// BET -> Benchmarks, Examples and Tests

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum(1, 2, 3, 4, 5)
	}
}

func BenchmarkMultiplication(b *testing.B) {
	for i := 0; i < b.N; i++ {
		multiplication(1, 2, 3, 4, 5)
	}
}
