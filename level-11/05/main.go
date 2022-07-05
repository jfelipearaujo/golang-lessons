package main

import "testing"

func TestAverage(t *testing.T) {
	v := Average([]float64{1, 2, 3, 4, 5})

	if v != 1.5 {
		t.Errorf("Expected 1.5, got %v", v)
	}
}

func Average(f []float64) float64 {
	var sum float64

	for _, v := range f {
		sum += v
	}

	return sum / float64(len(f))
}
