package utils

func GetPrime[T Int](n T) []T {

	not_prime := make([]bool, n)
	primes := make([]T, 0)

	i := T(2)
	for i < n {
		if !not_prime[i] {
			primes = append(primes, i)
		}

		for _, j := range primes {
			if i*(j) >= n {
				break
			}
			not_prime[i*j] = true
			if i%(j) == 0 {
				break
			}
		}
		i++
	}
	return primes
}
