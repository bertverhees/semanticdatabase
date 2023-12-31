package base

import (
	"fmt"
	"semanticdatabase/generics"
)

type Interval[T generics.Number] struct {
	lower T
	upper T
	// lower boundary open (i.e. = -infinity)
	lowerUnbounded bool
	// upper boundary open (i.e. = +infinity)
	upperUnbounded bool
	// lower boundary value included in range if not lower_unbounded.
	lowerIncluded bool
	// upper boundary value included in range if not upper_unbounded.
	upperIncluded bool
}

func (i *Interval[T]) Lower() T {
	return i.lower
}

func (i *Interval[T]) SetLower(lower T) error {
	i.lower = lower
	return nil
}

func (i *Interval[T]) Upper() T {
	return i.upper
}

func (i *Interval[T]) SetUpper(upper T) error {
	i.upper = upper
	return nil
}

func (i *Interval[T]) LowerUnbounded() bool {
	return i.lowerUnbounded
}

func (i *Interval[T]) SetLowerUnbounded(lowerUnbounded bool) error {
	i.lowerUnbounded = lowerUnbounded
	return nil
}

func (i *Interval[T]) UpperUnbounded() bool {
	return i.upperUnbounded
}

func (i *Interval[T]) SetUpperUnbounded(upperUnbounded bool) error {
	i.upperUnbounded = upperUnbounded
	return nil
}

func (i *Interval[T]) LowerIncluded() bool {
	return i.lowerIncluded
}

func (i *Interval[T]) SetLowerIncluded(lowerIncluded bool) error {
	i.lowerIncluded = lowerIncluded
	return nil
}

func (i *Interval[T]) UpperIncluded() bool {
	return i.upperIncluded
}

func (i *Interval[T]) SetUpperIncluded(upperIncluded bool) error {
	i.upperIncluded = upperIncluded
	return nil
}

func NewInterval[T generics.Number]() *Interval[T] {
	return new(Interval[T])
}

func LowerUnboundedInterval[T generics.Number](upper T, UpperIncluded bool) (*Interval[T], []error) {
	return NewIntervalBuilder[T]().setUpper(upper).setLowerIncluded(false).setUpperIncluded(UpperIncluded).setLowerUnbounded(true).Build()
}

func UpperUnboundedInterval[T generics.Number](lower T, LowerIncluded bool) (*Interval[T], []error) {
	return NewIntervalBuilder[T]().setLower(lower).setUpperIncluded(false).setLowerIncluded(LowerIncluded).setUpperUnbounded(true).Build()
}

func UnboundedInterval[T generics.Number](lower T, upper T) (*Interval[T], []error) {
	return NewIntervalBuilder[T]().setLower(lower).setUpper(upper).setUpperIncluded(false).setLowerIncluded(false).setUpperUnbounded(true).setLowerUnbounded(true).Build()
}

func (i *Interval[T]) Has(value T) bool {
	if i.lowerUnbounded && i.upperUnbounded {
		return true
	}
	returnValue := true
	if !i.lowerUnbounded && i.upperUnbounded {
		if i.lowerIncluded {
			returnValue = value >= i.lower
		} else {
			returnValue = value > i.lower
		}
	} else if !i.upperUnbounded && i.lowerUnbounded {
		if i.upperIncluded {
			returnValue = value <= i.upper
		} else {
			returnValue = value < i.upper
		}
	} else if !i.lowerUnbounded && !i.upperUnbounded {
		if i.lowerIncluded && i.upperIncluded {
			returnValue = value >= i.lower && value <= i.upper
		} else if !i.lowerIncluded && i.upperIncluded {
			returnValue = value > i.lower && value <= i.upper
		} else if i.lowerIncluded && !i.upperIncluded {
			returnValue = value >= i.lower && value < i.upper
		} else if !i.lowerIncluded && !i.upperIncluded {
			returnValue = value > i.lower && value < i.upper
		}
	}
	return returnValue
}

func (i *Interval[T]) BoundaryHas(value T, included bool) bool {
	if i.lowerUnbounded && i.upperUnbounded {
		return true
	}
	returnValue := true
	if !i.lowerUnbounded && i.upperUnbounded {
		if i.lowerIncluded {
			returnValue = value >= i.lower
			if included {
				returnValue = value > i.lower
			}
		} else {
			returnValue = value > i.lower
		}
	} else if !i.upperUnbounded && i.lowerUnbounded {
		if i.upperIncluded {
			returnValue = value <= i.upper
			if included {
				returnValue = value < i.lower
			}
		} else {
			returnValue = value < i.upper
		}
	} else if !i.lowerUnbounded && !i.upperUnbounded {
		if i.lowerIncluded && i.upperIncluded {
			returnValue = value >= i.lower && value <= i.upper
		} else if !i.lowerIncluded && i.upperIncluded {
			returnValue = value > i.lower && value <= i.upper
		} else if i.lowerIncluded && !i.upperIncluded {
			returnValue = value >= i.lower && value < i.upper
		} else if !i.lowerIncluded && !i.upperIncluded {
			returnValue = value > i.lower && value < i.upper
		}
		if included {
			returnValue = value > i.lower && value < i.upper
		}

	}
	return returnValue
}

/**
 * True if there is any overlap between intervals represented by Current and
 * `other'. True if at least one limit of other is strictly inside the limits
 * of this interval.
 */
