package lab1

func cleanOperation(input *string) {
	runes := []rune(*input)
	hashtagCount := 0
	rewritePosition := len(runes) - 1

	for i := rewritePosition; i >= 0; i-- {
		if runes[i] == '#' {
			hashtagCount++
		} else {
			if hashtagCount > 0 {
				hashtagCount--
			} else {
				runes[rewritePosition] = runes[i]
				rewritePosition--
			}
		}
	}

	runes = runes[rewritePosition+1:]

	*input = string(runes)
}

func CSolutionTest (firstInput, secondInput string) bool {
	cleanOperation(&firstInput)
	cleanOperation(&secondInput)

	return firstInput == secondInput
}