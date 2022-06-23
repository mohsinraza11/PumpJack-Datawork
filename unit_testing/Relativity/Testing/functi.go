package Testing

import "math"

func Relativity(Mass float64, velocity float64) float64 {
	c := 299792458.0
	generalRelativity := float64(Mass) * math.Pow(c, 2)
	LF_sub := 1 - (float64(velocity) / math.Pow(c, 2))
	lorentzFactor := math.Pow(LF_sub, (1 / 2))
	r := generalRelativity / lorentzFactor
	return r
}
