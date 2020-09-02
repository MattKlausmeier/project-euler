package main

import "fmt"

func n0002() {
	first := 1
	second := 1
	sum := 0
	for second <= 4000000 {
		temp := second
		second += first
		first = temp
		if second%2 == 0 && second <= 4000000 {
			sum += second
		}
	}
	fmt.Println(sum)
}
