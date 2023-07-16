package math

// Sum - sum numeric values
func Sum[T Num](o []T) T {
	var t T
	for _, v := range o {
		t += v
	}
	return t
}
