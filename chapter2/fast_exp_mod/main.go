package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func fast_exp(num, pow int) int {
	result := 1
	for pow > 0 {
		if pow%2 == 1 {
			result *= num
		}
		pow /= 2
		num *= num
	}
	return result
}

func fast_exp_mod(num, pow, mod int) int {
	result := 1
	// pow to binary with factors
	m := num % mod
	for ; pow > 0; pow /= 2 {
		if pow%2 == 1 {
			result *= m
			result %=mod
		}
		m = (m * m) % mod
	}
	return result % mod
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Calculator for fast exponentiation and fast exponentiation with modulo")
	fmt.Println("----------------------------------------------------------------------")
	for {
		var NUM, POW, MOD int
		var err error
		fmt.Printf("NUM= ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)
		NUM, err = strconv.Atoi(text)
		if err != nil {
			fmt.Printf("\nPlease enter only positive integer numbers or 0 to leave programm!\n")
			continue
		}
		if NUM < 1 {
			break
		}

		fmt.Printf("POW= ")
		text, _ = reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)
		POW, err = strconv.Atoi(text)
		if err != nil {
			fmt.Printf("\nPlease enter only positive integer numbers or 0 to leave programm!\n")
			continue
		}
		if POW < 1 {
			break
		}

		fmt.Printf("MOD= ")
		text, _ = reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)
		MOD, err = strconv.Atoi(text)
		if err != nil {
			fmt.Printf("\nPlease enter only positive integer numbers or 0 to leave programm!\n")
			continue
		}
		if MOD < 1 {
			break
		}

		testPOW := int(math.Pow(float64(NUM), float64(POW)))
		fmt.Printf("\nFast exponentiation: %d^%d=%d  (with math.Pow()=%d)\n", NUM, POW, fast_exp(NUM, POW), testPOW)
		testPOWMOD := int(math.Pow(float64(NUM), float64(POW))) % MOD
		fmt.Printf("Fast exponentiation with modulo: %d^%d mod %d=%d (with math.Pow()%%=%d)\n\n", NUM, POW, MOD, fast_exp_mod(NUM, POW, MOD), testPOWMOD)

	}
}
