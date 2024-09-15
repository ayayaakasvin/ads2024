package lab1

import "testing"

type FStruct struct {
	input int
	expected int
}

func TestGetPrime(t *testing.T) {
	tests := []FStruct {
		{input: 5, expected: 11},
		{input: 1, expected: 2},
		{input: 2, expected: 3},
		{input: 16, expected: 53},
	}

	for _, tt := range tests {
		result := GetPrime(tt.input)
		if result != tt.expected {
			t.Errorf("GetPrime(%d) = %d, want %d", tt.input, result, tt.expected)
		}
	}
}