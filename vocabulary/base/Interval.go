package base

import (
	"golang.org/x/exp/constraints"
)

type Interval[T constraints.Ordered] struct {
	Lower T
	Upper T
	// lower boundary open (i.e. = -infinity)
	LowerUnbounded bool
	// upper boundary open (i.e. = +infinity)
	UpperUnbounded bool
	// lower boundary value included in range if not lower_unbounded.
	LowerIncluded bool
	// upper boundary value included in range if not upper_unbounded.
	UpperIncluded bool
}

func NewInterval[T constraints.Ordered]() *Interval[T] {
	return new(Interval[T])
}
