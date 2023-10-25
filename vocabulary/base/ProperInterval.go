package base

/*
Type representing a 'proper' Interval, i.e. any two-sided or one-sided interval.
Inv_not_point: lower /= upper
*/
type ProperInterval[T Number] struct {
	Interval[T]
}
