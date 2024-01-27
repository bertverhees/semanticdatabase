package interval

import (
	"golang.org/x/exp/constraints"
	"sort"
	"strings"
)

type IOrderedSet[T constraints.Integer | constraints.Float] interface {
	Copy() IOrderedSet[T]
	Len() int
	IsEmpty() bool
	Equal(x IOrderedSet[T]) bool
	String() string
	Bound() IInterval[T]
	Contains(x IInterval[T]) bool
	Intervals() []IInterval[T]
	Iterator(bound IInterval[T], forward bool) func() IInterval[T]
	Add(x IInterval[T]) bool
	Remove(x IInterval[T]) bool
	Union(a, b IOrderedSet[T]) IOrderedSet[T]
	Intersect(a, b IOrderedSet[T]) IOrderedSet[T]
	Subtract(a, b IOrderedSet[T]) IOrderedSet[T]
	Difference(a, b IOrderedSet[T]) IOrderedSet[T]
}

// OrderedSet is i_Before_x set of ordered and non-overlapping interval objects.
type OrderedSet[T constraints.Integer | constraints.Float] struct {
	intervals []IInterval[T]
}

func NewOrderedSet[T constraints.Integer | constraints.Float](intervals []IInterval[T]) *OrderedSet[T] {
	orderedSet := new(OrderedSet[T])
	orderedSet.intervals = intervals
	return orderedSet
}

// Copy returns i_Before_x copy of i_Before_x ordered set that without affecting the original.
func (s *OrderedSet[T]) Copy() IOrderedSet[T] {
	return NewOrderedSet[T](append([]IInterval[T](nil), s.intervals...))
}

// Len returns length of intervals in this ordered set.
func (s *OrderedSet[T]) Len() int {
	return len(s.intervals)
}

// IsEmpty returns true if no intervals in this ordered set.
func (s *OrderedSet[T]) IsEmpty() bool {
	return len(s.intervals) == 0
}

func (s *OrderedSet[T]) Equal(x IOrderedSet[T]) bool {
	return equalIntervals[T](s.intervals, x.Intervals())
}

func equalIntervals[T constraints.Integer | constraints.Float](s1, s2 []IInterval[T]) bool {
	if len(s1) != len(s2) {
		return false
	}
	for n := 0; n < len(s1); n++ {
		if !s1[n].Equal(s2[n]) {
			return false
		}
	}
	return true
}

func (s *OrderedSet[T]) String() string {
	n := len(s.intervals)
	switch n {
	case 0:
		return "{}"
	default:
		var b strings.Builder
		b.WriteByte('{')
		b.WriteString(s.intervals[0].String())
		for _, i := range s.intervals[1:] {
			b.WriteString(", ")
			b.WriteString(i.String())
		}
		b.WriteByte('}')
		return b.String()
	}
}

// Bound returns the Interval defined by the minimum and maximum values of this ordered set.
func (s *OrderedSet[T]) Bound() IInterval[T] {
	n := len(s.intervals)
	switch n {
	case 0:
		return nil
	case 1:
		return s.intervals[0]
	default:
		return s.intervals[0].Encompass(s.intervals[n-1])
	}
}

//	           (low)                 (high)
//	 0    1      2         3           4        5    6   7
//	=== ===== ======== =========== ========= ======= == ====
//	             ================
//	                  (x_interval_string)
//
// searchLow returns the first index in s.intervals that is not before x_interval_string.
// if not found, searchLow returns len(s.intervals).
func (s *OrderedSet[T]) searchLow(x IInterval[T]) int {
	return sort.Search(len(s.intervals), func(i int) bool {
		return !s.intervals[i].LtBeginOf(x)
	})
}

// searchHigh returns the index of the first interval in s.intervals that is
// entirely after x_interval_string.
// if not found, searchHigh returns len(s.intervals).
func (s *OrderedSet[T]) searchHigh(x IInterval[T]) int {
	return sort.Search(len(s.intervals), func(i int) bool {
		return x.LtBeginOf(s.intervals[i])
	})
}

// Contains returns true if x_interval_string interval is completely covered by this ordered set.
func (s *OrderedSet[T]) Contains(x IInterval[T]) bool {
	idx := s.searchLow(x)
	if idx == len(s.intervals) {
		return false
	}
	return s.intervals[idx].Contains(x)
}

