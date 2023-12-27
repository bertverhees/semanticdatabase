package base

import (
	"errors"
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

type IntervalBuilder[T generics.Number] struct {
	Builder
	interval *Interval[T]
}

func NewIntervalBuilder[T generics.Number]() *IntervalBuilder[T] {
	builder := &IntervalBuilder[T]{}
	builder.object = NewInterval[T]()
	return builder
	//i := new(Interval[T])
	//i.lowerIncluded = true
	//i.upperIncluded = true
	//i.lowerUnbounded = true
	//i.upperUnbounded = true
	//ib := new(IntervalBuilder[T])
	//ib.interval = i
	//return ib
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
		i.AddError(errors.New("Impossible interval constellation with lower being equal to upper and lowerincluded or upperincluded being false"))
	}
	if i.object.(*Interval[T]).Lower() > i.object.(*Interval[T]).Upper() {
		i.AddError(errors.New("Impossible interval constellation with lower being higher to upper"))
	}
	if i.errors != nil {
		return nil, i.errors
	} else {
		return i.object.(*Interval[T]), nil
	}
}

func LowerUnboundedInterval[T generics.Number](upper T, UpperIncluded bool) (*Interval[T], []error) {
	return NewIntervalBuilder[T]().setUpper(upper).setLowerIncluded(false).setUpperIncluded(UpperIncluded).setLowerUnbounded(true).Build()
}

func UpperUnboundedInterval[T generics.Number](lower T, LowerIncluded bool) (*Interval[T], []error) {
	return NewIntervalBuilder[T]().setLower(lower).setUpperIncluded(false).setLowerIncluded(LowerIncluded).setUpperUnbounded(true).Build()
}

func UnboundedInterval[T generics.Number](lower T, LowerIncluded bool) (*Interval[T], []error) {
	return NewIntervalBuilder[T]().setUpperIncluded(false).setLowerIncluded(false).setUpperUnbounded(true).setLowerUnbounded(true).Build()
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
		}
	}
	return returnValue
}

func (i *Interval[T]) Intersects(other *Interval[T]) bool {
	return (i.lowerUnbounded && other.lowerUnbounded) ||
		(i.upperUnbounded && other.upperUnbounded) ||
		((i.lower-other.lower < 0) && (i.upper-other.upper < 0) && (other.lower-i.upper < 0)) ||
		((other.lower-other.lower < 0) && (other.upper-other.upper < 0) && (i.lower-i.upper < 0))
}

func (i *Interval[T]) Contains(other *Interval[T]) bool {
	otherHasLower := false
	otherHasUpper := false
	if other.lowerUnbounded {
		otherHasLower = i.lowerUnbounded
	} else {
		otherHasLower = i.Has(other.lower)
	}
	if other.upperUnbounded {
		otherHasUpper = i.upperUnbounded
	} else {
		otherHasUpper = i.Has(other.upper)
	}
	return otherHasUpper && otherHasLower
}
