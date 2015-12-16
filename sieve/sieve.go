package sieve

/*Sieve finds all prime numbers less than limit using the Sieve or Eratosthenes*/
func Sieve(limit int) []int {
	sieve := make([]bool, limit)
	primes := []int{2}
	for i := 3; i < limit; i += 2 {
		if !sieve[i] {
			for p := i * i; p < limit; p += i {
				sieve[p] = true
			}
			primes = append(primes, i)
		}
	}
	return primes
}
