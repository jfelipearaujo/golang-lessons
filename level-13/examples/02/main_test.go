package main

import "testing"

type test struct {
	input    []int
	expected int
}

func Test_sum(t *testing.T) {
	tests := []test{
		{
			input:    []int{1, 2, 3, 4, 5},
			expected: 15,
		},
		{
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expected: 55,
		},
		{
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
			expected: 210,
		},
	}

	for _, test := range tests {
		result := sum(test.input...)

		if result != test.expected {
			t.Error("Expected:", test.expected, "Result:", result)
		}
	}
}

func Test_multiplication(t *testing.T) {
	tests := []test{
		{
			input:    []int{1, 2, 3, 4, 5},
			expected: 120,
		},
		{
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expected: 3628800,
		},
		{
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
			expected: 2432902008176640000,
		},
	}

	for _, test := range tests {
		result := multiplication(test.input...)

		if result != test.expected {
			t.Error("Expected:", test.expected, "Result:", result)
		}
	}
}
