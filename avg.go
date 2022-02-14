package math

func Average(vals ...float64) float64 {
	var n = len(vals)
	if n == 0 {
		return 0
	}
	return SumF64(vals) / float64(n)
}
