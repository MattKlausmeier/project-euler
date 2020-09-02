package main

import (
	"fmt"
	"strconv"
)

func n0004() {
	var largest int
	for a := 999; a >= 0; a-- {
		for b := 999; b >= 0; b-- {
			if isPalindrome(strconv.Itoa(a*b)) && a*b > largest {
				largest = a * b
			}
		}
	}
	fmt.Println(largest)
}

func isPalindrome(str string) bool {
	var rev string
	for _, run := range str {
		rev = string(run) + rev
	}
	return str == rev
}
