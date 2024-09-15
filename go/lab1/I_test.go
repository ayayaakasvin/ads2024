package lab1

import "testing"

type IStruct struct {
	input string
	expected string
}

func TestISolution (t *testing.T) {
	tests := []IStruct {
		{input: "KSKS", expected: "KATSURAGI"},
		{input: "SSKKK", expected: "SAKAYANAGI"},
	}

	for _, tt := range tests {
		result := ISolution(tt.input)
		
		if result != tt.expected {
			t.Errorf("ISOlution(%s) = %s, want %s", tt.input, result, tt.expected)
		}
	}
}