package main

import "fmt"

func n0010() {
	var total int
	for a := 2; a < 2000000; a++ {
		if isPrime(a) {
			total += a
		}
	}
	fmt.Println(total)
}
