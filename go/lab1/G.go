package lab1

func GSolution(superTimes int) int {
	return GetPrime(GetPrime(superTimes))
}