package main

import (
	b "Relativity_c/Testing"
	"fmt"
)

func main() {
	var M, v float64
	fmt.Print("Value of Mass as INTEGER>")
	fmt.Scan(&M)
	fmt.Print("\nValue of Velocity as INTEGER>")
	fmt.Scan(&v)
	E := b.Relativity(M, v)
	fmt.Println("\nCalculated result of E =")
	fmt.Println(E)
}
