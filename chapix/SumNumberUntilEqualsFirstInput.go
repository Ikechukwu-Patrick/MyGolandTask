package main

import (
	"fmt"
)

func main() {
	var targetSum int
	fmt.Print("Enter a target sum: ")
	fmt.Scan(&targetSum)

	sum := 0
	for sum < targetSum {
		var num int
		fmt.Print("Enter an integer: ")
		fmt.Scan(&num)
		sum += num
	}

	if sum == targetSum {
		fmt.Println("The sum is equal to the target sum %d", targetSum)
	} else if sum != targetSum {
		fmt.Println("The sum is not equal to the target sum %d", targetSum)
	} else if sum < targetSum {
		fmt.Println("The sum is less than the target sum %d", targetSum)
	} else {
		fmt.Println("The sum is greater than the target sum %d", targetSum)
	}

}
