package math

// Mid2D computes the midpoint of coordinates
func Mid2D[T Num](a, b []T) []T {
	return []T{Mid(a[x], b[x]), Mid(a[y], b[y])}
}

// Mid computes the mean of two values
func Mid[T Num](x, y T) T {
	return (x + y) / T(2)
}
