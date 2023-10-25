package base

import "errors"

type Interval[T Number] struct {
	lower T
	upper T
	// lower boundary open (i.e. = -infinity)
	LowerUnbounded bool
	// upper boundary open (i.e. = +infinity)
	UpperUnbounded bool
	// lower boundary value included in range if not lower_unbounded.
	LowerIncluded bool
	// upper boundary value included in range if not upper_unbounded.
	UpperIncluded bool
}

func NewInterval[T Number]() *Interval[T] {
	return new(Interval[T])
}

type IntervalBuilder[T Number] struct {
	interval *Interval[T]
}

func NewIntervalBuilder[T Number]() *IntervalBuilder[T] {
	i := new(Interval[T])
	i.LowerIncluded = true
	i.UpperIncluded = true
	i.LowerUnbounded = true
	i.UpperUnbounded = true
	ib := new(IntervalBuilder[T])
	ib.interval = i
	return ib
}

func (i *IntervalBuilder[T]) setLower(lower T) *IntervalBuilder[T] {
	i.interval.lower = lower
	if i.interval.lower == i.interval.upper {
		errors.New("Lower cannot be equal with Upper")
	}
	i.interval.LowerUnbounded = false
	return i
}

func (i *IntervalBuilder[T]) setUpper(upper T) *IntervalBuilder[T] {
	i.interval.upper = upper
	if i.interval.lower == i.interval.upper {
		errors.New("Lower cannot be equal with Upper")
	}
	i.interval.UpperUnbounded = false
	return i
}

func (i *IntervalBuilder[T]) setLowerUnbounded(lowerUnbounded bool) *IntervalBuilder[T] {
	i.interval.LowerUnbounded = lowerUnbounded
	return i
}

func (i *IntervalBuilder[T]) setUpperUnbounded(upperUnbounded bool) *IntervalBuilder[T] {
	i.interval.UpperUnbounded = upperUnbounded
	return i
}

func (i *IntervalBuilder[T]) setLowerIncluded(lowerIncluded bool) *IntervalBuilder[T] {
	i.interval.LowerIncluded = lowerIncluded
	return i
}

func (i *IntervalBuilder[T]) setUpperIncluded(upperIncluded bool) *IntervalBuilder[T] {
	i.interval.UpperIncluded = upperIncluded
	return i
}

func (i *IntervalBuilder[T]) build() *Interval[T] {
	return i.interval
}

func LowerUnboundedInterval[T Number](upper T, UpperIncluded bool) *Interval[T] {
	return NewIntervalBuilder[T]().setUpper(upper).setLowerIncluded(false).setUpperIncluded(UpperIncluded).setLowerUnbounded(true).build()
}

func UpperUnboundedInterval[T Number](lower T, LowerIncluded bool) *Interval[T] {
	return NewIntervalBuilder[T]().setLower(lower).setUpperIncluded(false).setLowerIncluded(LowerIncluded).setUpperUnbounded(true).build()
}

func UnboundedInterval[T Number](lower T, LowerIncluded bool) *Interval[T] {
	return NewIntervalBuilder[T]().setUpperIncluded(false).setLowerIncluded(false).setUpperUnbounded(true).setLowerUnbounded(true).build()
}

func (i *Interval[T]) Has(value T) bool {
	if i.LowerUnbounded && i.UpperUnbounded {
		return true
	}
	returnValue := true
	if !i.LowerUnbounded && i.UpperUnbounded {
		if i.LowerIncluded {
			returnValue = value >= i.lower
		} else {
			returnValue = value > i.lower
		}
	} else if !i.UpperUnbounded && i.LowerUnbounded {
		if i.UpperIncluded {
			returnValue = value <= i.upper
		} else {
			returnValue = value < i.upper
		}
	} else if !i.LowerUnbounded && !i.UpperUnbounded {
		if i.LowerIncluded && i.UpperIncluded {
			returnValue = value >= i.lower && value <= i.upper
		} else if !i.LowerIncluded && i.UpperIncluded {
			returnValue = value > i.lower && value <= i.upper
		} else if i.LowerIncluded && !i.UpperIncluded {
			returnValue = value >= i.lower && value < i.upper
		}
	}
	return returnValue
}

func (i *Interval[T]) Intersects(other *Interval[T]) bool {
	return (i.LowerUnbounded && other.LowerUnbounded) ||
		(i.UpperUnbounded && other.UpperUnbounded) ||
		((i.lower-other.lower < 0) && (i.upper-other.upper < 0) && (other.lower-i.upper < 0)) ||
		((i.lower-other.lower < 0) && (i.upper-other.upper < 0) && (other.lower-i.upper < 0))
}

func (i *Interval[T]) Contains(other *Interval[T]) bool {
	otherHasLower := false
	otherHasUpper := false
	if other.LowerUnbounded {
		otherHasLower = i.LowerUnbounded
	} else {
		otherHasLower = i.Has(other.lower)
	}
	if other.UpperUnbounded {
		otherHasUpper = i.UpperUnbounded
	} else {
		otherHasUpper = i.Has(other.upper)
	}
	return otherHasUpper && otherHasLower
}
