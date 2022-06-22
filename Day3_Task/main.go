package main

import (
  "fmt"
)

func slope(){
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

fmt.Println("Enter your choice in number: ")
fmt.Scanln(&choice)


switch choice {
case 1:
    fmt.Println("Monday")
case 2:
    fmt.Println("Tuesday")
case 3:
    fmt.Println("Tuesday")
case 4:
    slope()
case 5:
    fmt.Println("Friday")
case 6:
    fmt.Println("Saturday")
case 7:
	Sum()
default:
    fmt.Println("Invalid")
}
}
