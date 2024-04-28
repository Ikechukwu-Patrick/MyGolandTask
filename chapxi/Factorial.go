package main

import (
	"fmt"
)

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

func main() {
	var num int

	fmt.Print("Enter a non-negative integer: ")
	fmt.Scan(&num)

	if num < 0 {
		fmt.Println("Please enter a non-negative integer.")
		return
	}

	result := factorial(num)
	fmt.Printf("The factorial of %d! is = %d\n", num, result)
}
