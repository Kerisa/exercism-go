package sieve

import (
)

// Sieve uses the Sieve of Eratosthenes to find all the primes from 2 up to a given number
func Sieve(limit int) []int {
	var prime = []int{}
	for i := 2; i <= limit; i++ {
		isPrime := true
		for _,n := range prime {
			if i % n == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			prime = append(prime, i)
		}
	}
	return prime
}