package lab1

import "testing"

type CStruct struct {
	a, b string
	expected bool
}

func TestCSolutionTest(t *testing.T)  {
	tests := []CStruct {
		{a: "ab#c", b: "ad#c", expected: true},
		{a: "abc", b: "aaa#bbb####bcc#", expected: true},
		{a: "a#sn##d#kj#f", b: "l#df##n#vo", expected: false},
		{a: "Jokerge", b: "SratVechno", expected: false},
	}

	for _, tt := range tests {
		result := CSolutionTest(tt.a, tt.b)
		if result != tt.expected {
			t.Errorf("CSolutionTest(%s, %s) = %t; want %t", tt.a, tt.b, result, tt.expected)
		}
	}
}