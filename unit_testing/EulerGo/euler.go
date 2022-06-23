package EulerGo

import (
	"math"
	"math/cmplx"
)

func euler() complex128 {
	total := 0 + 0i
	total = cmplx.Exp(math.Pi*1i) + 1.0
	return total
}
