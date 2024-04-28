package main

import "fmt"

func main() {
	var base int

	fmt.Print("Enter the length of the base of the triangle (1-10): ")
	fmt.Scanln(&base)

	if base < 1 || base > 10 {
		fmt.Println("Invalid input. Please enter a base length between 1 and 10.")
		return
	}

	for i := 1; i <= base; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
