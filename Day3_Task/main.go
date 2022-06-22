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

func quad(){
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
    CheckPrime() 
case 2:
    fmt.Println("Tuesday")
case 3:
    quad()
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
