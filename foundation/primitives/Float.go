package primitives

import "math"

type Float struct{}

func ToFixedNumberOfDecimals(f float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(f*output)) / output
}
func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}
