package main

import "fmt"

func n0009() {
	for a := 0; a < 1000; a++ {
		for b := a + 1; b < 1000; b++ {
			for c := b + 1; c < 1000; c++ {
				if a*a+b*b == c*c {
					if a+b+c == 1000 {
						fmt.Println(a * b * c)
					}
				}
			}
		}
	}
}
