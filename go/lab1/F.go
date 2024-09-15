package lab1

func isPrime (numero int) (bool, int) {
	if numero <= 1 {
		return false, -1
	}
	if numero <= 3 {
		return true, numero
	}
	if numero % 2 == 0 || numero % 3 == 0 {
		return false, -1
	}

	for i := 5; i * i <= numero; i += 6 {
		if numero % i == 0 || numero % (i + 2) == 0 {
			return false, -1
		}
	}

	return true, numero
}

func GetPrime(times int) int {
	var counter, lastPrime int

	for i := 2; i <= 1000 && counter < times; i++ {
		if isTrue, prime := isPrime(i); isTrue {
			counter++
			lastPrime = prime
		}
	}

	return lastPrime
}