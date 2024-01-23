package interval

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"strings"
)

type IInterval[T constraints.Integer | constraints.Float] interface {
	Lower() T
	SetLower(lower T) error
	Upper() T
	SetUpper(upper T) error
	LowerUnbounded() bool
	SetLowerUnbounded(lowerUnbounded bool) error
	UpperUnbounded() bool
	SetUpperUnbounded(upperUnbounded bool) error
	LowerIncluded() bool
	SetLowerIncluded(lowerIncluded bool) error
	UpperIncluded() bool
	SetUpperIncluded(upperIncluded bool) error
	String() string
	Equal(x Interval[T]) bool
	IsEmpty() bool
	LtBeginOf(x Interval[T]) bool
	LeEndOf(x Interval[T]) bool
	Contains(x Interval[T]) bool
	Has(value T) bool
	Intersect(x Interval[T]) Interval[T]
	Move(x T) Interval[T]
	Subtract(x Interval[T]) (Interval[T], Interval[T])
	Adjoin(x Interval[T]) Interval[T]
	Encompass(x Interval[T]) Interval[T]
}

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

func lowerIncluded[T constraints.Integer | constraints.Float](i, x Interval[T]) bool {
	return x.lowerIncluded && i.lowerIncluded
}

func upperIncluded[T constraints.Integer | constraints.Float](i, x Interval[T]) bool {
	return x.upperIncluded && i.upperIncluded
}

func (i Interval[T]) Lower() T {
	return i.lower
}

func (i *Interval[T]) SetLower(lower T) error {
	i.lower = lower
	return nil
}

func (i Interval[T]) Upper() T {
	return i.upper
}

func (i *Interval[T]) SetUpper(upper T) error {
	i.upper = upper
	return nil
}

func (i Interval[T]) LowerUnbounded() bool {
	return i.lowerUnbounded
}

func (i *Interval[T]) SetLowerUnbounded(lowerUnbounded bool) error {
	i.lowerUnbounded = lowerUnbounded
	return nil
}

func (i Interval[T]) UpperUnbounded() bool {
	return i.upperUnbounded
}

func (i *Interval[T]) SetUpperUnbounded(upperUnbounded bool) error {
	i.upperUnbounded = upperUnbounded
	return nil
}

func (i Interval[T]) LowerIncluded() bool {
	return i.lowerIncluded
}

func (i *Interval[T]) SetLowerIncluded(lowerIncluded bool) error {
	i.lowerIncluded = lowerIncluded
	return nil
}

func (i Interval[T]) UpperIncluded() bool {
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
		b.WriteByte('(')
	}
	fmt.Fprintf(&b, "%v", i.lower)
	b.WriteString(", ")
	fmt.Fprintf(&b, "%v", i.upper)
	if i.upperIncluded {
		b.WriteByte(']')
	} else {
		b.WriteByte(')')
	}
	if i.upperUnbounded {
		b.WriteByte('>')
	}
	return b.String()
}

// Equal returns true if receiver interval is equals x_interval_string interval.
func (i Interval[T]) Equal(x Interval[T]) bool {
	if x.upperUnbounded && i.upperUnbounded {
		return (i.lower == x.lower &&
			i.lowerIncluded == x.lowerIncluded &&
			i.lowerUnbounded == x.lowerUnbounded) || (x.IsEmpty() && i.IsEmpty())
	}
	if x.lowerUnbounded && i.lowerUnbounded {
		return (i.upper == x.upper &&
			i.upperIncluded == x.upperIncluded &&
			i.upperUnbounded == x.upperUnbounded) || (x.IsEmpty() && i.IsEmpty())
	}
	return (i.lower == x.lower &&
		i.upper == x.upper &&
		i.lowerIncluded == x.lowerIncluded &&
		i.upperIncluded == x.upperIncluded &&
		i.lowerUnbounded == x.lowerUnbounded &&
		i.upperUnbounded == x.upperUnbounded) || (x.IsEmpty() && i.IsEmpty())
}

