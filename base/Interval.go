package base

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"strings"
)

type Interval[T constraints.Integer | constraints.Float] struct {
	// begin of this interval.
	lower T
	// if lowerIncluded is true, this interval is inclusive of the lower point.
	lowerIncluded bool
	// if lowerUnbounded is true, the value of lower is ignored and the interval is considered endles on the lowerside
	lowerUnbounded bool
	// end of this interval.
	upper T
	// if upperIncluded is true, this interval is inclusive of the upper point.
	upperIncluded bool
	// if upperUnbounded is true, the value of upper is ignored and the interval is considered endles on the upperside
	upperUnbounded bool
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

func NewInterval[T constraints.Integer | constraints.Float]() *Interval[T] {
	i := new(Interval[T])
	i.lowerUnbounded = false
	i.upperUnbounded = false
	i.lowerIncluded = true
	i.upperIncluded = true
	return i
}

func (i Interval[T]) String() string {
	var b strings.Builder
	if i.lowerUnbounded {
		b.WriteByte('<')
	}
	if i.lowerIncluded {
		b.WriteByte('[')
	} else {
		b.WriteByte('*')
	}
	fmt.Fprintf(&b, "%v", i.lower)
	b.WriteString(", ")
	fmt.Fprintf(&b, "%v", i.upper)
	if i.upperIncluded {
		b.WriteByte(']')
	} else {
		b.WriteByte('*')
	}
	if i.upperUnbounded {
		b.WriteByte('>')
	}
	return b.String()
}

// Equal returns true if receiver interval is equals x_interval_string interval.
func (i Interval[T]) Equal(x Interval[T]) bool {
	result := Interval[T]{}
	result.lowerUnbounded = x.lowerUnbounded == true && i.lowerUnbounded == true
	result.upperUnbounded = x.upperUnbounded == true && i.upperUnbounded == true
	if result.upperUnbounded {
		return (i.lower == x.lower &&
			i.lowerIncluded == x.lowerIncluded &&
			i.lowerUnbounded == x.lowerUnbounded &&
			i.upperUnbounded == x.upperUnbounded) || (x.IsEmpty() && i.IsEmpty())
	}
	if result.lowerUnbounded {
		return (i.upper == x.upper &&
			i.upperIncluded == x.upperIncluded &&
			i.lowerUnbounded == x.lowerUnbounded &&
			i.upperUnbounded == x.upperUnbounded) || (x.IsEmpty() && i.IsEmpty())
	}
	return (i.lower == x.lower &&
		i.upper == x.upper &&
		i.lowerIncluded == x.lowerIncluded &&
		i.upperIncluded == x.upperIncluded &&
		i.lowerUnbounded == x.lowerUnbounded &&
		i.upperUnbounded == x.upperUnbounded) || (x.IsEmpty() && i.IsEmpty())
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

// IsEmpty returns true if receiver interval has no value.
func (i Interval[T]) IsEmpty() bool {
	if i.upperUnbounded || i.lowerUnbounded {
		return false
	}
	if i.lower < i.upper && !i.upperUnbounded && !i.lowerUnbounded {
		return false
	} else if i.lower == i.upper {
		return !i.lowerIncluded || !i.upperIncluded
	}
	return true
}

// LtBeginOf returns true if receiver interval is less than begin of x_interval_string interval.
func (i Interval[T]) LtBeginOf(x Interval[T]) bool {
	if x.IsEmpty() {
		return false
	}
	if i.IsEmpty() {
		return false
	}
	if i.upperUnbounded {
		return false
	}
	if x.lowerUnbounded {
		return false
	}
	if i.upper < x.lower {
		return true
	} else if i.upper == x.lower {
		return !i.upperIncluded || !x.lowerIncluded
	}
	return false
}

// LeEndOf returns true if receiver interval is less than or equal to end of x_interval_string interval.
func (i Interval[T]) LeEndOf(x Interval[T]) bool {
	if x.IsEmpty() {
		return false
	}
	if i.IsEmpty() {
		return false
	}
	if i.upperUnbounded {
		return false
	} else if x.upperUnbounded {
		return true
	}
	if i.upper < x.upper {
		return true
	}
	if i.upper == x.upper {
		return (i.upperIncluded && x.upperIncluded) || !i.upperIncluded
	}
	return false
}

// Contains returns true if x_interval_string interval is completely covered by receiver interval.
func (i Interval[T]) Contains(x Interval[T]) bool {

	if x.IsEmpty() {
		return true
	}
	if i.IsEmpty() {
		return false
	}
	if i.lower > x.lower || (x.lowerUnbounded && !i.lowerUnbounded) {
		return false
	}
	if i.upper < x.upper || (x.upperUnbounded && !i.upperUnbounded) {
		return false
	}
	if i.lower < x.lower && i.upper > x.upper &&
		(unbounded(i, x) || bounded(i, x) ||
			((!x.lowerUnbounded && i.lowerUnbounded) && (!x.upperUnbounded && i.upperUnbounded))) {
		return true
	}
	if i.lower == x.lower && (i.lowerIncluded || !x.lowerIncluded) &&
		(unbounded(i, x) || bounded(i, x) ||
			((!x.lowerUnbounded && i.lowerUnbounded) && (!x.upperUnbounded && i.upperUnbounded))) {
		return true
	}
	return i.upper == x.upper && (i.upperIncluded || !x.upperIncluded) &&
		(unbounded(i, x) || bounded(i, x) ||
			((!x.lowerUnbounded && i.lowerUnbounded) && (!x.upperUnbounded && i.upperUnbounded)))
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

// Intersect returns the intersection of receiver interval with x_interval_string interval.
func (i Interval[T]) Intersect(x Interval[T]) Interval[T] {
	if x.IsEmpty() || i.IsEmpty() {
		return Interval[T]{}
	}
	result := Interval[T]{}
	result.lowerUnbounded = x.lowerUnbounded == true && i.lowerUnbounded == true
	result.upperUnbounded = x.upperUnbounded == true && i.upperUnbounded == true
	if result.upperUnbounded {
		result.upper = 0
		result.upperIncluded = false
	} else {
		if i.upperUnbounded && !x.upperUnbounded { //i.upperUnbounded
			result.upper = x.upper
			result.upperIncluded = x.upperIncluded
		} else if !i.upperUnbounded && x.upperUnbounded { //x.upperUnbounded
			result.upper = i.upper
			result.upperIncluded = i.upperIncluded
		} else if (i.upper >= x.upper && i.upperIncluded) || (i.upper > x.upper && !i.upperIncluded) {
			result.upper = x.upper
			result.upperIncluded = x.upperIncluded
		} else if (x.upper >= i.upper && x.upperIncluded) || (x.upper > i.upper && !x.upperIncluded) {
			result.upper = i.upper
			result.upperIncluded = i.upperIncluded
		}
	}
	if result.lowerUnbounded {
		result.lower = 0
		result.lowerIncluded = false
	} else {
		if i.lowerUnbounded && !x.lowerUnbounded { //i.upperUnbounded
			result.lower = x.lower
			result.lowerIncluded = x.lowerIncluded
		} else if !i.lowerUnbounded && x.lowerUnbounded { //x.upperUnbounded
			result.lower = i.lower
			result.lowerIncluded = i.lowerIncluded
		} else if (i.lower <= x.lower && i.lowerIncluded) || (i.lower < x.lower && !i.lowerIncluded) {
			result.lower = x.lower
			result.lowerIncluded = x.lowerIncluded
		} else if (x.lower <= i.lower && x.lowerIncluded) || (x.lower < i.lower && !x.lowerIncluded) {
			result.lower = i.lower
			result.lowerIncluded = i.lowerIncluded
		}
	}
	return maybeEmpty[T](result)
}

func maybeEmpty[T constraints.Integer | constraints.Float](x Interval[T]) Interval[T] {
	if x.IsEmpty() {
		return Interval[T]{}
	}
	return x
}

// Move returns an interval that adds number x_interval_string to begin and end of receiver interval.
func (i Interval[T]) Move(x T) Interval[T] {
	if i.IsEmpty() {
		return Interval[T]{}
	}
	return Interval[T]{
		lower:          i.lower + x,
		lowerIncluded:  i.lowerIncluded,
		upper:          i.upper + x,
		upperIncluded:  i.upperIncluded,
		upperUnbounded: i.upperUnbounded,
		lowerUnbounded: i.lowerUnbounded,
	}
}

// Subtract returns two intervals, one on the before of x_interval_string and one on the
// after of x_interval_string, corresponding to the subtraction of x_interval_string from the receiver
// interval. The returned intervals are always within the range of the
// receiver interval.
func (i Interval[T]) Subtract(x Interval[T]) (Interval[T], Interval[T]) {
	in := i.Intersect(x)
	if in.IsEmpty() {
		if i.LtBeginOf(x) {
			return i, Interval[T]{}
		}
		return Interval[T]{}, i
	}
	var r1, r2 Interval[T]
	var lower1, lower2, upper1, upper2 T
	var lowerIncluded1, lowerIncluded2, upperIncluded1, upperIncluded2 bool
	var lowerUnbounded1, lowerUnbounded2, upperUnbounded1, upperUnbounded2 bool
	if i.lower < x.lower || i.lowerUnbounded {
		lower1 = i.lower
		lowerIncluded1 = i.lowerIncluded
		lowerUnbounded1 = i.lowerUnbounded
		upper1 = x.lower
		upperIncluded1 = !x.lowerIncluded
		upperUnbounded1 = false
	}
	if i.lower == x.lower && !i.lowerUnbounded {
		lower1 = i.lower
		lowerIncluded1 = i.lowerIncluded
		lowerUnbounded1 = false
		upper1 = x.lower
		upperIncluded1 = !x.lowerIncluded
		upperUnbounded1 = false
	}
	if i.lower > x.lower && !i.lowerUnbounded {
		lower1 = i.lower
		if i.lowerIncluded != x.lowerIncluded {
			lowerIncluded1 = false
		}
		if i.lowerIncluded == x.lowerIncluded {
			lowerIncluded1 = i.lowerIncluded
		}
		lowerUnbounded1 = false
		upper1 = i.lower
		upperIncluded1 = !x.lowerIncluded
		upperUnbounded1 = false
	}
	if i.upper > x.upper || i.upperUnbounded {
		upper2 = i.upper
		upperIncluded2 = i.upperIncluded
		upperUnbounded2 = i.upperUnbounded
		lower2 = x.upper
		lowerIncluded2 = !x.lowerIncluded
		lowerUnbounded2 = false
	}
	if i.upper == x.upper && !i.upperUnbounded {
		upper2 = i.upper
		upperIncluded2 = i.upperIncluded
		upperUnbounded2 = false
		lower2 = x.upper
		lowerIncluded2 = !x.upperIncluded
		lowerUnbounded2 = false
	}
	if i.upper < x.upper || !i.upperUnbounded {
		upper2 = i.upper
		if i.upperIncluded != x.upperIncluded {
			upperIncluded2 = false
		}
		if i.upperIncluded == x.upperIncluded {
			upperIncluded2 = i.upperIncluded
		}
		upperUnbounded2 = false
		lower2 = i.upper
		lowerIncluded2 = !x.upperIncluded
		lowerUnbounded2 = false
	}
	r1 = maybeEmpty[T](Interval[T]{
		lower:          lower1,
		lowerIncluded:  lowerIncluded1,
		upper:          upper1,
		upperIncluded:  upperIncluded1,
		lowerUnbounded: lowerUnbounded1,
		upperUnbounded: upperUnbounded1,
	})
	r2 = maybeEmpty[T](Interval[T]{
		lower:          lower2,
		lowerIncluded:  lowerIncluded2,
		upper:          upper2,
		upperIncluded:  upperIncluded2,
		lowerUnbounded: lowerUnbounded2,
		upperUnbounded: upperUnbounded2,
	})
	return r1, r2
}

// Adjoin returns the union of two intervals, if the intervals are exactly
// adjacent, or the zero interval if they are not.
func (i Interval[T]) Adjoin(x Interval[T]) Interval[T] {
	if x.IsEmpty() || i.IsEmpty() {
		return Interval[T]{}
	}
	if i.lower == x.upper && (i.lowerIncluded || x.upperIncluded) {
		x.upper = i.upper
		x.upperIncluded = i.upperIncluded
		return x
	}
	if i.upper == x.lower && (i.upperIncluded || x.lowerIncluded) {
		x.lower = i.lower
		x.lowerIncluded = i.lowerIncluded
		return x
	}
	return Interval[T]{}
}

// Encompass returns an interval that covers the exact extents of two intervals.
func (i Interval[T]) Encompass(x Interval[T]) Interval[T] {
	if x.IsEmpty() {
		return i
	}
	if i.IsEmpty() {
		return x
	}
	if i.lower < x.lower {
		x.lower = i.lower
		x.lowerIncluded = i.lowerIncluded
	} else if i.lower == x.lower && i.lowerIncluded {
		x.lowerIncluded = true
	}
	if i.upper > x.upper {
		x.upper = i.upper
		x.upperIncluded = i.upperIncluded
	} else if i.upper == x.upper && i.upperIncluded {
		x.upperIncluded = true
	}
	return x
}
