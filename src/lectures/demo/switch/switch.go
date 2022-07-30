package main

import "fmt"

func price() int {
	return 1
}

const (
	Economy    = 0
	Business   = 1
	FirstClass = 2
)

func main() {

	switch price() {
	case 0:
		fmt.Println("Economy")
	case 1:
		fmt.Println("Business")
	case 2:
		fmt.Println("First")
	}
}
