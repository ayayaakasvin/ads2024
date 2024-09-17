package lab2

import "testing"

type AStruct struct {
	numbers []int
	target int
	expected int
}

func TestNearestNumber(t *testing.T)  {
	tests := []AStruct {
		{[]int{7, 8, -10, 4, 2, -1}, 5, 3},
		{[]int{1, 2, 3}, -10, 0},
		{[]int{1, 1, 1, 1, 1}, 1, 0},
		{[]int{1, 2, 90, 32, 2, 2}, 10, 1},
	}

	for _, tt := range tests {
		result := NearestNumber(tt.numbers, tt.target)
		if result != tt.expected {
			t.Errorf("NearestNumber(%v, %d) = %d; want %d", tt.numbers, tt.target, result, tt.expected)
		}
	}
}