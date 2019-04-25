package is_prime_number

func IsPrimeNumber(num int) (int, bool) {
	var remainder int

	if num == 1 {
		return num, false
	}

	if num == 2 {
		return num, true
	}

	if num > 2 {
		for i := 2; i <= num-1; i++ {
			remainder = num % i
			if remainder == 0 {
				return num, false
			}
		}
	}

	return num, true
}
