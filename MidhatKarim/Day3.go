package main

import (
	"fmt"
)

func main() {
	var x1, x2, y1, y2 float64
	fmt.Println("x1: ")
	fmt.Scanln(&x1)
	fmt.Println("x2: ")
	fmt.Scanln(&x2)
	fmt.Println("y1: ")
	fmt.Scanln(&y1)
	fmt.Println("y2: ")
	fmt.Scanln(&y2)

	m := (y2 - y1) / (x2 - x1)
	fmt.Println("Slope = ", m)
}
