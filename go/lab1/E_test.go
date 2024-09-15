package lab1

import "testing"

type EStruct struct {
	Nursik []int
	Boris []int
	expected string
}

func TestESolution(t *testing.T) {
	testCases := []EStruct {
		{Nursik: []int{1, 3, 5, 7, 9}, Boris: []int{2, 4, 6, 8, 0}, expected: "Nursik 5"},
		{Nursik: []int{0, 1, 2, 3, 4}, Boris: []int{9, 8, 7, 6, 5}, expected: "Boris 6"},
		{Nursik: []int{6, 7, 8, 9, 0}, Boris: []int{1, 2, 3, 4, 5}, expected: "Nursik 10"},
		{Nursik: []int{4, 3, 2, 1, 0}, Boris: []int{9, 8, 7, 6, 5}, expected: "Boris 5"},
		{Nursik: []int{2, 4, 6, 8, 0}, Boris: []int{1, 3, 5, 7, 9}, expected: "Nursik 5"},
		{Nursik: []int{9, 0, 1, 2, 3}, Boris: []int{8, 7, 6, 5, 4}, expected: "Boris 10"},
		{Nursik: []int{0, 1, 2, 3, 4}, Boris: []int{5, 6, 7, 8, 9}, expected: "Draw"},
	}
	
	for _, tt := range testCases {
		result := ESolution(tt.Nursik, tt.Boris)

		if result != tt.expected {
			t.Errorf("Expected %s, got %s", tt.expected, result)
		}
	}
}