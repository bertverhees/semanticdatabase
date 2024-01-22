package base

import (
	"fmt"
	interval2 "semanticdatabase/foundation/interval"
	"semanticdatabase/generics"
)

type IntervalBuilder[T generics.Number] struct {
	Builder
	interval *interval2.Interval[T]
}

func NewIntervalBuilder[T generics.Number]() *IntervalBuilder[T] {
	builder := &IntervalBuilder[T]{}
	builder.object = interval2.NewInterval[T]()
	return builder
}

func (i *IntervalBuilder[T]) setLower(lower T) *IntervalBuilder[T] {
	i.AddError(i.object.(*interval2.Interval[T]).SetLower(lower))
	return i
}

func (i *IntervalBuilder[T]) setUpper(upper T) *IntervalBuilder[T] {
	i.AddError(i.object.(*interval2.Interval[T]).SetUpper(upper))
	return i
}

func (i *IntervalBuilder[T]) setLowerUnbounded(lowerUnbounded bool) *IntervalBuilder[T] {
	i.AddError(i.object.(*interval2.Interval[T]).SetLowerUnbounded(lowerUnbounded))
	return i
}

func (i *IntervalBuilder[T]) setUpperUnbounded(upperUnbounded bool) *IntervalBuilder[T] {
	i.AddError(i.object.(*interval2.Interval[T]).SetUpperUnbounded(upperUnbounded))
	return i
}

func (i *IntervalBuilder[T]) setLowerIncluded(lowerIncluded bool) *IntervalBuilder[T] {
	i.AddError(i.object.(*interval2.Interval[T]).SetLowerIncluded(lowerIncluded))
	return i
}

func (i *IntervalBuilder[T]) setUpperIncluded(upperIncluded bool) *IntervalBuilder[T] {
	i.AddError(i.object.(*interval2.Interval[T]).SetUpperIncluded(upperIncluded))
	return i
}

func (i *IntervalBuilder[T]) Build() (*interval2.Interval[T], []error) {
	if i.object.(*interval2.Interval[T]).Lower() == i.object.(*interval2.Interval[T]).Upper() && (i.object.(*interval2.Interval[T]).LowerIncluded() == false || i.object.(*interval2.Interval[T]).UpperIncluded() == false) {
		i.AddError(fmt.Errorf("Impossible interval constellation with lower: %v == upper: %v and lowerincluded or upperincluded being false", i.object.(*interval2.Interval[T]).Lower(), i.object.(*interval2.Interval[T]).Lower()))
	}
	if i.object.(*interval2.Interval[T]).Lower() > i.object.(*interval2.Interval[T]).Upper() {
		i.AddError(fmt.Errorf("Impossible interval constellation with lower: %v being higher to upper: %v", i.object.(*interval2.Interval[T]).Lower(), i.object.(*interval2.Interval[T]).Upper()))
	}
	if i.errors != nil {
		return nil, i.errors
	} else {
		return i.object.(*interval2.Interval[T]), nil
	}
}