// IsEmpty returns true if receiver interval has no value.
func (i Interval[T]) IsEmpty() bool {
	if i.upperUnbounded || i.lowerUnbounded {
		return false
	}
	if i.lower < i.upper {
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
	lowerSide := false
	upperSide := false
	if x.lowerUnbounded && !i.lowerUnbounded {
		lowerSide = false
	} else {
		if i.lowerUnbounded {
			lowerSide = true
		}
		if i.lower < x.lower {
			lowerSide = true
		}
		if i.lower == x.lower && !x.lowerIncluded {
			lowerSide = true
		}
		if i.lower == x.lower && i.lowerIncluded {
			lowerSide = true
		}
	}
	if x.upperUnbounded && !i.upperUnbounded {
		upperSide = false
	} else {
		if i.upperUnbounded {
			upperSide = true
		}
		if i.upper > x.upper {
			upperSide = true
		}
		if i.upper == x.upper && !x.upperIncluded {
			upperSide = true
		}
		if i.upper == x.upper && i.upperIncluded {
			upperSide = true
		}
	}
	return lowerSide && upperSide
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

// Intersect returns the intersection of receiver interval with x_interval_string interval.
func (i Interval[T]) Intersect(x Interval[T]) Interval[T] {
	if x.IsEmpty() || i.IsEmpty() {
		return Interval[T]{}
	}
	if !i.lowerUnbounded && !x.lowerUnbounded {
		if i.lower > x.lower {
			x.lower = i.lower
			x.lowerIncluded = i.lowerIncluded
		} else if i.lower == x.lower && !i.lowerIncluded {
			x.lowerIncluded = false
		}
	} else if x.lowerUnbounded && !i.lowerUnbounded {
		x.lower = i.lower
		x.lowerIncluded = i.lowerIncluded
		x.lowerUnbounded = false
	}
	if !i.upperUnbounded && !x.upperUnbounded {
		if i.upper < x.upper {
			x.upper = i.upper
			x.upperIncluded = i.upperIncluded
		} else if i.upper == x.upper && !i.upperIncluded {
			x.upperIncluded = false
		}
	} else if x.upperUnbounded && !i.upperUnbounded {
		x.upper = i.upper
		x.upperIncluded = i.upperIncluded
		x.upperUnbounded = false
	}
	return maybeEmpty(x)
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
	if x.lowerUnbounded {
		r1 = Interval[T]{}
	} else {
		lower := i.lower
		if i.lower > in.lower {
			lower = in.lower
		}
		r1 = maybeEmpty(Interval[T]{
			lower:          lower,
			lowerIncluded:  i.lowerIncluded,
			upper:          in.lower,
			upperIncluded:  !in.lowerIncluded,
			lowerUnbounded: i.lowerUnbounded,
			upperUnbounded: false,
		})
	}
	if x.upperUnbounded {
		r2 = Interval[T]{}
	} else {
		r2 = maybeEmpty(Interval[T]{
			lower:          in.upper,
			lowerIncluded:  !in.upperIncluded,
			upper:          i.upper,
			upperIncluded:  i.upperIncluded,
			lowerUnbounded: false,
			upperUnbounded: i.upperUnbounded,
		})
		if r2.lower > r2.upper && i.upperUnbounded {
			r2.upper = r2.lower
		}
	}
	return r1, r2
}

// Adjoin returns the union of two intervals, if the intervals are exactly
// adjacent, or the zero interval if they are not.
// Monica says: The Adjoin keyword in Swift is used to combine two ranges, specifically closed ranges, into a single range. It operates on ranges of any type that conforms to the Comparable protocol, allowing for flexibility in combining different types of ranges.
//
// The Adjoin keyword takes two closed ranges as input and returns a new closed range that represents the union of the two input ranges.
// If the input ranges are exactly adjacent, meaning they share a common boundary point,
// the Adjoin operation successfully combines them into a single range.
// However, if the input ranges are not adjacent, the Adjoin operation results in an empty range,
// represented by an empty closed range with the same lower and upper bounds.
func (i Interval[T]) Adjoin(x Interval[T]) Interval[T] {
	if x.IsEmpty() || i.IsEmpty() {
		return Interval[T]{}
	}
	if i.lowerUnbounded || i.upperUnbounded || x.upperUnbounded || x.lowerUnbounded {
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
	if i.lowerUnbounded {
		x.lowerUnbounded = true
	}
	if i.upperUnbounded {
		x.upperUnbounded = true
	}
	return x
}
