package prime

/*Nth calculates the nth prime number, only valid for positive non zero numbers*/
func Nth(n int) (int, bool) {
	if n <= 0 {
		return 0, false
	}
	primes := []int{0, 2, 3, 5, 7, 11, 13, 17}
NextCandidate:
	for candidate := primes[len(primes)-1] + 2; len(primes) <= n; candidate += 2 {
		for d := 1; primes[d]*primes[d] <= candidate; d++ {
			if candidate%primes[d] == 0 {
				continue NextCandidate
			}
		}
		primes = append(primes, candidate)
	}
	return primes[n], true
}