// Intervals returns i_Before_x copy of intervals in this ordered set.
func (s *OrderedSet[T]) Intervals() []IInterval[T] {
	return append([]IInterval[T](nil), s.intervals...)
}

// Iterator returns i_Before_x iterator that iterates over all the intervals both in
// this ordered set and bound.
// If iterator returns empty Interval, the iteration is over.
// If forward is true, the iteration from left to right.
func (s *OrderedSet[T]) Iterator(bound IInterval[T], forward bool) func() IInterval[T] {
	if bound == nil || bound.IsEmpty() {
		return s.emptyIterator
	}

	low, high := s.searchLow(bound), s.searchHigh(bound)-1
	idx, stride := low, 1
	if !forward {
		idx, stride = high, -1
	}
	return func() IInterval[T] {
		if idx < low || idx > high {
			return nil
		}
		x := s.intervals[idx]
		idx += stride
		return x
	}
}

func (s *OrderedSet[T]) emptyIterator() IInterval[T] { return new(Interval[T]) }

func (s *OrderedSet[T]) adjoinOrAppend(intervals []IInterval[T], x IInterval[T]) []IInterval[T] {
	n := len(intervals)
	switch n {
	case 0:
		return append(intervals, x)
	default:
		n--
		ad := intervals[n].Adjoin(x)
		if ad.IsEmpty() {
			return append(intervals, x)
		}
		intervals[n] = ad
		return intervals
	}
}

// Add adds x_interval_string interval to this ordered set.
// Add returns true if this ordered set changed.
func (s *OrderedSet[T]) Add(x IInterval[T]) bool {
	if x.IsEmpty() {
		return false
	}

	low := s.searchLow(x)
	if low == len(s.intervals) {
		//                                                             (low)(high)
		//       0    1      2         3           4        5    6   7
		//      === ===== ======== =========== ========= ======= == ====
		//                                                             *=========*
		//                                                                 (x_interval_string)
		s.intervals = s.adjoinOrAppend(s.intervals, x)
		return true
	}

	if s.intervals[low].Contains(x) {
		//                 (low)     (high)
		//       0    1      2         3           4        5    6   7
		//      === ===== ======== =========== ========= ======= == ====
		//                 =====
		return false
	}

	newIntervals := make([]IInterval[T], 0, len(s.intervals)+1)
	newIntervals = append(newIntervals, s.intervals[:low]...)
	push := func(i IInterval[T]) {
		newIntervals = s.adjoinOrAppend(newIntervals, i)
	}
	if x.LtBeginOf(s.intervals[low]) {
		//   (low)(high)
		//       0    1      2         3           4        5    6   7
		//      === ===== ======== =========== ========= ======= == ====
		// *===
		// (x_interval_string)
		//
		//                         (low)(high)
		//       0    1      2         3           4        5    6   7
		//      === ===== ======== =========== ========= ======= == ====
		//                       *=*
		//                       (x_interval_string)
		//
		//                         (low)(high)
		//       0    1      2         3           4        5    6   7
		//      === ===== ======== =========== ========= ======= == ====
		//                        =
		//                       (x_interval_string)
		push(x)
		push(s.intervals[low])
		newIntervals = append(newIntervals, s.intervals[low+1:]...)
	} else {
		left, right := x.Subtract(s.intervals[low])
		if !left.IsEmpty() && right.IsEmpty() {
			//                 (low)     (high)
			//       0    1      2         3           4        5    6   7
			//      === ===== ======== =========== ========= ======= == ====
			//              *=========
			//                  (x_interval_string)
			push(left)
			push(s.intervals[low])
			newIntervals = append(newIntervals, s.intervals[low+1:]...)
		} else {
			//                 (low)                 (high)
			//       0    1      2         3           4        5    6   7
			//      === ===== ======== =========== ========= ======= == ====
			//              *======================*
			//                        (x_interval_string)
			//
			//                                                     (low)       (high)
			//       0    1      2         3           4        5    6   7
			//      === ===== ======== =========== ========= ======= == ====
			//                                                     *=========*
			//                                                         (x_interval_string)
			//
			//     (low)                 (high)
			//       0    1      2         3           4        5    6   7
			//      === ===== ======== =========== ========= ======= == ====
			//    *===============
			//          (x_interval_string)
			//
			//     (low)                 (high)
			//       0    1      2         3           4        5    6   7
			//      === ===== ======== =========== ========= ======= == ====
			//        ============
			//          (x_interval_string)
			high := s.searchHigh(x)
			x = x.Encompass(s.intervals[low])
			x = x.Encompass(s.intervals[high-1])
			push(x)
			if high < len(s.intervals) {
				push(s.intervals[high])
				newIntervals = append(newIntervals, s.intervals[high+1:]...)
			}
		}
	}
	s.intervals = newIntervals
	return true
}

