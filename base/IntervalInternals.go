package base

import "golang.org/x/exp/constraints"

func iL_l_xL[T constraints.Integer | constraints.Float](i, x Interval[T]) bool {
	return (i.lower < x.lower) || (i.lowerUnbounded && !x.lowerUnbounded)
}

func iL_le_xL[T constraints.Integer | constraints.Float](i, x Interval[T]) bool {
	return (((i.lower <= x.lower) && (i.lowerIncluded || (i.lowerIncluded == x.lowerIncluded))) || i.lowerUnbounded || !(!i.lowerUnbounded && x.lowerUnbounded))
}

func iL_e_xL[T constraints.Integer | constraints.Float](i, x Interval[T]) bool {
	return ((i.lower == x.lower) && (!i.lowerUnbounded && !x.lowerUnbounded) && (i.lowerIncluded == x.lowerIncluded)) || (i.lowerUnbounded && x.lowerUnbounded)
}

func iU_e_xU[T constraints.Integer | constraints.Float](i, x Interval[T]) bool {
	return ((i.upper == x.upper) && (!i.upperUnbounded && !x.upperUnbounded) && (i.upperIncluded == x.upperIncluded)) || (i.upperUnbounded && x.upperUnbounded)
}

func iU_g_xU[T constraints.Integer | constraints.Float](i, x Interval[T]) bool {
	return (i.upper > x.upper) || (i.upperUnbounded && !x.upperUnbounded)
}

func iU_ge_xU[T constraints.Integer | constraints.Float](i, x Interval[T]) bool {
	return (((i.upper >= x.upper) && (i.upperIncluded || (i.upperIncluded == x.upperIncluded))) || i.upperUnbounded || !(!i.upperUnbounded && x.upperUnbounded))
}

func lowerIncluded[T constraints.Integer | constraints.Float](i, x Interval[T]) bool {
	return x.lowerIncluded && i.lowerIncluded
}

func upperIncluded[T constraints.Integer | constraints.Float](i, x Interval[T]) bool {
	return x.upperIncluded && i.upperIncluded
}

func included[T constraints.Integer | constraints.Float](i, x Interval[T]) bool {
	return lowerIncluded(i, x) && upperIncluded(i, x)
}

func lowerNotIncluded[T constraints.Integer | constraints.Float](i, x Interval[T]) bool {
	return !x.lowerIncluded && !i.lowerIncluded
}

func upperNotIncluded[T constraints.Integer | constraints.Float](i, x Interval[T]) bool {
	return !x.upperIncluded && !i.upperIncluded
}

func notIncluded[T constraints.Integer | constraints.Float](i, x Interval[T]) bool {
	return lowerIncluded(i, x) && upperIncluded(i, x)
}

func lowerUnbounded[T constraints.Integer | constraints.Float](i, x Interval[T]) bool {
	return x.lowerUnbounded && i.lowerUnbounded
}

func upperUnbounded[T constraints.Integer | constraints.Float](i, x Interval[T]) bool {
	return x.upperUnbounded && i.upperUnbounded
}

func unbounded[T constraints.Integer | constraints.Float](i, x Interval[T]) bool {
	return lowerUnbounded(i, x) && upperUnbounded(i, x)
}

func lowerBounded[T constraints.Integer | constraints.Float](i, x Interval[T]) bool {
	return !x.lowerUnbounded && !i.lowerUnbounded
}

func upperBounded[T constraints.Integer | constraints.Float](i, x Interval[T]) bool {
	return !x.upperUnbounded && !i.upperUnbounded
}

func bounded[T constraints.Integer | constraints.Float](i, x Interval[T]) bool {
	return lowerBounded(i, x) && upperBounded(i, x)
}
