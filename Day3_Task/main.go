package main

import (
	"fmt"
	"math"
)

func CheckPrime() {
	var num int
	fmt.Print("Enter Number To check Prime:")
	fmt.Scanln(&num)
	isPrime := true
	if num == 0 || num == 1 {
		fmt.Printf(" %d is not a  prime number\n", num)
	} else {
		for i := 2; i <= num/2; i++ {
			if num%i == 0 {
				fmt.Printf(" %d is not a  prime number\n", num)
				isPrime = false
				break
			}
		}
		if isPrime == true {
			fmt.Printf(" %d is a prime number\n", num)
		}
	}
}

func quad() {
	var a, b, c, root1, root2, imaginary, discriminant float64

	fmt.Print("Enter the a, b, c of Quadratic equation = ")
	fmt.Scanln(&a, &b, &c)

	discriminant = (b * b) - (4 * a * c)

	if discriminant > 0 {
		root1 = (-b + math.Sqrt(discriminant)/(2*a))
		root2 = (-b - math.Sqrt(discriminant)/(2*a))
		fmt.Println("Two Distinct Real Roots Exist: root1 = ", root1, " and root2 = ", root2)
	} else if discriminant == 0 {
		root1 = -b / (2 * a)
		root2 = -b / (2 * a)
		fmt.Println("Two Equal and Real Roots Exist: root1 = ", root1, " and root2 = ", root2)
	} else if discriminant < 0 {
		root1 = -b / (2 * a)
		root2 = -b / (2 * a)
		imaginary = math.Sqrt(-discriminant) / (2 * a)
		fmt.Println("Two Distinct Complex Roots Exist: root1 = ", root1, "+", imaginary, " and root2 = ", root2, "-", imaginary)
	}
}

//--------TRIG FUNC--------//

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

//---------TRIG ENDS HERE------//

func slope() {
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

func Sum() {
	var num1, num2 int
	fmt.Print("Enter the First Number = ")
	fmt.Scanln(&num1)
	fmt.Print(" \n")

	fmt.Print("Enter the Second Number = ")
	fmt.Scanln(&num2)
	fmt.Print(" \n")

	fmt.Println("The Sum of num1 and num2  = ", num1+num2)
}

var choice int

//--------- my function --mustafain

func main() {
	fmt.Println("Choose the formula. \n")
	fmt.Println("1-Prime Function \n")
	fmt.Println("2-Distance Formula \n")
	fmt.Println("3-Quadratic Formula \n")
	fmt.Println("4-Slope formula \n")
	fmt.Println("5-Trignometric Identity \n")
	fmt.Println("6-Sum of numbers \n")
	fmt.Println("Enter your choice in number: ")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		CheckPrime()
	case 2:
		fmt.Println("Tuesday")
	case 3:
		quad()
	case 4:
		slope()
	case 5:
		trig()
	case 6:
		Sum()
	default:
		fmt.Println("Invalid")
	}
}
