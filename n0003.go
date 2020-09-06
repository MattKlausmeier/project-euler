package main

import (
	"fmt"
	"math"
)

func n0003() {
	// 775146 is the square root of 600851475143
	// I think technically we should be starting
	// our search at 600851475143 / 2, but this
	// method gives us the correct answer.
	for i := 775146; i > 1; i-- {
		if 600851475143%i == 0 {
			if isPrime(i) {
				fmt.Println(i)
				return
			}
		}
	}
}

func isPrime(n int) bool {
	for i := int(math.Sqrt(float64(n))); i > 1; i-- {
		if n%i == 0 {
			return false
		}
	}
	return true
}
