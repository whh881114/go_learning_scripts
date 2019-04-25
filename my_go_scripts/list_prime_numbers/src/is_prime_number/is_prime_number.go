package is_prime_number

func IsPrimeNumber(ch chan int, num int) (int, bool) {
	var remainder int

	if num == 1 {
		ch <- 0
		return num, false
	}

	if num == 2 {
		ch <- num
		return num, true
	}

	if num > 2 {
		for i := 2; i <= num-1; i++ {
			remainder = num % i
			if remainder == 0 {
				ch <- 0
				return num, false
			}
		}
	}

	// fmt.Println(num)

	ch <- num
	return num, true
}
