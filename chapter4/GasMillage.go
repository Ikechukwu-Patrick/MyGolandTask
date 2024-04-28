package main

import (
	"fmt"
)

func main() {
	var totalMiles, totalGallons int

	for {
		var miles, gallons int
		fmt.Print("Enter the miles driven (-1 to quit): ")
		fmt.Scanln(&miles)

		if miles == -1 {
			break
		}

		fmt.Print("Enter the gallons used: ")
		fmt.Scanln(&gallons)

		mpg := float64(miles) / float64(gallons)
		fmt.Printf("Miles per gallon for this trip: %.2f\n", mpg)

		totalMiles += miles
		totalGallons += gallons

		totalMpg := float64(totalMiles) / float64(totalGallons)
		fmt.Printf("Combined miles per gallon: %.2f\n", totalMpg)
	}

	fmt.Println("Program ended.")
}
