package math

import (
	"math"
)

const (
	PRECISION              = 12
	EPSILON                = 1.0e-12
	Pi                     = math.Pi
	Tau                    = 2.0 * Pi
	Sqrt2                  = math.Sqrt2
	E                      = math.E
	MaxFloat64             = math.MaxFloat64
	SmallestNonzeroFloat64 = math.SmallestNonzeroFloat64
)

const (
	x = iota
	y
)
// Sqrt returns the square root of x.
// Special cases are:
//	Sqrt(+Inf) = +Inf
//	Sqrt(±0) = ±0
//	Sqrt(x < 0) = NaN
//	Sqrt(NaN) = NaN
func Sqrt(x float64) float64 {
	return math.Sqrt(x)
}

// Pow returns x**y, the base-x exponential of y.
func Pow(x, y float64) float64 {
	return math.Pow(x, y)
}

// Pow10 returns 10**n, the base-10 exponential of n.
// Special cases are:
//	Pow10(n) =    0 for n < -323
//	Pow10(n) = +Inf for n > 308
func Pow10(n int) float64 {
	return math.Pow10(n)
}

