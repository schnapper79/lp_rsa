package main

import (
	"fmt"
	"math"
	"time"
)

func sieve_of_eratosthenes(max int) []bool {
	sieve := make([]bool, max+1)

	for i := 2; i <= max; i++ {
		sieve[i] = true
	}

	for i := 2; i <= int(math.Sqrt(float64(max)))+1; i += 1 {
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

func print_sieve(sieve []bool) {
	fmt.Println("Sieve:")
	for i := 1; i < len(sieve); i += 2 {
		if sieve[i] {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()
}

func sieve_to_primes(sieve []bool) []int {
	primes := make([]int, 0)
	for i := 1; i < len(sieve); i += 2 {
		if sieve[i] {
			primes = append(primes, i)
		}
	}
	return primes
}

func main() {
	var max int
	fmt.Printf("Max: ")
	fmt.Scan(&max)

	start := time.Now()
	sieve := sieve_of_eratosthenes(max)
	elapsed := time.Since(start)
	fmt.Printf("Elapsed: %f seconds\n", elapsed.Seconds())

	if max <= 1000 {
		print_sieve(sieve)

		primes := sieve_to_primes(sieve)
		fmt.Println(primes)
	}
}
