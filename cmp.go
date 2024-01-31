package math

// Min - returns the min of x and y.
func Min[T Num](x, y T) T {
	return min(x, y)
}

// Max - returns the max of x and y.
func Max[T Num](x, y T) T {
	return max(x, y)
}

// Equals - equality of two numeric values
func Equals[T Num](x, y T) bool {
	var a, b = min(x, y), max(x, y)
	return (a == b) || (float64(b-a) < EPSILON)
}

// FloatEqual - compare the equality of floats
// Ref: http://floating-point-gui.de/errors/comparison/
// compare floating point precision
func FloatEqual(x, y float64) bool {
	//var eps = EPSILON
	//if len(epsilon) > 0 {
	//	eps = epsilon[0]
	//}
	//if x == y {
	//	// shortcut, handles infinities
	//	return true
	//}
	//var diff = x - y
	//if diff < 0 {diff  = -diff}
	//if diff == 0 {diff = 0} //-0
	return Equals(x, y)
	//if x == y {
	//	// shortcut, handles infinities
	//	return true
	//} else if x == 0 || y == 0 || diff < eps {
	//	// x or y is zero or both are extremely close to it
	//	// relative error is less meaningful here
	//	return diff < eps || diff < (eps*eps)
	//}
	//
	//if x < 0 {x = -x}
	//if y < 0 {y = -y}
	//
	//// use relative error
	//return diff/math.Min(x+y, math.MaxFloat64) < eps
}
