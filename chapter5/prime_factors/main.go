package main

import (
	"fmt"
	"time"
)

func sieve_of_eratosthenes(max uint64) []bool {
	sieve := make([]bool, max+1)

	for i := uint64(2); i <= max; i++ {
		sieve[i] = true
	}

	for i := uint64(2); i*i <= max; i += 1 {
		if sieve[i] {
			for j := i * i; j <= max; j += i {
				if sieve[j] {
					sieve[j] = false
				}
			}
		}
	}

	return sieve
}

func sieve_to_primes(sieve []bool) []uint64 {
	primes := []uint64{2}
	for i := uint64(3); i < uint64(len(sieve)); i += 2 {
		if sieve[i] {
			primes = append(primes, i)
		}
	}
	return primes
}

func find_factors(n uint64) []uint64 {
	factors := make([]uint64, 0)

	for ; n%2 == 0; n /= 2 {
		factors = append(factors, 2)
	}

	factor := uint64(3)
	for ; factor*factor <= n; factor += 2 {
		for ; n%factor == 0; n /= factor {
			factors = append(factors, factor)
		}
	}

	if n > 1 {
		factors = append(factors, n)
	}

	return factors
}

func multiply_factors(factors []uint64) uint64 {
	product := uint64(1)
	for _, factor := range factors {
		product *= factor
	}
	return product
}

func fast_find_factors(n uint64) []uint64 {
	factors := make([]uint64, 0)

	for _, prime := range primes {
		if prime*prime > n {
			break
		}
		for ; n%prime == 0; n /= prime {
			factors = append(factors, prime)
		}

	}

	if n > 1 {
		factors = append(factors, n)
	}

	return factors
}

var primes []uint64

func main() {

	//fill primes
	primes = sieve_to_primes(sieve_of_eratosthenes(1600000000))

	var num uint64
	fmt.Printf("Number: ")
	fmt.Scan(&num)

	start := time.Now()
	factors := find_factors(num)
	elapsed := time.Since(start)
	fmt.Printf("Elapsed slow: %f seconds\n", elapsed.Seconds())

	start2 := time.Now()
	factors2 := fast_find_factors(num)
	elapsed2 := time.Since(start2)

	fmt.Printf("Elapsed fast: %f seconds\n", elapsed2.Seconds())

	fmt.Printf("slow Factors: %v\n", factors)
	fmt.Printf("fast Factors: %v\n", factors2)

	fmt.Printf("slow Product: %d\n\n", multiply_factors(factors))
	fmt.Printf("fast Product: %d\n\n", multiply_factors(factors2))
}
