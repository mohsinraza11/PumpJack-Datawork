package NewTask

import (
	"math"
)

func Quadratic_formula(a float64, b float64, c float64) (float64, float64) {

	var result1 = (-b + math.Pow((b*b-4*a*c), (0.5))) / (2 * a)
	var result2 = (-b - math.Pow((b*b-4*a*c), (0.5))) / (2 * a)

	return result1, result2
}
