package math

func Average[T Num](values ...T) T {
	var n = len(values)
	if n == 0 {
		return 0
	}
	return Sum(values) / T(n)
}
