package Task3

import (
	"fmt"
	"math"
)

func Distance_formula(x1 float64, x2 float64, y1 float64, y2 float64) float64 {
	var distance float64
	distance = math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
	return math.Round(distance*100) / 100
}

func Dist2() {
	var x1, x2, y1, y2 float64
	fmt.Print("Enter the First X co-oridnate(x1): ")
	fmt.Scanln(&x1)
	fmt.Print("Enter the Second X Co-oridnate(x2):")
	fmt.Scanln(&x2)
	fmt.Print("Enter the Second Y Co-oridnate(y1):")
	fmt.Scanln(&y1)
	fmt.Print("Enter the Second Y Co-oridnate(y2):")
	fmt.Scanln(&y2)
	var output float64
	output = Distance_formula(x1, x2, y1, y2)
	fmt.Println("The distance between the two points is:", output)

}
