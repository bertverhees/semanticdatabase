package base

import "golang.org/x/exp/constraints"

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
