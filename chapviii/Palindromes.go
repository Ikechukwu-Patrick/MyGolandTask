package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(num int) bool {
	s := strconv.Itoa(num)
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var num int
	for {
		fmt.Print("Enter a five-digit integer: ")
		fmt.Scan(&num)

		if num < 10000 || num > 99999 {
			fmt.Println("Error: Please enter a five-digit integer.")
		} else {
			break
		}
	}

	if isPalindrome(num) {
		fmt.Println("The number is a palindrome.")
	} else {
		fmt.Println("The number is not a palindrome.")
	}
}
