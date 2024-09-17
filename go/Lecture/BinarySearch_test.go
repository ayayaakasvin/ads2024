package lecture

import (
	"fmt"
	"testing"
)

type OutputCheck struct {
	input         []int
	target        int
	expectedInt   int
	expectedError error
}

func TestBinarySearchFound(t *testing.T) {
	testCases := []OutputCheck{
		{input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, target: 2, expectedInt: 1, expectedError: nil},
	}
	
	for _, tt := range testCases {
		var bss = BinarySearchConstructor(tt.input)

		result, err := bss.Search(tt.target)
		if result != tt.expectedInt || err != tt.expectedError {
			t.Errorf("Expected %d, %v : Got %d, %v", tt.expectedInt, tt.expectedError, result, err)
		}
	}
}

func TestBinarySearchNotFound(t *testing.T) {
	testCases := []OutputCheck{
		{input: []int{1, 2, 4, 5, 6, 7, 8, 9}, target: 3, expectedInt: -1, expectedError: fmt.Errorf("not found")},
	}
	
	for _, tt := range testCases {
		var bss = BinarySearchConstructor(tt.input)

		result, err := bss.Search(tt.target)
		if result != tt.expectedInt || err.Error() != tt.expectedError.Error() {
			t.Errorf("Expected %d, %v : Got %d, %v", tt.expectedInt, tt.expectedError, result, err)
		}
	}
}
