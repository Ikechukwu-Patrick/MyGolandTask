package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var counter, _, largest int
	scanner := bufio.NewScanner(os.Stdin)

	for counter < 10 {
		fmt.Printf("Enter integer %d: ", counter+1)
		scanner.Scan()
		input := scanner.Text()

		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid integer.")
			continue
		}

		if counter == 0 || num > largest {
			largest = num
		}

		counter++
	}

	fmt.Printf("The largest integer entered is: %d\n", largest)
}
