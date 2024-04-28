package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	citizens := make([]struct {
		name     string
		earnings float64
	}, 3)

	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; i < 3; i++ {
		fmt.Printf("Enter name for Citizen %d: ", i+1)
		scanner.Scan()
		citizens[i].name = scanner.Text()

		fmt.Printf("Enter earnings for %s: ", citizens[i].name)
		scanner.Scan()
		earningsStr := scanner.Text()
		earnings, err := strconv.ParseFloat(earningsStr, 64)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid number.")
			return
		}
		citizens[i].earnings = earnings
	}

	for _, citizen := range citizens {
		tax := calculateTax(citizen.earnings)
		fmt.Printf("%s's total tax: %.2f USD\n", citizen.name, tax)
	}
}

func calculateTax(earnings float64) float64 {
	var tax float64

	if earnings <= 30000 {
		tax = earnings * 0.15
	} else {
		tax = 30000*0.15 + (earnings-30000)*0.20
	}

	return tax
}
