package ExpandGo

import (
	"fmt"
	"math"
)

func expand(x float64) float64 {
	fmt.Println("(x-1)^3 = x**3 - 3*x**2 + 3*x - 1")
	res1 := (math.Pow(x, 3))
	res2 := -(3 * math.Pow(x, 2))
	res3 := 3*x - 1
	result := res1 + res2 + res3
	return result
}
