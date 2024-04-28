package main

import "fmt"

func main() {
	fmt.Println("N\tN^2\tN^3\tN^4")

	for N := 1; N <= 5; N++ {
		N2 := N * N
		N3 := N * N * N
		N4 := N * N * N * N

		fmt.Printf("%d\t%d\t%d\t%d\n", N, N2, N3, N4)
	}
}
