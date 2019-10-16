package math

import (
	"math"
)

//@formatter:off

//Float_equal compare the equality of floats
//Ref: http://floating-point-gui.de/errors/comparison/
//compare floating point precision
func FloatEqual(a, b float64) bool {
	//var eps = EPSILON
	//if len(epsilon) > 0 {
	//	eps = epsilon[0]
	//}
	//if a == b {
	//	// shortcut, handles infinities
	//	return true
	//}
	//var diff = a - b
	//if diff < 0 {diff  = -diff}
	//if diff == 0 {diff = 0} //-0
	return a == b || Abs(a - b) < EPSILON
	//if a == b {
	//	// shortcut, handles infinities
	//	return true
	//} else if a == 0 || b == 0 || diff < eps {
	//	// a or b is zero or both are extremely close to it
	//	// relative error is less meaningful here
	//	return diff < eps || diff < (eps*eps)
	//}
	//
	//if a < 0 {a = -a}
	//if b < 0 {b = -b}
	//
	//// use relative error
	//return diff/math.Min(a+b, math.MaxFloat64) < eps
}

// Floor returns the greatest integer value less than or equal to x.
//
// Special cases are:
//	Floor(±0) = ±0
//	Floor(±Inf) = ±Inf
//	Floor(NaN) = NaN
func Floor(x float64) float64 {
	return math.Floor(x)
}

// Ceil returns the least integer value greater than or equal to x.
//
// Special cases are:
//	Ceil(±0) = ±0
//	Ceil(±Inf) = ±Inf
//	Ceil(NaN) = NaN
func Ceil(x float64) float64 {
	return math.Ceil(x)
}

// Trunc returns the integer value of x.
//
// Special cases are:
//	Trunc(±0) = ±0
//	Trunc(±Inf) = ±Inf
//	Trunc(NaN) = NaN
func Trunc(x float64) float64 {
	return math.Trunc(x)
}

// NaN returns an IEEE 754 ``not-a-number'' value.
func NaN() float64 {
	return math.NaN()
}


// IsNaN reports whether f is an IEEE 754 ``not-a-number'' value.
func IsNaN(f float64) bool {
	return math.IsNaN(f)
}

// Inf returns positive infinity if sign >= 0, negative infinity if sign < 0.
func Inf(sign int) float64 {
	return math.Inf(sign)
}

// IsInf reports whether f is an infinity, according to sign.
// If sign > 0, IsInf reports whether f is positive infinity.
// If sign < 0, IsInf reports whether f is negative infinity.
// If sign == 0, IsInf reports whether f is either infinity.
func IsInf(f float64, sign int) bool {
	return math.IsInf(f, sign)
}
