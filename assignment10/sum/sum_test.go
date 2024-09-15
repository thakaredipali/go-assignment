package main

import "testing"

func TestSumOfSquares(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{1, 2, 3}, 14},
		{[]int{4, 5, 6}, 77},
		{[]int{7, 8, 9}, 194},
	}

	for _, test := range tests {
		result := SumOfSquares(test.input)
		if result != test.expected {
			t.Errorf("For input %v, expected %d, got %d", test.input, test.expected, result)
		}
	}
}
