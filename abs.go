package math

import (
	"golang.org/x/exp/constraints"
)

type Num interface {
	constraints.Integer | constraints.Float
}

// Abs returns the absolute value of x.
//
// Special cases are:
//
//	Abs(±Inf) = +Inf
//	Abs(NaN) = NaN
func Abs[T Num](v T) T {
	if v < 0 {
		return -v
	}
	return v
}
