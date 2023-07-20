package main

import (
	"fmt"
	"math/rand"
	"time"
)

func gcd(a, b int) int {
	// fix negative numbers according to task
	if a < 0 {
		a *= -1
	}
	if b < 0 {
		b *= -1
	}

	// first special case: both numbers are same
	if a == b {
		return a
	}

	// second special case (most common when algorithm is done)
	if b == 0 {
		return a
	}
	if a == 0 {
		return b
	}

	// do the math
	r := a % b

	// and finally run gcd recursively
	return gcd(b, r)
}

func lcm(a, b int) int {
	g := gcd(a, b)
	return (a / g) * (b)
}

func fast_exp(num, pow int) int {
	n := int(num)
	result := int(1)
	for pow > 0 {
		if pow%2 == 1 {
			result *= n
		}
		pow /= 2
		n *= n
	}
	return result
}

func fast_exp_mod(num, pow, mod int) int {
	result := int(1)
	// pow to binary with factors
	m := int(num % mod)
	for ; pow > 0; pow /= 2 {
		if pow%2 == 1 {
			result *= m
			result = result % mod
		}
		m = (m * m) % mod
	}
	rs := result % (mod)
	return int(rs)
}

// Return a pseudo random number in the range [min, max).
func rand_range(min int, max int) int {
	return min + random.Intn(max-min)
}

func is_probable_prime(p int) (bool, int, int) {
	for i := 0; i < num_tests; i++ {
		n := rand_range(2, p/100)
		if r := fast_exp_mod(n, p-1, p); r != 1 {
			return false, n, fast_exp_mod(n, p-1, p)
		}
	}
	return true, 0, 0
}

func find_prime(min, max int) int {
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

func totient(p, q int) int {
	return lcm((p - 1), (q - 1))
}

func random_exponent(phi int) int {
	for {
		e := rand_range(2, phi)
		if gcd(e, phi) == 1 {
			return e
		}
	}
}

func inverse_mod(a, n int) int {
	t := 0
	newt := 1
	r := n
	newr := a
	for newr != 0 {
		quotient := r / newr
		t, newt = newt, t-quotient*newt
		r, newr = newr, r-quotient*newr
	}
	if r > 1 {
		return -1
	}
	if t < 0 {
		t += n
	}
	return t
}

var random = rand.New(rand.NewSource(time.Now().UnixNano())) // Initialize with a changing seed
const num_tests = 20

func main() {
	p := find_prime(10000, 50000)
	q := find_prime(10000, 50000)

	n := p * q
	phi := totient(p, q)
	e := random_exponent(phi)
	d := inverse_mod(e, phi)

	fmt.Println("*** Public ***")
	fmt.Println("public key modulus:", n)
	fmt.Println("public key exponent:", e)
	fmt.Println()
	fmt.Println("*** Private ***")
	fmt.Printf("primes: %d, %d\n", p, q)
	fmt.Printf("phi(n)): %d\n", phi)
	fmt.Printf("d: %d\n", d)

	// test
	var m int
	for {
		fmt.Print("Enter a message to encrypt (less than ", n-1, "): ")
		fmt.Scanf("%d", &m)
		if m == 1 {
			break
		}
		if m >= n-1 {
			fmt.Println("Message must be less than", n-1)
			continue
		}
		cr := fast_exp_mod(m, e, n)
		pl := fast_exp_mod(cr, d, n)
		fmt.Println("Encrypted message:", cr)
		fmt.Println("Decrypted message:", pl)
		fmt.Println()
	}
}
