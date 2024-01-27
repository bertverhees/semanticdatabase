package interval

import (
	"github.com/bertverhees/interval"
	"semanticdatabase/generics"
)

/*
_type representing a 'proper' Interval, i.e. any two-sided or one-sided interval.
Inv_not_point: lower /= upper
*/
type ProperInterval[T generics.Number] struct {
	interval.Interval[T]
}
