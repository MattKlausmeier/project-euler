package main

import "fmt"

func n0005() {
	for a := 1; a < 999999999; a++ {
		var found bool
		for b := 2; b <= 20; b++ {
			if a%b != 0 {
				found = true
				break
			}
		}
		if !found {
			fmt.Println(a)
			return
		}
	}
}
