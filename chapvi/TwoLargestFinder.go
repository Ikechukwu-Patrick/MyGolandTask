package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var numbers []int

	for i := 0; i < 10; i++ {
		var input string
		fmt.Print("Enter a number: ")
		fmt.Scanln(&input)

		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid number.")
			i--
			continue
		}

		numbers = append(numbers, num)
	}

	sort.Ints(numbers)

	largest := numbers[len(numbers)-1]
	secondLargest := numbers[len(numbers)-2]

	fmt.Printf("The largest number is: %d\n", largest)
	fmt.Printf("The second largest number is: %d\n", secondLargest)
}
