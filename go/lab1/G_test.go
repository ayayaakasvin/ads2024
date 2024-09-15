package lab1

import "testing"

type GStruct struct {
	input int
	expected int
}

func TestGSolution (t *testing.T) {
	tests := []GStruct {
		{input: 1, expected: 3},
		{input: 2, expected: 5},
		{input: 3, expected: 11},
		{input: 5, expected: 31},
	}

	for _, tt := range tests {
		result := GSolution(tt.input)
		
		if result != tt.expected {
			t.Errorf("GSOlution(%d) = %d, want %d", tt.input, result, tt.expected)
		}
	}
}