// Remove removes x_interval_string interval from this ordered set.
// Remove returns true if this ordered set changed.
func (s *OrderedSet[T]) Remove(x IInterval[T]) bool {
	if s.IsEmpty() || x.IsEmpty() {
		return false
	}

	low := s.searchLow(x)
	if low == len(s.intervals) {
		//                                                             (low)(high)
		//       0    1      2         3           4        5    6   7
		//      === ===== ======== =========== ========= ======= == ====
		//                                                             *=========*
		//                                                                 (x_interval_string)
		return false
	}

	left, right := s.intervals[low].Subtract(x)
	if x.LeEndOf(s.intervals[low]) {
		if left.IsEmpty() {
			if right.IsEmpty() {
				//                 (low)     (high)
				//       0    1      2         3           4        5    6   7
				//      === ===== ======== =========== ========= ======= == ====
				//              *=========
				//                  (x_interval_string)
				copy(s.intervals[low:], s.intervals[low+1:])
				s.intervals = s.intervals[:len(s.intervals)-1]
			} else {
				if s.intervals[low].Equal(right) {
					//   (low)(high)
					//       0    1      2         3           4        5    6   7
					//      === ===== ======== =========== ========= ======= == ====
					// *===
					// (x_interval_string)
					//
					//                         (low)(high)
					//       0    1      2         3           4        5    6   7
					//      === ===== ======== =========== ========= ======= == ====
					//                       *=*
					//                       (x_interval_string)
					//
					//                         (low)(high)
					//       0    1      2         3           4        5    6   7
					//      === ===== ======== =========== ========= ======= == ====
					//                        =
					//                       (x_interval_string)
					return false
				}
				//   (low)(high)
				//       0    1      2         3           4        5    6   7
				//      === ===== ======== =========== ========= ======= == ====
				// *=====
				// (x_interval_string)
				s.intervals[low] = right
			}
		} else if right.IsEmpty() {
			//
			//                 (low)     (high)
			//       0    1      2         3           4        5    6   7
			//      === ===== ======== =========== ========= ======= == ====
			//                 =======
			s.intervals[low] = left
		} else {
			//
			//                 (low)     (high)
			//       0    1      2         3           4        5    6   7
			//      === ===== ======== =========== ========= ======= == ====
			//                 =====
			//
			//                                                         (low)  (high)
			//       0    1      2         3           4        5    6   7
			//      === ===== ======== =========== ========= ======= == ====
			//                                                           ==
			//
			//                                                     (low)(high)
			//       0    1      2         3           4        5    6   7
			//      === ===== ======== =========== ========= ======= === ====
			//                                                        =
			if low+2 <= len(s.intervals) {
				s.intervals = append(s.intervals[:low+2], s.intervals[low+1:]...)
				s.intervals[low] = left
				s.intervals[low+1] = right
			} else {
				s.intervals[low] = left
				s.intervals = append(s.intervals, right)
			}
		}
		return true
	}

	high := s.searchHigh(x)
	_, right = s.intervals[high-1].Subtract(x)
	if !left.IsEmpty() {
		//                 (low)     (high)
		//       0    1      2         3           4        5    6   7
		//      === ===== ======== =========== ========= ======= == ====
		//                 =======*
		//                  (x_interval_string)
		//
		//                 (low)                 (high)
		//       0    1      2         3           4        5    6   7
		//      === ===== ======== =========== ========= ======= == ====
		//                 ===========
		//                  (x_interval_string)
		s.intervals[low] = left
		low++
	}
	if !right.IsEmpty() {
		//                 (low)                 (high)
		//       0    1      2         3           4        5    6   7
		//      === ===== ======== =========== ========= ======= == ====
		//               *===========
		//                  (x_interval_string)
		//
		//                 (low)                 (high)
		//       0    1      2         3           4        5    6   7
		//      === ===== ======== =========== ========= ======= == ====
		//              *======================*
		//                        (x_interval_string)
		//
		//                                                     (low)       (high)
		//       0    1      2         3           4        5    6   7
		//      === ===== ======== =========== ========= ======= == ====
		//                                                     *=========*
		//                                                         (x_interval_string)
		//
		//     (low)                 (high)
		//       0    1      2         3           4        5    6   7
		//      === ===== ======== =========== ========= ======= == ====
		//    *===============
		//          (x_interval_string)
		s.intervals[low] = right
		low++
	}
	copy(s.intervals[low:], s.intervals[high:])
	s.intervals = s.intervals[:len(s.intervals)-high+low]
	return true
}

