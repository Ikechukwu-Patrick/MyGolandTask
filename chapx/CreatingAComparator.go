package main

import (
	"fmt"
)

func compareNumbers(num1, num2 int) int {
	if num1 == num2 {
		return 0
	} else if num1 > num2 {
		return 1
	} else if num2 > num1 {
		return -1
	}
	return 00
}

func main() {
	var num1, num2 int

	fmt.Print("Enter the first number: ")
	fmt.Scan(&num1)

	fmt.Print("Enter the second number: ")
	fmt.Scan(&num2)

	result := compareNumbers(num1, num2)

	switch result {
	case 0:
		fmt.Println("0")
	case 1:
		fmt.Println("1")
	case -1:
		fmt.Println("-1")
	}
}
