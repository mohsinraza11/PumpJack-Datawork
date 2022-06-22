package main

import (
  "fmt"
)

var choice int
//--------- my function --mustafain


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
    fmt.Println("Thursday")
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
