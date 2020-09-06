package main

import "fmt"
import "os"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Enter the number of the problem you'd like to solve as the first command line argument.")
		return
	}
	switch os.Args[1] {
	case "1":
		n0001()
	case "2":
		n0002()
	case "3":
		n0003()
	case "4":
		n0004()
	case "5":
		n0005()
	case "6":
		n0006()
	case "7":
		n0007()
	case "8":
		n0008()
	case "9":
		n0009()
	case "10":
		n0010()
	case "11":
		n0011()
	case "13":
		n0013()
	default:
		fmt.Println("I can't solve that one... yet.")
	}
}
