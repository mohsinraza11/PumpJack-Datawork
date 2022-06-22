package main

import (
	"fmt"
	"math"
)

func main() {
	trig()
}

func trig() {
	var chce int
	var num float64

	for i := 1; i < 2; i++ {
		fmt.Print("Chose from the following\n1)Sin\n2)Cos\n3)Tan\n4)Exit\n\n")
		fmt.Scan(&chce)
		if (chce < 1) || (chce > 4) {
			fmt.Print("Invalid value. Try again....   \n")
			i = i - 1
		} else {
			i = i + 3
		}
	}

	if chce == 1 {
		fmt.Print("Enter the number\n")
		fmt.Scan(&num)
		fmt.Print("Sin:   ", math.Sin(num))

	} else if chce == 2 {
		fmt.Print("Enter the number\n")
		fmt.Scan(&num)
		fmt.Print("Cos:   ", math.Cos(num))

	} else if chce == 3 {
		fmt.Print("Enter the number\n")
		fmt.Scan(&num)
		fmt.Print("Tan:   ", math.Tan(num))

	} else {
		fmt.Print("Good-bye!")
	}

}
