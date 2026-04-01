package utils

import "math"

func RoundTo(f float64, p uint) float64 {
	r := math.Pow(10, float64(p))
	return math.Round(f*r) / r
}
