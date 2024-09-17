package lab2

func abs (number int) int {
	if number < 0 {
		return -number
	}
	return number
}

func NearestNumber (numbers []int, target int) int {
	var (
		finalIndex int
		difference int = abs(target - numbers[finalIndex])
 	)

	for index, value := range numbers {
		if target == value {
			return index
		}

		if abs(target - value) < difference {
			finalIndex = index
			difference = abs(target - value)
		}
	}

	return finalIndex
}