func (i *Interval[T]) Intersects(other *Interval[T]) bool {
	b1 := i.lowerUnbounded && other.lowerUnbounded
	b2 := i.upperUnbounded && other.upperUnbounded
	b3 := (i.lower-other.lower < 0) && (i.upper-other.upper < 0) && (other.lower-i.upper < 0)
	b4 := (other.lower-i.lower < 0) && (other.upper-i.upper < 0) && (i.lower-other.upper < 0)
	b5 := other.Contains(i)
	b6 := i.Contains(other)
	b7 := other.lowerUnbounded && i.Has(other.upper)
	b8 := other.upperUnbounded && i.Has(other.lower)
	return b1 || b2 || b3 || b4 || b5 || b6 || b7 || b8
}

func (i *Interval[T]) IntersectsAsInterVal(other *Interval[T]) *Interval[T] {
	var lowerUnbounded, upperUnbounded, lowerIncluded, upperIncluded bool
	var lower, upper T
	if i.Intersects(other) {
		lowerIncluded = i.lowerIncluded && other.lowerIncluded
		upperIncluded = i.upperIncluded && other.upperIncluded
		lowerUnbounded = i.lowerUnbounded && other.lowerUnbounded
		upperUnbounded = i.upperUnbounded && other.upperUnbounded
		if !i.upperUnbounded && other.upperUnbounded {
			upper = i.upper
		} else if !other.upperUnbounded && i.upperUnbounded {
			upper = other.upper
		} else if !i.upperUnbounded && !other.upperUnbounded {
			upper = i.upper
			if i.upper > other.upper {
				upper = other.upper
			}
		}
		if !i.lowerUnbounded && other.lowerUnbounded {
			lower = i.lower
		} else if !other.lowerUnbounded && i.lowerUnbounded {
			lower = other.lower
		} else if !i.lowerUnbounded && !other.lowerUnbounded {
			lower = i.lower
			if i.lower < other.lower {
				lower = other.lower
			}
		}
		if lower > upper {
			if lowerUnbounded {
				lower = upper
			}
			if upperUnbounded {
				upper = lower
			}
		}
		return &Interval[T]{lower, upper, lowerUnbounded, upperUnbounded, lowerIncluded, upperIncluded}
	}
	return nil
}

/*
The included properties of "other" are ignored, but handled as true
*/
func (i *Interval[T]) Contains(other *Interval[T]) bool {
	otherHasLower := false
	otherHasUpper := false
	if other.lowerUnbounded {
		otherHasLower = i.lowerUnbounded
	} else {
		otherHasLower = i.BoundaryHas(other.lower, other.lowerIncluded)
	}
	if other.upperUnbounded {
		otherHasUpper = i.upperUnbounded
	} else {
		otherHasUpper = i.BoundaryHas(other.upper, other.upperIncluded)
	}
	return otherHasUpper && otherHasLower
}

type IntervalBuilder[T generics.Number] struct {
	Builder
	interval *Interval[T]
}

func NewIntervalBuilder[T generics.Number]() *IntervalBuilder[T] {
	builder := &IntervalBuilder[T]{}
	builder.object = NewInterval[T]()
	return builder
}

func (i *IntervalBuilder[T]) setLower(lower T) *IntervalBuilder[T] {
	i.AddError(i.object.(*Interval[T]).SetLower(lower))
	return i
}

func (i *IntervalBuilder[T]) setUpper(upper T) *IntervalBuilder[T] {
	i.AddError(i.object.(*Interval[T]).SetUpper(upper))
	return i
}

func (i *IntervalBuilder[T]) setLowerUnbounded(lowerUnbounded bool) *IntervalBuilder[T] {
	i.AddError(i.object.(*Interval[T]).SetLowerUnbounded(lowerUnbounded))
	return i
}

func (i *IntervalBuilder[T]) setUpperUnbounded(upperUnbounded bool) *IntervalBuilder[T] {
	i.AddError(i.object.(*Interval[T]).SetUpperUnbounded(upperUnbounded))
	return i
}

func (i *IntervalBuilder[T]) setLowerIncluded(lowerIncluded bool) *IntervalBuilder[T] {
	i.AddError(i.object.(*Interval[T]).SetLowerIncluded(lowerIncluded))
	return i
}

func (i *IntervalBuilder[T]) setUpperIncluded(upperIncluded bool) *IntervalBuilder[T] {
	i.AddError(i.object.(*Interval[T]).SetUpperIncluded(upperIncluded))
	return i
}

func (i *IntervalBuilder[T]) Build() (*Interval[T], []error) {
	if i.object.(*Interval[T]).Lower() == i.object.(*Interval[T]).Upper() && (i.object.(*Interval[T]).LowerIncluded() == false || i.object.(*Interval[T]).UpperIncluded() == false) {
		i.AddError(fmt.Errorf("Impossible interval constellation with lower: %v == upper: %v and lowerincluded or upperincluded being false", i.object.(*Interval[T]).Lower(), i.object.(*Interval[T]).Lower()))
	}
	if i.object.(*Interval[T]).Lower() > i.object.(*Interval[T]).Upper() {
		i.AddError(fmt.Errorf("Impossible interval constellation with lower: %v being higher to upper: %v", i.object.(*Interval[T]).Lower(), i.object.(*Interval[T]).Upper()))
	}
	if i.errors != nil {
		return nil, i.errors
	} else {
		return i.object.(*Interval[T]), nil
	}
}
