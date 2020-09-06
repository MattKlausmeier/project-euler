package main

import "fmt"

func n0012() {
	trinum := 0
	for n := 0; n < 10000000; n++ {
		trinum += n
		divisorcount := 0
		for x := 1; x <= trinum; x++ {
			if trinum%x == 0 {
				divisorcount++
			}
		}
		if divisorcount > 500 {
			fmt.Println(trinum)
			return
		}
	}
}
