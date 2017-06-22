package main

import (
)

func eratosthenes(max int64) ([]int64, error){
	primes := []int64{2}
	for n := int64(2) ; n < max ; n++ {
		prime := true
		for _, p := range primes {
			if n % p == 0 { //n = k p -> n not prime
				prime = false
				break
			}
		}
		if prime {
			primes = append(primes, n)
		}
	}
	return primes, nil
}
