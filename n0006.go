package main

import "fmt"

func n0006() {
	sumofthesquares := 0
	squareofthesums := 0
	for a := 1; a <= 100; a++ {
		sumofthesquares += a * a
	}
	for a := 1; a <= 100; a++ {
		squareofthesums += a
	}
	squareofthesums = squareofthesums * squareofthesums
	fmt.Println(squareofthesums - sumofthesquares)
}