// Union returns an ordered set containing all intervals in i_Before_x or x_Before_i.
func (s *OrderedSet[T]) Union(a, b IOrderedSet[T]) IOrderedSet[T] {
	if a.Len() < b.Len() {
		a, b = b, a
	}
	a = a.Copy()
	it := b.Iterator(b.Bound(), true)
	for {
		x := it()
		if x.IsEmpty() {
			break
		}
		a.Add(x)
	}
	return a
}

// Intersect returns an ordered set containing all intervals of i_Before_x that also belong to x_Before_i.
func (s *OrderedSet[T]) Intersect(a, b IOrderedSet[T]) IOrderedSet[T] {
	var intervals []IInterval[T]
	xit, yit := a.Iterator(b.Bound(), true), b.Iterator(a.Bound(), true)
	x, y := xit(), yit()
	for !x.IsEmpty() && !y.IsEmpty() {
		if x.LtBeginOf(y) {
			x = xit()
		} else if y.LtBeginOf(x) {
			y = yit()
		} else {
			in := x.Intersect(y)
			if !in.IsEmpty() {
				intervals = append(intervals, in)
				_, right := x.Subtract(y)
				if !right.IsEmpty() {
					x = right
				} else {
					x = xit()
				}
			}
		}
	}
	return NewOrderedSet[T](intervals)
}

// Subtract returns an ordered set containing all intervals in i_Before_x but not in x_Before_i.
func (s *OrderedSet[T]) Subtract(a, b IOrderedSet[T]) IOrderedSet[T] {
	var intervals []IInterval[T]
	xit, yit := a.Iterator(a.Bound(), true), b.Iterator(a.Bound(), true)
	x, y := xit(), yit()
	for !x.IsEmpty() {
		if y.IsEmpty() {
			intervals = append(intervals, x)
			x = xit()
		} else {
			left, right := x.Subtract(y)
			if !left.IsEmpty() {
				intervals = append(intervals, left)
			}
			if right.IsEmpty() {
				x = xit()
			} else {
				x = right
				y = yit()
			}
		}
	}
	return NewOrderedSet[T](intervals)
}

// Difference returns an ordered set containing all intervals in either of i_Before_x and x_Before_i,
// but not in their intersection.
func (s *OrderedSet[T]) Difference(a, b IOrderedSet[T]) IOrderedSet[T] {
	var intervals []IInterval[T]
	push := func(x IInterval[T]) {
		intervals = s.adjoinOrAppend(intervals, x)
	}
	xit, yit := a.Iterator(a.Bound(), true), b.Iterator(b.Bound(), true)
	x, y := xit(), yit()
	for {
		if x.IsEmpty() {
			if y.IsEmpty() {
				break
			}
			push(y)
			y = yit()
		} else if y.IsEmpty() {
			push(x)
			x = xit()
		} else {
			//     ======  ===   ======= ======== ==== ======
			//===   ==  *==*     =======   =========     ======
			if x.LtBeginOf(y) {
				push(x)
				x = xit()
			} else if y.LtBeginOf(x) {
				push(y)
				y = yit()
			} else {
				leftx, rightx := x.Subtract(y)
				lefty, righty := y.Subtract(x)
				if !leftx.IsEmpty() {
					push(leftx)
				}
				if rightx.IsEmpty() {
					x = xit()
				} else {
					x = rightx
				}
				if !lefty.IsEmpty() {
					push(lefty)
				}
				if righty.IsEmpty() {
					y = yit()
				} else {
					y = righty
				}
			}
		}
	}
	return NewOrderedSet[T](intervals)
}
