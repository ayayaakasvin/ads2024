package lecture

import (
	"fmt"
)

type binarySearchSlice struct {
	elements []int
	beginning int
	end int
}

func InsertionSort (input []int)  {
	for i := 0; i < len(input); i++ {
		var insertionElem int = input[i]
		var insertionPosition int = i
		for j := insertionPosition - 1; j >= 0; j-- {
			if insertionElem < input[j] {
				input[j + 1] = insertionElem
				insertionPosition--
			}
		}

		input[insertionPosition] = insertionElem
	}
}

func (slice *binarySearchSlice) Search (target int) (int, error) {
	var (
		mid int = (slice.beginning + slice.end) / 2
		value int = slice.elements[mid]
	)

	if value == target {
		return mid, nil
	} else if slice.beginning == slice.end {
		return -1, fmt.Errorf("not found")
	}

	if value > target {
		slice.end = mid - 1
	} else {
		slice.beginning = mid + 1
	}

	return slice.Search(target)
}

func BinarySearchConstructor (input []int) *binarySearchSlice {
	InsertionSort(input)
	var result *binarySearchSlice = &binarySearchSlice{
		elements: input,
		end:  len(input) - 1,
		beginning: 0,
	}

	return result
}