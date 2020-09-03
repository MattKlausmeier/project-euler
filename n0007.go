package main

import "fmt"

func n0007() {
	primecount := 0
	lastprime := 0
	for a := 2; primecount < 10001; a++ {
		if isPrime(a) {
			lastprime = a
			primecount++
		}
	}
	fmt.Println(lastprime)
}
