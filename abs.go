package math

// Abs returns the absolute value of x.
//
// Special cases are:
//	Abs(±Inf) = +Inf
//	Abs(NaN) = NaN
func Abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}
//func Abs(x float64) float64 {
//	return math.Abs(x)
//}
