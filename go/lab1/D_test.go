package lab1

import "testing"

type DStruct struct {
	input string
	expected bool
}

func TestIsBalanced (t *testing.T) {
	tests := []DStruct {
		{input: "sbaabsss", expected: true},
		{input: "sbabasss", expected: false},
		{input: "baab", expected: true},
		{input: "abpa", expected: false},
		{input: "abbbbbbba", expected: false},
	}

	for _, tt := range tests {
		result := IsBalanced(tt.input)

		if result != tt.expected {
			t.Errorf("IsBalanced(%s) = %t, want %t", tt.input, result, tt.expected)
		}
	}
}