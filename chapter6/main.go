package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func fast_exp(num, pow uint64) uint64 {
	result := uint64(1)
	for pow > 0 {
		if pow%2 == 1 {
			result *= num
		}
		pow /= 2
		num *= num
	}
	return result
}

func fast_exp_mod(num, pow, mod uint64) uint64 {
	if pow == 0 {
		return 1 % mod
	} else if pow == 1 {
		return num % mod
	}

	result := uint64(1)
	// pow to binary with factors
	m := num % mod
	for ; pow > 0; pow /= 2 {
		if pow%2 == 1 {
			result *= m
		}
		m = (m * m) % mod
	}
	return result % mod
}

// Return a pseudo random number in the range [min, max).
func rand_range(min uint64, max uint64) uint64 {
	return  min + random.(max-min)
}

func is_probable_prime(p uint64) (bool, uint64, uint64) {
	for i := 0; i < num_tests; i++ {
		n := rand_range(2, p-1)
		if fast_exp_mod(n, p-1, p) != 1 {

			return false, n, fast_exp_mod(n, p-1, p)
		}
	}
	return true, 0, 0
}

func find_prime(min, max uint64) uint64 {
	for {
		p := rand_range(min, max+1)
		if p%2 == 0 {
			p += 1
		}
		if t, _, _ := is_probable_prime(p); t {
			return p
		}
	}
}

func test_known_values() {
	primes := []uint64{
		10009, 11113, 11699, 12809, 14149,
		15643, 17107, 17881, 19301, 19793,
	}
	composites := []uint64{
		10323, 11397, 12212, 13503, 14599,
		16113, 17547, 17549, 18893, 19999,
	}
	fmt.Println("Testing known values")
	fmt.Println("Primes")
	for _, p := range primes {
		if t, n, r := is_probable_prime(p); t {
			fmt.Printf("%d prime\n", p)
		} else {
			fmt.Printf("%d composite (%d^p-1 mod p=%d)\n", p, n, r)
		}
	}
	fmt.Println("\nComposites")
	for _, c := range composites {
		if t, n, r := is_probable_prime(c); t {
			fmt.Printf("%d prime\n", c)
		} else {
			fmt.Printf("%d composite (%d^p-1 mod p=%d)\n", c, n, r)
		}
	}
}

var random = rand.New(rand.NewSource(time.Now().UnixNano())) // Initialize with a changing seed
const num_tests = 20

func main() {
	// print probability of false positive for num_tests tests in percent
	fmt.Printf("Probability of false positive: %f\n", 1.0/(math.Pow(2, float64(num_tests))))

	// test_known_values()
	test_known_values()

	var num uint64
	for {
		fmt.Printf("\ndigits: ")
		fmt.Scan(&num)

		max := fast_exp(10, num)
		start := time.Now()
		p := find_prime(max/10, max)
		elapsed := time.Since(start)
		fmt.Printf("Prime: %d in %f seconds\n", p, elapsed.Seconds())
	}
}
