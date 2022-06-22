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
    fmt.Println("Sunday")
default:
    fmt.Println("Invalid")
}
}
