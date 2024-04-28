package main

import "fmt"

func main() {
	var salesAmount float64
	var totalSale float64
	var commissionRate float64 = 0.09
	var baseSalary float64 = 200

	fmt.Println("Enter sales amount for the week:")
	fmt.Scanln(&salesAmount)

	totalSale = baseSalary + (salesAmount * commissionRate)

	fmt.Printf("The sales person's total earnings for the week is $%.2f\n", totalSale)
}
