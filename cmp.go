package math

import "math"

// Min returns the min of x and y.
//
// Special cases are:
//
//	Min(x, -Inf) = Min(-Inf, x) = -Inf
//	Min(x, NaN) = Min(NaN, x) = NaN
//	Min(-0, ±0) = Min(±0, -0) = -0
func Min[T Num](x, y T) T {
	return T(math.Min(float64(x), float64(y)))
}

// Max returns the max of x and y.
//
// Special cases are:
//
//	Max(x, +Inf) = Max(+Inf, x) = +Inf
//	Max(x, NaN) = Max(NaN, x) = NaN
//	Max(+0, ±0) = Max(±0, +0) = +0
//	Max(-0, -0) = -0
func Max[T Num](x, y T) T {
	return T(math.Max(float64(x), float64(y)))
}
