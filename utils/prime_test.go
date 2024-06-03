package utils

import (
	"testing"
)

func TestPrime(t *testing.T) {
	primes := GetPrime(100)
	if len(primes) != 25 {
		t.Fatalf("len(primes) != 25, primes: %v", primes)
	}
}
