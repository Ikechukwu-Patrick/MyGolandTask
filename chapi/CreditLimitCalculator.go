package main

import "fmt"

func main() {
	var accountNumber, beginningBalance, totalCharges, totalCredits, creditLimit int

	fmt.Print("Enter account number (-1 to quit): ")
	fmt.Scanln(&accountNumber)

	for accountNumber != -1 {
		fmt.Print("Enter beginning balance: ")
		fmt.Scanln(&beginningBalance)

		fmt.Print("Enter total charges: ")
		fmt.Scanln(&totalCharges)

		fmt.Print("Enter total credits: ")
		fmt.Scanln(&totalCredits)

		fmt.Print("Enter credit limit: ")
		fmt.Scanln(&creditLimit)

		newBalance := beginningBalance + totalCharges - totalCredits
		fmt.Printf("New balance: %d\n", newBalance)

		if newBalance > creditLimit {
			fmt.Println("Credit limit exceeded")
		}

		fmt.Print("Enter account number (-1 to quit): ")
		fmt.Scanln(&accountNumber)
	}

}
