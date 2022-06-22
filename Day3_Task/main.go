package main

import (
  "fmt"
)

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
    fmt.Println("Thursday")
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
