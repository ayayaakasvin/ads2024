package lab1

func IsBalanced (input string) bool {
	var anyChanges bool = true
	var mutableVersion []byte = []byte(input)
	
	for ; anyChanges ; {
		anyChanges = false
		i := 0
        for i < len(mutableVersion) - 1 {
            if i == 0 && mutableVersion[i] == mutableVersion[len(mutableVersion) - 1] {
                mutableVersion = append(mutableVersion[:i], mutableVersion[i + 1:len(mutableVersion) - 1]...)
                anyChanges = true
                continue
            }

            if mutableVersion[i] == mutableVersion[i + 1] {
                mutableVersion = append(mutableVersion[:i], mutableVersion[i + 2:]...)
                anyChanges = true
                continue
            }
            i++
        }
	}

	return len(mutableVersion) == 0
}