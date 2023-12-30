package base

import (
	"errors"
	"reflect"
	"semanticdatabase/generics"
	"testing"
)

func TestIntervalBuilder_build_errors(t *testing.T) {
	type testCase[T generics.Number] struct {
		name   string
		i      IntervalBuilder[T]
		want   *Interval[T]
		errors []error
	}
	tests := []testCase[int]{
		{name: "Test Builder for Default Interval", i: *NewIntervalBuilder[int](), want: nil,
			errors: []error{errors.New("Impossible interval constellation with lower: 0 == upper: 0 and lowerincluded or upperincluded being false")}},
		{name: "Test Builder for Default lower > upper", i: *NewIntervalBuilder[int]().setLower(4).setUpper(3), want: nil,
			errors: []error{errors.New("Impossible interval constellation with lower: 4 being higher to upper: 3")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got_errors := tt.i.Build()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Build() = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got_errors, tt.errors) {
				t.Errorf("Errors from Build() = %v, errors %v", got_errors, tt.errors)
			}
		})
	}
}

func TestIntervalBuilder_build(t *testing.T) {
	type testCase[T generics.Number] struct {
		name   string
		i      IntervalBuilder[T]
		want   *Interval[T]
		errors []error
	}
	tests := []testCase[int]{
		{name: "Test for Default Items", i: *NewIntervalBuilder[int](), want: nil,
			errors: []error{errors.New("Impossible interval constellation with lower: 0 == upper: 0 and lowerincluded or upperincluded being false")}},
		{name: "Test for Default included true and unbounded false booleans",
			i:    *NewIntervalBuilder[int]().setLower(0).setUpper(1).setUpperIncluded(true).setLowerIncluded(true),
			want: &Interval[int]{lower: 0, upper: 1, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
		{name: "Test for setted Unbounded booleans",
			i:    *NewIntervalBuilder[int]().setLower(0).setUpper(1).setLowerUnbounded(true).setUpperUnbounded(true),
			want: &Interval[int]{lower: 0, upper: 1, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: false, upperIncluded: false}},
		{name: "Test for setted Included and unbounded booleans",
			i:    *NewIntervalBuilder[int]().setLower(0).setUpper(1).setLowerIncluded(true).setUpperIncluded(true).setUpperUnbounded(true).setLowerUnbounded(true),
			want: &Interval[int]{lower: 0, upper: 1, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got_errors := tt.i.Build()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Build() = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got_errors, tt.errors) {
				t.Errorf("Errors from Build() = %v, errors %v", got_errors, tt.errors)
			}
		})
	}
}

func TestInterval_Contains(t *testing.T) {
	type args[T generics.Number] struct {
		other *Interval[T]
	}
	type testCase[T generics.Number] struct {
		name string
		i    Interval[T]
		args args[T]
		want bool
	}
	tests := []testCase[int]{
		// container included true unbounded false, arg included false unbounded false
		{name: "0-1 Down and out container: 0..10, li true ui true, lu false uu false, arg:-5..-1 li false ui false, lu false uu false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: -1, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: false},
		{name: "0-2 Down and intersecting: container: 0..10, li true ui true, lu false uu false, arg:-2..2 li false ui false, lu false uu false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -2, upper: 2, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: false},
		{name: "0-3 Down Limit and In: container: 0..10, li true ui true, lu false uu false, arg:0..6 li false ui false, lu false uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 0, upper: 6, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: true},
		{name: "0-4 In: container: 0..10, li true ui true, lu false uu false, arg:2..6 li false ui false, lu false uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 2, upper: 6, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: true},
		{name: "0-5 Up Limit and In: container: 0..10, li true ui true, lu false uu false, arg:8..10 li false ui false, lu false uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 0, upper: 6, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: true},
		{name: "0-6 Up and intersecting: container: 0..10, li true ui true, lu false uu false, arg:8..12 li false ui false, lu false uu false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 8, upper: 12, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: false},
		{name: "0-7 Up and out: container: 0..10, li true ui true, lu false uu false, arg:12..16 li false ui false, lu false uu false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 12, upper: 16, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: false},
		// container included upper false included lower true unbounded false, arg included false unbounded false
		{name: "0-8 Down and out container: 0..10, li true ui false, lu false uu false,  arg:-5..-1 li false ui false, lu false uu false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: -5, upper: -1, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: false},
		{name: "0-9 Down and intersecting: container: 0..10, li true ui false, lu false uu false,  arg:-2..2 li false ui false, lu false uu false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: -2, upper: 2, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: false},
		{name: "0-10 Down Limit and In: container: 0..10, li true ui false, lu false uu false,  arg:0..6 li false ui false, lu false uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: 0, upper: 6, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: true},
		{name: "0-11 In: container: 0..10, li true ui false, lu false uu false,  arg:2..6 li false ui false, lu false uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: 2, upper: 6, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: true},
		{name: "0-12 Up Limit and In: container: 0..10, li true ui false, lu false uu false,  arg:8..10 li false ui false, lu false uu false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: 8, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: false},
		{name: "0-13 Up and intersecting: container: 0..10, li true ui false, lu false uu false,  arg:8..12 li false ui false, lu false uu false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: 8, upper: 12, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: false},
		{name: "0-14 Up and out: container: 0..10, li true ui false, lu false uu false,  arg:12..16 li false ui false, lu false uu false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: 12, upper: 16, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: false},
		// container included false unbounded false, arg included false unbounded false
		{name: "0-15 Down and out container: 0..10,  li false ui false,  lu false uu false,  arg:-5..-1 li false ui false, lu false uu false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: -5, upper: -1, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: false},
		{name: "0-16 Down and intersecting: container: 0..10,  li false ui false,  lu false uu false,  arg:-2..2 li false ui false, lu false uu false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: -2, upper: 2, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: false},
		{name: "0-17 Down Boundary and In: container: 0..10,  li false ui false,  lu false uu false,  arg:0..6 li false ui false, lu false uu false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: 0, upper: 6, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: false},
		{name: "0-18 In: container: 0..10, included:upper:  li false ui false,  lu false uu false,  arg:2..6 li false ui false, lu false uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: 2, upper: 6, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: true},
		{name: "0-19 Up Boundary and In: container: 0..10,  li false ui false,  lu false uu false,  arg:8..10 li false ui false, lu false uu false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: 8, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: false},
		{name: "0-20 Up and intersecting: container: 0..10,  li false ui false,  lu false uu false,  arg:8..12 li false ui false, lu false uu false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: 8, upper: 12, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: false},
		{name: "0-21 Up and out: container: 0..10,  li false ui false,  lu false uu false,  arg:12..16 li false ui false, lu false uu false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: 12, upper: 16, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: false},
		// container included false unbounded false, arg included false unbounded false
		{name: "0-22 Down and out container: 0..10, included false and lower unbounded: true upper unbounded: false,  arg:-5..-1 included: false, unbounded: false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: false, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: -5, upper: -1, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: true},
		{name: "0-23 Down and intersecting: container: 0..10, included false and lower unbounded: true upper unbounded: false,  arg:-2..2 included: false, unbounded: false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: false, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: -2, upper: 2, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: true},
		{name: "0-24 Down Limit and In: container: 0..10, included false and lower unbounded: true upper unbounded: false,  arg:0..6 included: false, unbounded: false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: false, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: 0, upper: 6, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: true},
		{name: "0-25 In: container: 0..10, included:upper: included false and lower unbounded: true upper unbounded: false,  arg:2..6 included: false, unbounded: false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: false, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: 2, upper: 6, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: true},
		{name: "0-26 Up Limit and In: container: 0..10, included false and lower unbounded: true upper unbounded: false,  arg:8..10 included: false, unbounded: false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: false, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: 8, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: false},
		{name: "0-27 Up and intersecting: container: 0..10, included false and lower unbounded: true upper unbounded: false,  arg:8..12 included: false, unbounded: false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: false, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: 8, upper: 12, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: false},
		{name: "0-28 Up and out: container: 0..10, included false and lower unbounded: true upper unbounded: false,  arg:12..16 included: false, unbounded: false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: false, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: 12, upper: 16, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: false},
		// container included false unbounded true, arg included false unbounded false
		{name: "0-29 Down and out container: 0..10, included false and lower unbounded: true upper unbounded: true,  arg:-5..-1 included: false, unbounded: false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: false, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: -5, upper: -1, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: true},
		{name: "0-30 Down and intersecting: container: 0..10, included false and lower unbounded: true upper unbounded: true,  arg:-2..2 included: false, unbounded: false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: false, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: -2, upper: 2, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: true},
		{name: "0-31 Down Limit and In: container: 0..10, included false and lower unbounded: true upper unbounded: true,  arg:0..6 included: false, unbounded: false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: false, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: 0, upper: 6, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: true},
		{name: "0-32 In: container: 0..10, included:upper: included false and lower unbounded: true upper unbounded: true,  arg:2..6 included: false, unbounded: false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: false, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: 2, upper: 6, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: true},
		{name: "0-33 Up Limit and In: container: 0..10, included false and lower unbounded: true upper unbounded: true,  arg:8..10 included: false, unbounded: false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: false, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: 8, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: true},
		{name: "0-34 Up and intersecting: container: 0..10, included false and lower unbounded: true upper unbounded: true,  arg:8..12 included: false, unbounded: false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: false, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: 8, upper: 12, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: true},
		{name: "0-35 Up and out: container: 0..10, included false and lower unbounded: true upper unbounded: true,  arg:12..16 included: false, unbounded: false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: false, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: 12, upper: 16, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: true},
		// container included true unbounded false, arg included false unbounded low true up false
		{name: "1-1 Down and out container: 0..10, included: true and unbounded: false,  arg:-5..-1 included: false, lower unbounded true upper unbounded: false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: -1, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: false},
		{name: "1-2 Down and intersecting: container: 0..10, included: true and unbounded: false,  arg:-2..2 included: false, lower unbounded true upper: false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -2, upper: 2, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: false},
		{name: "1-3 Down Limit and In: container: 0..10, included: true and unbounded: false,  arg:0..6 included: false, lower unbounded true upper: false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 0, upper: 6, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: false},
		{name: "1-4 In: container: 0..10, included: true and unbounded: false,  arg:2..6 included: false, lower unbounded true upper: false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 2, upper: 6, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: false},
		{name: "1-5 Up Limit and In: container: 0..10, included: true and unbounded: false,  arg:8..10 included: false, lower unbounded true upper: false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: 0, upper: 6, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: false},
		{name: "1-6 Up and intersecting: container: 0..10, included: true and unbounded: false,  arg:8..12 included: false, lower unbounded true upper: false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 8, upper: 12, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: false},
		{name: "1-7 Up and out: container: 0..10, included: true and unbounded: false,  arg:12..16 included: false, lower unbounded true upper: false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 12, upper: 16, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: false},
		// container included upper false included lower true unbounded false, arg included false unbounded low true up false
		{name: "1-8 Down and out container: 0..10, included:upper: false lower: true and unbounded: false,  arg:-5..-1 included: false, lower unbounded true upper: false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: -5, upper: -1, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: false},
		{name: "1-9 Down and intersecting: container: 0..10, included:upper: false lower: true and unbounded: false,  arg:-2..2 included: false, lower unbounded true upper: false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: -2, upper: 2, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: false},
		{name: "1-10 Down Limit and In: container: 0..10, included:upper: false lower: true and unbounded: false,  arg:0..6 included: false, lower unbounded true upper: false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: 0, upper: 6, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: false},
		{name: "1-11 In: container: 0..10, included:upper: false lower: true and unbounded: false,  arg:2..6 included: false, lower unbounded true upper: false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: 2, upper: 6, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: false},
		{name: "1-12 Up Limit and In: container: 0..10, included:upper: false lower: true and unbounded: false,  arg:8..10 included: lower unbounded true upper, unbounded: false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: 8, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: false},
		{name: "1-13 Up and intersecting: container: 0..10, included:upper: false lower: true and unbounded: false,  arg:8..12 included: false, lower unbounded true upper: false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: 8, upper: 12, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: false},
		{name: "1-14 Up and out: container: 0..10, included:upper: false lower: true and unbounded: false,  arg:12..16 included: false, lower unbounded true upper: false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: 12, upper: 16, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: false},
		// container included false unbounded false, arg included false unbounded low true up false
		{name: "1-15 Down and out container: 0..10, included false and unbounded: false,  arg:-5..-1 included: false, lower unbounded true upper: false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: -5, upper: -1, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: false},
		{name: "1-16 Down and intersecting: container: 0..10, included false and unbounded: false,  arg:-2..2 included: false, lower unbounded true upper: false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: -2, upper: 2, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: false},
		{name: "1-17 Down Boundary and In: container: 0..10, included false and unbounded: false,  arg:0..6 included: false, lower unbounded true upper: false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: 0, upper: 6, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: false},
		{name: "1-18 In: container: 0..10, included:upper: included false and unbounded: false,  arg:2..6 included: false, lower unbounded true upper: false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: 2, upper: 6, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: false},
		{name: "1-19 Up Boundary and In: container: 0..10, included false and unbounded: false,  arg:8..10 included: false, lower unbounded true upper: false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: 8, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: false},
		{name: "1-20 Up and intersecting: container: 0..10, included false and unbounded: false,  arg:8..12 included: false, lower unbounded true upper: false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: 8, upper: 12, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: false},
		{name: "1-21 Up and out: container: 0..10, included false and unbounded: false,  arg:12..16 included: false, lower unbounded true upper: false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: 12, upper: 16, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: false},
		// container included false unbounded lower true upper false, arg included false unbounded low true up false
		{name: "1-22 Down and out container: 0..10, included false and lower unbounded: true upper unbounded: false,  arg:-5..-1 included: false, upper unbounded true upper: false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: false, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: -5, upper: -1, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: true},
		{name: "1-23 Down and intersecting: container: 0..10, included false and lower unbounded: true upper unbounded: false,  arg:-2..2 included: false, lower unbounded true upper: false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: false, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: -2, upper: 2, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: true},
		{name: "1-24 Down Limit and In: container: 0..10, included false and lower unbounded: true upper unbounded: false,  arg:0..6 included: false, lower unbounded true upper: false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: false, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: 0, upper: 6, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: true},
		{name: "1-25 In: container: 0..10, included:upper: included false and lower unbounded: true upper unbounded: false,  arg:2..6 included: false, lower unbounded true upper: false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: false, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: 2, upper: 6, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: true},
		{name: "1-26 Up Limit and In: container: 0..10, included false and lower unbounded: true upper unbounded: false,  arg:8..10 included: false, lower unbounded true upper: false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: false, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: 8, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: false},
		{name: "1-27 Up and intersecting: container: 0..10, included false and lower unbounded: true upper unbounded: false,  arg:8..12 included: false, lower unbounded true upper: false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: false, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: 8, upper: 12, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: false},
		{name: "1-28 Up and out: container: 0..10, included false and lower unbounded: true upper unbounded: false,  arg:12..16 included: false, lower unbounded true upper: false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: false, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: 12, upper: 16, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: false},
		// container included false unbounded true, arg included false unbounded low true up false
		{name: "1-29 Down and out container: 0..10, included false and lower unbounded: true upper unbounded: true,  arg:-5..-1 included: false, lower unbounded true upper: false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: false, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: -5, upper: -1, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: true},
		{name: "1-30 Down and intersecting: container: 0..10, included false and lower unbounded: true upper unbounded: true,  arg:-2..2 included: false, lower unbounded true upper: false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: false, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: -2, upper: 2, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: true},
		{name: "1-31 Down Limit and In: container: 0..10, included false and lower unbounded: true upper unbounded: true,  arg:0..6 included: false, lower unbounded true upper: false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: false, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: 0, upper: 6, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: true},
		{name: "1-32 In: container: 0..10, included:upper: included false and lower unbounded: true upper unbounded: true,  arg:2..6 included: false, lower unbounded true upper: false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: false, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: 2, upper: 6, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: true},
		{name: "1-33 Up Limit and In: container: 0..10, included false and lower unbounded: true upper unbounded: true,  arg:8..10 included: false, lower unbounded true upper: false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: false, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: 8, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: true},
		{name: "1-34 Up and intersecting: container: 0..10, included false and lower unbounded: true upper unbounded: true,  arg:8..12 included: false, lower unbounded true upper: false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: false, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: 8, upper: 12, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: true},
		{name: "1-35 Up and out: container: 0..10, included false and lower unbounded: true upper unbounded: true,  arg:12..16 included: false, lower unbounded true upper: false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: false, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: 12, upper: 16, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: true},
		// container included true unbounded false, arg included false unbounded low false up true
		{name: "2-1 Down and out container: 0..10, included: true and unbounded: false,  arg:-5..-1 included: false, upper unbounded true lower: false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: -1, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: false, upperIncluded: false}},
			want: false},
		{name: "2-2 Down and intersecting: container: 0..10, included: true and unbounded: false,  arg:-2..2 included: false, upper unbounded true lower: false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -2, upper: 2, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: false, upperIncluded: false}},
			want: false},
		{name: "2-3 Down Limit and In: container: 0..10, included: true and unbounded: false,  arg:0..6 included: false, upper unbounded true lower: false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 0, upper: 6, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: false, upperIncluded: false}},
			want: false},
		{name: "2-4 In: container: 0..10, included: true and unbounded: false,  arg:2..6 included: false, upper unbounded true lower: false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 2, upper: 6, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: false, upperIncluded: false}},
			want: false},
		{name: "2-5 Up Limit and In: container: 0..10, included: true and unbounded: false,  arg:8..10 included: false, upper unbounded true lower: false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: 0, upper: 6, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: false, upperIncluded: false}},
			want: false},
		{name: "2-6 Up and intersecting: container: 0..10, included: true and unbounded: false,  arg:8..12 included: false, upper unbounded true lower: false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 8, upper: 12, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: false, upperIncluded: false}},
			want: false},
		{name: "2-7 Up and out: container: 0..10, included: true and unbounded: false,  arg:12..16 included: false, upper unbounded true lower: false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 12, upper: 16, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: false, upperIncluded: false}},
			want: false},
		// container included upper false included lower true unbounded false, arg included false unbounded low true up false
		{name: "2-8 Down and out container: 0..10, included:upper: false lower: true and unbounded: false,  arg:-5..-1 included: false, upper unbounded true lower: false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: -5, upper: -1, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: false, upperIncluded: false}},
			want: false},
		{name: "2-9 Down and intersecting: container: 0..10, included:upper: false lower: true and unbounded: false,  arg:-2..2 included: false, upper unbounded true lower: false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: -2, upper: 2, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: false, upperIncluded: false}},
			want: false},
		{name: "2-10 Down Limit and In: container: 0..10, included:upper: false lower: true and unbounded: false,  arg:0..6 included: false, upper unbounded true lower: false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: 0, upper: 6, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: false, upperIncluded: false}},
			want: false},
		{name: "2-11 In: container: 0..10, included:upper: false lower: true and unbounded: false,  arg:2..6 included: false, upper unbounded true lower: false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: 2, upper: 6, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: false, upperIncluded: false}},
			want: false},
		{name: "2-12 Up Limit and In: container: 0..10, included:upper: false lower: true and unbounded: false,  arg:8..10 included: false, upper unbounded true lower: false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: 8, upper: 10, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: false, upperIncluded: false}},
			want: false},
		{name: "2-13 Up and intersecting: container: 0..10, included:upper: false lower: true and unbounded: false,  arg:8..12 included: false, upper unbounded true lower: false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: 8, upper: 12, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: false, upperIncluded: false}},
			want: false},
		{name: "2-14 Up and out: container: 0..10, included:upper: false lower: true and unbounded: false,  arg:12..16 included: false, upper unbounded true lower: false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: 12, upper: 16, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: false, upperIncluded: false}},
			want: false},
		// container included false unbounded false, arg included false unbounded low true up false
		{name: "2-15 Down and out container: 0..10, included false and unbounded: false,  arg:-5..-1 included: false, upper unbounded true lower: false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: -5, upper: -1, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: false, upperIncluded: false}},
			want: false},
		{name: "2-16 Down and intersecting: container: 0..10, included false and unbounded: false,  arg:-2..2 included: false, upper unbounded true lower: false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: -2, upper: 2, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: false, upperIncluded: false}},
			want: false},
		{name: "2-17 Down Boundary and In: container: 0..10, included false and unbounded: false,  arg:0..6 included: false, upper unbounded true lower: false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: 0, upper: 6, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: false, upperIncluded: false}},
			want: false},
		{name: "2-18 In: container: 0..10, included:upper: included false and unbounded: false,  arg:2..6 included: false, upper unbounded true lower: false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: 2, upper: 6, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: false, upperIncluded: false}},
			want: false},
		{name: "2-19 Up Boundary and In: container: 0..10, included false and unbounded: false,  arg:8..10 included: false, upper unbounded true lower: false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: 8, upper: 10, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: false, upperIncluded: false}},
			want: false},
		{name: "2-20 Up and intersecting: container: 0..10, included false and unbounded: false,  arg:8..12 included: false, upper unbounded true lower: false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: 8, upper: 12, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: false, upperIncluded: false}},
			want: false},
		{name: "2-21 Up and out: container: 0..10, included false and unbounded: false,  arg:12..16 included: false, upper unbounded true lower: false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: 12, upper: 16, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: false, upperIncluded: false}},
			want: false},
		// container included false unbounded lower true upper false, arg included false unbounded low true up false
		{name: "2-22 Down and out container: 0..10, included false and lower unbounded: true upper unbounded: false,  arg:-5..-1 included: false, upper unbounded true lower: false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: false, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: -5, upper: -1, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: false, upperIncluded: false}},
			want: false},
		{name: "2-23 Down and intersecting: container: 0..10, included false and lower unbounded: true upper unbounded: false,  arg:-2..2 included: false, upper unbounded true lower: false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: false, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: -2, upper: 2, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: false, upperIncluded: false}},
			want: false},
		{name: "2-24 Down Limit and In: container: 0..10, included false and lower unbounded: true upper unbounded: false,  arg:0..6 included: false, upper unbounded true lower: false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: false, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: 0, upper: 6, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: false, upperIncluded: false}},
			want: false},
		{name: "2-25 In: container: 0..10, included:upper: included false and lower unbounded: true upper unbounded: false,  arg:2..6 included: false, upper unbounded true lower: false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: false, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: 2, upper: 6, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: false, upperIncluded: false}},
			want: false},
		{name: "2-26 Up Limit and In: container: 0..10, included false and lower unbounded: true upper unbounded: false,  arg:8..10 included: false, upper unbounded true lower: false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: false, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: 8, upper: 10, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: false, upperIncluded: false}},
			want: false},
		{name: "2-27 Up and intersecting: container: 0..10, included false and lower unbounded: true upper unbounded: false,  arg:8..12 included: false, upper unbounded true lower: false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: false, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: 8, upper: 12, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: false, upperIncluded: false}},
			want: false},
		{name: "2-28 Up and out: container: 0..10, included false and lower unbounded: true upper unbounded: false,  arg:12..16 included: false, upper unbounded true lower: false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: false, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: 12, upper: 16, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: false, upperIncluded: false}},
			want: false},
		// container included false unbounded true, arg included false unbounded low true up false
		{name: "2-29 Down and out container: 0..10, included false and lower unbounded: true upper unbounded: true,  arg:-5..-1 included: false, upper unbounded true lower: false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: false, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: -5, upper: -1, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: false, upperIncluded: false}},
			want: true},
		{name: "2-30 Down and intersecting: container: 0..10, included false and lower unbounded: true upper unbounded: true,  arg:-2..2 included: false, upper unbounded true lower: false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: false, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: -2, upper: 2, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: false, upperIncluded: false}},
			want: true},
		{name: "2-31 Down Limit and In: container: 0..10, included false and lower unbounded: true upper unbounded: true,  arg:0..6 included: false, upper unbounded true lower: false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: false, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: 0, upper: 6, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: false, upperIncluded: false}},
			want: true},
		{name: "2-32 In: container: 0..10, included:upper: included false and lower unbounded: true upper unbounded: true,  arg:2..6 included: false, upper unbounded true lower: false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: false, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: 2, upper: 6, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: false, upperIncluded: false}},
			want: true},
		{name: "2-33 Up Limit and In: container: 0..10, included false and lower unbounded: true upper unbounded: true,  arg:8..10 included: false, upper unbounded true lower: false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: false, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: 8, upper: 10, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: false, upperIncluded: false}},
			want: true},
		{name: "2-34 Up and intersecting: container: 0..10, included false and lower unbounded: true upper unbounded: true,  arg:8..12 included: false, upper unbounded true lower: false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: false, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: 8, upper: 12, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: false, upperIncluded: false}},
			want: true},
		{name: "2-35 Up and out: container: 0..10, included false and lower unbounded: true upper unbounded: true,  arg:12..16 included: false, upper unbounded true lower: false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: false, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: 12, upper: 16, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: false, upperIncluded: false}},
			want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.Contains(tt.args.other); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterval_Has(t *testing.T) {
	type args[T generics.Number] struct {
		value T
	}
	type testCase[T generics.Number] struct {
		name string
		i    Interval[T]
		args args[T]
		want bool
	}
	tests := []testCase[int]{
		{name: "1-1 Value -2, lu false, uu false, li false ui false, w false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false},
			args: args[int]{value: -2},
			want: false},
		{name: "1-2 Value 0, lu false, uu false, li false ui false, w false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false},
			args: args[int]{value: 0},
			want: false},
		{name: "1-3 Value 5, lu false, uu false, li false ui false, w true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false},
			args: args[int]{value: 5},
			want: true},
		{name: "1-4 Value 10, lu false, uu false, li false ui false, w false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false},
			args: args[int]{value: 10},
			want: false},
		{name: "1-5 Value 12, lu false, uu false, li false ui false, w false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false},
			args: args[int]{value: 12},
			want: false},

		{name: "2-1 Value -2, lu true, uu false, li false ui false, w true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: false, upperIncluded: false},
			args: args[int]{value: -2},
			want: true},
		{name: "2-2 Value 0, lu true, uu false, li false ui false, w true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: false, upperIncluded: false},
			args: args[int]{value: 0},
			want: true},
		{name: "2-3 Value 5, lu true, uu false, li false ui false, w true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: false, upperIncluded: false},
			args: args[int]{value: 5},
			want: true},
		{name: "2-4 Value 10, lu true, uu false, li false ui false, w false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: false, upperIncluded: false},
			args: args[int]{value: 10},
			want: false},
		{name: "2-5 Value 12, lu true, uu false, li false ui false, w false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: false, upperIncluded: false},
			args: args[int]{value: 12},
			want: false},

		{name: "3-1 Value -2, lu false, uu true, li false ui false, w false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: false, upperIncluded: false},
			args: args[int]{value: -2},
			want: false},
		{name: "3-2 Value 0, lu false, uu true, li false ui false, w false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: false, upperIncluded: false},
			args: args[int]{value: 0},
			want: false},
		{name: "3-3 Value 5, lu false, uu true, li false ui false, w true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: false, upperIncluded: false},
			args: args[int]{value: 5},
			want: true},
		{name: "3-4 Value 10, lu false, uu true, li false ui false, w true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: false, upperIncluded: false},
			args: args[int]{value: 10},
			want: true},
		{name: "3-5 Value 12, lu false, uu true, li false ui false, w true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: false, upperIncluded: false},
			args: args[int]{value: 12},
			want: true},

		{name: "4-1 Value -2, lu true, uu true, li false ui false, w true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: false, upperIncluded: false},
			args: args[int]{value: -2},
			want: true},
		{name: "4-2 Value 0, lu true, uu true, li false ui false, w true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: false, upperIncluded: false},
			args: args[int]{value: 0},
			want: true},
		{name: "4-3 Value 5, lu true, uu true, li false ui false, w true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: false, upperIncluded: false},
			args: args[int]{value: 5},
			want: true},
		{name: "4-4 Value 10, lu true, uu true, li false ui false, w true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: false, upperIncluded: false},
			args: args[int]{value: 10},
			want: true},
		{name: "4-5 Value 12, lu true, uu true, li false ui false, w true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: false, upperIncluded: false},
			args: args[int]{value: 12},
			want: true},

		{name: "5-1 Value -2, lu false, uu false, li true ui false, w false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: false},
			args: args[int]{value: -2},
			want: false},
		{name: "5-2 Value 0, lu false, uu false, li true ui false, w true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: false},
			args: args[int]{value: 0},
			want: true},
		{name: "5-3 Value 5, lu false, uu false, li true ui false, w true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: false},
			args: args[int]{value: 5},
			want: true},
		{name: "5-4 Value 10, lu false, uu false, li true ui false, w false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: false},
			args: args[int]{value: 10},
			want: false},
		{name: "5-5 Value 12, lu false, uu false, li true ui false, w false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: false},
			args: args[int]{value: 12},
			want: false},

		{name: "6-1 Value -2, lu false, uu false, li false ui false, w false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: true},
			args: args[int]{value: -2},
			want: false},
		{name: "6-2 Value 0, lu false, uu false, li true ui false, w false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: true},
			args: args[int]{value: 0},
			want: false},
		{name: "6-3 Value 5, lu false, uu false, li true ui false, w true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: true},
			args: args[int]{value: 5},
			want: true},
		{name: "6-4 Value 10, lu false, uu false, li true ui false, w true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: true},
			args: args[int]{value: 10},
			want: true},
		{name: "6-5 Value 12, lu false, uu false, li true ui false, w false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: true},
			args: args[int]{value: 12},
			want: false},

		{name: "7-1 Value -2, lu false, uu false, li true ui true, w false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{value: -2},
			want: false},
		{name: "7-2 Value 0, lu false, uu false, li true ui true, w true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{value: 0},
			want: true},
		{name: "7-3 Value 5, lu false, uu false, li true ui true, w true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{value: 5},
			want: true},
		{name: "7-4 Value 10, lu false, uu false, li true ui true, w true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{value: 10},
			want: true},
		{name: "7-5 Value 12, lu false, uu false, li true ui true, w false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{value: 12},
			want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.Has(tt.args.value); got != tt.want {
				t.Errorf("Has() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterval_Intersects(t *testing.T) {
	type args[T generics.Number] struct {
		other *Interval[T]
	}
	type testCase[T generics.Number] struct {
		name string
		i    Interval[T]
		args args[T]
		want bool
	}
	// false false false false 0
	// true false false false 1
	// false true false false 2
	// true true false false 3

	// false false true false 4
	// true false true false 5
	// false true true false 6
	// true true true false 7

	// false false false true 8
	// true false false true 9
	// false true false true 10
	// true true false true 11

	// false false true true 12
	// true false true true 13
	// false true true true 14
	// true true true true 15

	tests := []testCase[int]{
		//false false false false
		{name: "0-1 Down and out : 0..10, lu false uu false, arg:-5..-1 lu false uu false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: -1, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: false},
		{name: "0-2 Down and Boundary : 0..10, lu false uu false, arg:-5..0 lu false uu false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: 0, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: false},
		{name: "0-3 Down and Intersects : lu false uu false, arg:-5..5 lu false uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: 5, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "0-4 In : 0..10, lu false uu false, arg:3..7 lu false uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 3, upper: 7, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "0-5 Up and Intersects : 0..10, lu false uu false, arg:7..12 lu false uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 7, upper: 12, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "0-6 Boundary and Up : 0..10, lu false uu false, arg:10..12 lu false uu false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 10, upper: 12, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: false},
		{name: "0-7 Up and Out: 0..10, lu false uu false, arg:12..15 lu false uu false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 12, upper: 15, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: false},
		// true false false false
		{name: "1-1 Down and out : 0..10, lu true uu false, arg:-5..-1 lu false uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: -1, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "1-2 Down and Boundary : 0..10, lu true uu false, arg:-5..0 lu false uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: 0, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "1-3 Down and Intersects : lu true uu false, arg:-5..5 lu false uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: 5, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "1-4 In : 0..10, lu true uu false, arg:3..7 lu false uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 3, upper: 7, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "1-5 Up and Intersects : 0..10, lu true uu false, arg:7..12 lu false uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 7, upper: 12, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "1-6 Boundary and Up : 0..10, lu true uu false, arg:10..12 lu false uu false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 10, upper: 12, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: false},
		{name: "1-7 Up and Out: 0..10, lu true uu false, arg:12..15 lu false uu false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 12, upper: 15, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: false},
		// false true false false
		{name: "2-1 Down and out : 0..10, lu false uu true, arg:-5..-1 lu false uu false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: -1, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: false},
		{name: "2-2 Down and Boundary : 0..10, lu false uu true, arg:-5..0 lu false uu false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: 0, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: false},
		{name: "2-3 Down and Intersects : lu false uu true, arg:-5..5 lu false uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: 5, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "2-4 In : 0..10, lu false uu true, arg:3..7 lu false uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 3, upper: 7, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "2-5 Up and Intersects : 0..10, lu false uu true, arg:7..12 lu false uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 7, upper: 12, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "2-6 Boundary and Up : 0..10, lu false uu true, arg:10..12 lu false uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 10, upper: 12, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "2-7 Up and Out: 0..10, lu false uu true, arg:12..15 lu false uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 12, upper: 15, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: true},
		// true true false false
		{name: "3-1 Down and out : 0..10, lu true uu true, arg:-5..-1 lu false uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: -1, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "3-2 Down and Boundary : 0..10, lu true uu true, arg:-5..0 lu false uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: 0, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "3-3 Down and Intersects : lu true uu true, arg:-5..5 lu false uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: 5, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "3-4 In : 0..10, lu true uu true, arg:3..7 lu false uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 3, upper: 7, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "3-5 Up and Intersects : 0..10, lu true uu true, arg:7..12 lu false uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 7, upper: 12, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "3-6 Boundary and Up : 0..10, lu true uu true, arg:10..12 lu false uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 10, upper: 12, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "3-7 Up and Out: 0..10, lu true uu true, arg:12..15 lu false uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 12, upper: 15, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: true},
		// ===============================
		// false false true false
		{name: "4-1 Down and out : 0..10, lu false uu false, arg:-5..-1 lu true uu false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: -1, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: false},
		{name: "4-2 Down and Boundary : 0..10, lu false uu false, arg:-5..0 lu true uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: 0, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "4-3 Down and Intersects : lu false uu false, arg:-5..5 lu true uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: 5, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "4-4 In : 0..10, lu false uu false, arg:3..7 lu true uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 3, upper: 7, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "4-5 Up and Intersects : 0..10, lu false uu false, arg:7..12 lu true uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 7, upper: 12, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "4-6 Boundary and Up : 0..10, lu false uu false, arg:10..12 lu true uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 10, upper: 12, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "4-7 Up and Out: 0..10, lu false uu false, arg:12..15 lu true uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 12, upper: 15, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: true},
		// true false true false
		{name: "5-1 Down and out : 0..10, lu true uu false, arg:-5..-1 lu true uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: -1, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "5-2 Down and Boundary : 0..10, lu true uu false, arg:-5..0 lu true uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: 0, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "5-3 Down and Intersects : lu true uu false, arg:-5..5 lu true uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: 5, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "5-4 In : 0..10, lu true uu false, arg:3..7 lu true uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 3, upper: 7, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "5-5 Up and Intersects : 0..10, lu true uu false, arg:7..12 lu true uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 7, upper: 12, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "5-6 Boundary and Up : 0..10, lu true uu false, arg:10..12 lu true uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 10, upper: 12, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "5-7 Up and Out: 0..10, lu true uu false, arg:12..15 lu true uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 12, upper: 15, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: true},
		// false true true false
		{name: "6-1 Down and out : 0..10, lu false uu true, arg:-5..-1 lu true uu false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: -1, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: false},
		{name: "6-2 Down and Boundary : 0..10, lu false uu true, arg:-5..0 lu true uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: 0, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "6-3 Down and Intersects : lu false uu true, arg:-5..5 lu true uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: 5, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "6-4 In : 0..10, lu false uu true, arg:3..7 lu true uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 3, upper: 7, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "6-5 Up and Intersects : 0..10, lu false uu true, arg:7..12 lu true uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 7, upper: 12, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "6-6 Boundary and Up : 0..10, lu false uu true, arg:10..12 lu true uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 10, upper: 12, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "6-7 Up and Out: 0..10, lu false uu true, arg:12..15 lu true uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 12, upper: 15, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: true},
		// true true true false
		{name: "7-1 Down and out : 0..10, lu true uu true, arg:-5..-1 lu true uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: -1, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "7-2 Down and Boundary : 0..10, lu true uu true, arg:-5..0 lu true uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: 0, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "7-3 Down and Intersects : lu true uu true, arg:-5..5 lu true uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: 5, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "7-4 In : 0..10, lu true uu true, arg:3..7 lu true uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 3, upper: 7, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "7-5 Up and Intersects : 0..10, lu true uu true, arg:7..12 lu true uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 7, upper: 12, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "7-6 Boundary and Up : 0..10, lu true uu true, arg:10..12 lu true uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 10, upper: 12, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "7-7 Up and Out: 0..10, lu true uu true, arg:12..15 lu true uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 12, upper: 15, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: true},
		// ===============================
		// false false false true
		{name: "8-1 Down and out : 0..10, lu false uu false, arg:-5..-1 lu false uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: -1, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "8-2 Down and Boundary : 0..10, lu false uu false, arg:-5..0 lu false uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: 0, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "8-3 Down and Intersects : lu false uu false, arg:-5..5 lu false uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: 5, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "8-4 In : 0..10, lu false uu false, arg:3..7 lu false uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 3, upper: 7, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "8-5 Up and Intersects : 0..10, lu false uu false, arg:7..12 lu false uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 7, upper: 12, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "8-6 Boundary and Up : 0..10, lu false uu false, arg:10..12 lu false uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 10, upper: 12, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "8-7 Up and Out: 0..10, lu false uu false, arg:12..15 lu false uu true, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 12, upper: 15, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: false},
		// true false false true
		{name: "9-1 Down and out : 0..10, lu true uu false, arg:-5..-1 lu false uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: -1, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "9-2 Down and Boundary : 0..10, lu true uu false, arg:-5..0 lu false uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: 0, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "9-3 Down and Intersects : lu true uu false, arg:-5..5 lu false uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: 5, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "9-4 In : 0..10, lu true uu false, arg:3..7 lu false uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 3, upper: 7, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "9-5 Up and Intersects : 0..10, lu true uu false, arg:7..12 lu false uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 7, upper: 12, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "9-6 Boundary and Up : 0..10, lu true uu false, arg:10..12 lu false uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 10, upper: 12, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "9-7 Up and Out: 0..10, lu true uu false, arg:12..15 lu false uu true, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 12, upper: 15, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: false},
		// false true false true
		{name: "10-1 Down and out : 0..10, lu false uu true, arg:-5..-1 lu false uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: -1, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "10-2 Down and Boundary : 0..10, lu false uu true, arg:-5..0 lu false uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: 0, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "10-3 Down and Intersects : lu false uu true, arg:-5..5 lu false uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: 5, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "10-4 In : 0..10, lu false uu true, arg:3..7 lu false uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 3, upper: 7, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "10-5 Up and Intersects : 0..10, lu false uu true, arg:7..12 lu false uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 7, upper: 12, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "10-6 Boundary and Up : 0..10, lu false uu true, arg:10..12 lu false uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 10, upper: 12, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "10-7 Up and Out: 0..10, lu false uu true, arg:12..15 lu false uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 12, upper: 15, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: true},
		// true true false true
		{name: "11-1 Down and out : 0..10, lu true uu true, arg:-5..-1 lu false uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: -1, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "11-2 Down and Boundary : 0..10, lu true uu true, arg:-5..0 lu false uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: 0, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "11-3 Down and Intersects : lu true uu true, arg:-5..5 lu false uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: 5, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "11-4 In : 0..10, lu true uu true, arg:3..7 lu false uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 3, upper: 7, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "11-5 Up and Intersects : 0..10, lu true uu true, arg:7..12 lu false uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 7, upper: 12, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "11-6 Boundary and Up : 0..10, lu true uu true, arg:10..12 lu false uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 10, upper: 12, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "11-7 Up and Out: 0..10, lu true uu true, arg:12..15 lu false uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 12, upper: 15, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: true},
		// ===============================
		// false false true true

		{name: "12-1 Down and out : 0..10, lu false uu false, arg:-5..-1 lu true uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: -1, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "12-2 Down and Boundary : 0..10, lu false uu false, arg:-5..0 lu true uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: 0, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "12-3 Down and Intersects : lu false uu false, arg:-5..5 lu true uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: 5, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "12-4 In : 0..10, lu false uu false, arg:3..7 lu true uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 3, upper: 7, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "12-5 Up and Intersects : 0..10, lu false uu false, arg:7..12 lu true uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 7, upper: 12, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "12-6 Boundary and Up : 0..10, lu false uu false, arg:10..12 lu true uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 10, upper: 12, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "12-7 Up and Out: 0..10, lu false uu false, arg:12..15 lu true uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 12, upper: 15, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: true},
		// true false true true
		{name: "13-1 Down and out : 0..10, lu true uu false, arg:-5..-1 lu true uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: -1, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "13-2 Down and Boundary : 0..10, lu true uu false, arg:-5..0 lu true uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: 0, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "13-3 Down and Intersects : lu true uu false, arg:-5..5 lu true uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: 5, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "13-4 In : 0..10, lu true uu false, arg:3..7 lu true uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 3, upper: 7, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "13-5 Up and Intersects : 0..10, lu true uu false, arg:7..12 lu true uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 7, upper: 12, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "13-6 Boundary and Up : 0..10, lu true uu false, arg:10..12 lu true uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 10, upper: 12, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "13-7 Up and Out: 0..10, lu true uu false, arg:12..15 lu true uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 12, upper: 15, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: true},
		// false true true true
		{name: "14-1 Down and out : 0..10, lu false uu true, arg:-5..-1 lu true uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: -1, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "14-2 Down and Boundary : 0..10, lu false uu true, arg:-5..0 lu true uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: 0, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "14-3 Down and Intersects : lu false uu true, arg:-5..5 lu true uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: 5, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "14-4 In : 0..10, lu false uu true, arg:3..7 lu true uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 3, upper: 7, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "14-5 Up and Intersects : 0..10, lu false uu true, arg:7..12 lu true uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 7, upper: 12, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "14-6 Boundary and Up : 0..10, lu false uu true, arg:10..12 lu true uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 10, upper: 12, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "14-7 Up and Out: 0..10, lu false uu true, arg:12..15 lu true uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 12, upper: 15, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: true},
		// true true true true
		{name: "15-1 Down and out : 0..10, lu true uu true, arg:-5..-1 lu true uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: -1, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "15-2 Down and Boundary : 0..10, lu true uu true, arg:-5..0 lu true uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: 0, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "15-3 Down and Intersects : lu true uu true, arg:-5..5 lu true uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: 5, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "15-4 In : 0..10, lu true uu true, arg:3..7 lu true uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 3, upper: 7, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "15-5 Up and Intersects : 0..10, lu true uu true, arg:7..12 lu true uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 7, upper: 12, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "15-6 Boundary and Up : 0..10, lu true uu true, arg:10..12 lu true uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 10, upper: 12, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "15-7 Up and Out: 0..10, lu true uu true, arg:12..15 lu true uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 12, upper: 15, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: true},
		// ===============================
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.Intersects(tt.args.other); got != tt.want {
				t.Errorf("Intersects() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterval_IntersectsAsInterval(t *testing.T) {
	type args[T generics.Number] struct {
		other *Interval[T]
	}
	type testCase[T generics.Number] struct {
		name string
		i    Interval[T]
		args args[T]
		want *Interval[T]
	}
	tests := []testCase[int]{
		//false false false false
		{name: "0-1 Down and out : 0..10, lu false uu false, arg:-5..-1 lu false uu false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: -1, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "0-2 Down and Boundary : 0..10, lu false uu false, arg:-5..0 lu false uu false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: 0, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "0-3 Down and Intersects : lu false uu false, arg:-5..5 lu false uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: 5, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: &Interval[int]{0, 5, false, false, true, true}},
		{name: "0-4 In : 0..10, lu false uu false, arg:3..7 lu false uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 3, upper: 7, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: &Interval[int]{3, 7, false, false, true, true}},
		{name: "0-5 Up and Intersects : 0..10, lu false uu false, arg:7..12 lu false uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 7, upper: 12, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "0-6 Boundary and Up : 0..10, lu false uu false, arg:10..12 lu false uu false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 10, upper: 12, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "0-7 Up and Out: 0..10, lu false uu false, arg:12..15 lu false uu false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 12, upper: 15, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		// true false false false
		{name: "1-1 Down and out : 0..10, lu true uu false, arg:-5..-1 lu false uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: -1, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "1-2 Down and Boundary : 0..10, lu true uu false, arg:-5..0 lu false uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: 0, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "1-3 Down and Intersects : lu true uu false, arg:-5..5 lu false uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: 5, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "1-4 In : 0..10, lu true uu false, arg:3..7 lu false uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 3, upper: 7, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "1-5 Up and Intersects : 0..10, lu true uu false, arg:7..12 lu false uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 7, upper: 12, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "1-6 Boundary and Up : 0..10, lu true uu false, arg:10..12 lu false uu false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 10, upper: 12, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "1-7 Up and Out: 0..10, lu true uu false, arg:12..15 lu false uu false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 12, upper: 15, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		// false true false false
		{name: "2-1 Down and out : 0..10, lu false uu true, arg:-5..-1 lu false uu false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: -1, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "2-2 Down and Boundary : 0..10, lu false uu true, arg:-5..0 lu false uu false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: 0, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "2-3 Down and Intersects : lu false uu true, arg:-5..5 lu false uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: 5, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "2-4 In : 0..10, lu false uu true, arg:3..7 lu false uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 3, upper: 7, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "2-5 Up and Intersects : 0..10, lu false uu true, arg:7..12 lu false uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 7, upper: 12, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "2-6 Boundary and Up : 0..10, lu false uu true, arg:10..12 lu false uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 10, upper: 12, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "2-7 Up and Out: 0..10, lu false uu true, arg:12..15 lu false uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 12, upper: 15, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		// true true false false
		{name: "3-1 Down and out : 0..10, lu true uu true, arg:-5..-1 lu false uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: -1, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "3-2 Down and Boundary : 0..10, lu true uu true, arg:-5..0 lu false uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: 0, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "3-3 Down and Intersects : lu true uu true, arg:-5..5 lu false uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: 5, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "3-4 In : 0..10, lu true uu true, arg:3..7 lu false uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 3, upper: 7, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "3-5 Up and Intersects : 0..10, lu true uu true, arg:7..12 lu false uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 7, upper: 12, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "3-6 Boundary and Up : 0..10, lu true uu true, arg:10..12 lu false uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 10, upper: 12, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "3-7 Up and Out: 0..10, lu true uu true, arg:12..15 lu false uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 12, upper: 15, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		// ===============================
		// false false true false
		{name: "4-1 Down and out : 0..10, lu false uu false, arg:-5..-1 lu true uu false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: -1, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "4-2 Down and Boundary : 0..10, lu false uu false, arg:-5..0 lu true uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: 0, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "4-3 Down and Intersects : lu false uu false, arg:-5..5 lu true uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: 5, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "4-4 In : 0..10, lu false uu false, arg:3..7 lu true uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 3, upper: 7, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "4-5 Up and Intersects : 0..10, lu false uu false, arg:7..12 lu true uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 7, upper: 12, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "4-6 Boundary and Up : 0..10, lu false uu false, arg:10..12 lu true uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 10, upper: 12, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "4-7 Up and Out: 0..10, lu false uu false, arg:12..15 lu true uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 12, upper: 15, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		// true false true false
		{name: "5-1 Down and out : 0..10, lu true uu false, arg:-5..-1 lu true uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: -1, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "5-2 Down and Boundary : 0..10, lu true uu false, arg:-5..0 lu true uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: 0, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "5-3 Down and Intersects : lu true uu false, arg:-5..5 lu true uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: 5, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "5-4 In : 0..10, lu true uu false, arg:3..7 lu true uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 3, upper: 7, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "5-5 Up and Intersects : 0..10, lu true uu false, arg:7..12 lu true uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 7, upper: 12, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "5-6 Boundary and Up : 0..10, lu true uu false, arg:10..12 lu true uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 10, upper: 12, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "5-7 Up and Out: 0..10, lu true uu false, arg:12..15 lu true uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 12, upper: 15, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		// false true true false
		{name: "6-1 Down and out : 0..10, lu false uu true, arg:-5..-1 lu true uu false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: -1, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "6-2 Down and Boundary : 0..10, lu false uu true, arg:-5..0 lu true uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: 0, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "6-3 Down and Intersects : lu false uu true, arg:-5..5 lu true uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: 5, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "6-4 In : 0..10, lu false uu true, arg:3..7 lu true uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 3, upper: 7, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "6-5 Up and Intersects : 0..10, lu false uu true, arg:7..12 lu true uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 7, upper: 12, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "6-6 Boundary and Up : 0..10, lu false uu true, arg:10..12 lu true uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 10, upper: 12, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "6-7 Up and Out: 0..10, lu false uu true, arg:12..15 lu true uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 12, upper: 15, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		// true true true false
		{name: "7-1 Down and out : 0..10, lu true uu true, arg:-5..-1 lu true uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: -1, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "7-2 Down and Boundary : 0..10, lu true uu true, arg:-5..0 lu true uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: 0, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "7-3 Down and Intersects : lu true uu true, arg:-5..5 lu true uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: 5, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "7-4 In : 0..10, lu true uu true, arg:3..7 lu true uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 3, upper: 7, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "7-5 Up and Intersects : 0..10, lu true uu true, arg:7..12 lu true uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 7, upper: 12, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "7-6 Boundary and Up : 0..10, lu true uu true, arg:10..12 lu true uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 10, upper: 12, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "7-7 Up and Out: 0..10, lu true uu true, arg:12..15 lu true uu false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 12, upper: 15, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		// ===============================
		// false false false true
		{name: "8-1 Down and out : 0..10, lu false uu false, arg:-5..-1 lu false uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: -1, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "8-2 Down and Boundary : 0..10, lu false uu false, arg:-5..0 lu false uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: 0, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "8-3 Down and Intersects : lu false uu false, arg:-5..5 lu false uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: 5, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "8-4 In : 0..10, lu false uu false, arg:3..7 lu false uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 3, upper: 7, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "8-5 Up and Intersects : 0..10, lu false uu false, arg:7..12 lu false uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 7, upper: 12, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "8-6 Boundary and Up : 0..10, lu false uu false, arg:10..12 lu false uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 10, upper: 12, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "8-7 Up and Out: 0..10, lu false uu false, arg:12..15 lu false uu true, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 12, upper: 15, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		// true false false true
		{name: "9-1 Down and out : 0..10, lu true uu false, arg:-5..-1 lu false uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: -1, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "9-2 Down and Boundary : 0..10, lu true uu false, arg:-5..0 lu false uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: 0, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "9-3 Down and Intersects : lu true uu false, arg:-5..5 lu false uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: 5, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "9-4 In : 0..10, lu true uu false, arg:3..7 lu false uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 3, upper: 7, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "9-5 Up and Intersects : 0..10, lu true uu false, arg:7..12 lu false uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 7, upper: 12, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "9-6 Boundary and Up : 0..10, lu true uu false, arg:10..12 lu false uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 10, upper: 12, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "9-7 Up and Out: 0..10, lu true uu false, arg:12..15 lu false uu true, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 12, upper: 15, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		// false true false true
		{name: "10-1 Down and out : 0..10, lu false uu true, arg:-5..-1 lu false uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: -1, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "10-2 Down and Boundary : 0..10, lu false uu true, arg:-5..0 lu false uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: 0, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "10-3 Down and Intersects : lu false uu true, arg:-5..5 lu false uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: 5, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "10-4 In : 0..10, lu false uu true, arg:3..7 lu false uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 3, upper: 7, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "10-5 Up and Intersects : 0..10, lu false uu true, arg:7..12 lu false uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 7, upper: 12, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "10-6 Boundary and Up : 0..10, lu false uu true, arg:10..12 lu false uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 10, upper: 12, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "10-7 Up and Out: 0..10, lu false uu true, arg:12..15 lu false uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 12, upper: 15, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		// true true false true
		{name: "11-1 Down and out : 0..10, lu true uu true, arg:-5..-1 lu false uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: -1, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "11-2 Down and Boundary : 0..10, lu true uu true, arg:-5..0 lu false uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: 0, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "11-3 Down and Intersects : lu true uu true, arg:-5..5 lu false uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: 5, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "11-4 In : 0..10, lu true uu true, arg:3..7 lu false uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 3, upper: 7, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "11-5 Up and Intersects : 0..10, lu true uu true, arg:7..12 lu false uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 7, upper: 12, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "11-6 Boundary and Up : 0..10, lu true uu true, arg:10..12 lu false uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 10, upper: 12, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "11-7 Up and Out: 0..10, lu true uu true, arg:12..15 lu false uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 12, upper: 15, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		// ===============================
		// false false true true

		{name: "12-1 Down and out : 0..10, lu false uu false, arg:-5..-1 lu true uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: -1, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "12-2 Down and Boundary : 0..10, lu false uu false, arg:-5..0 lu true uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: 0, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "12-3 Down and Intersects : lu false uu false, arg:-5..5 lu true uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: 5, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "12-4 In : 0..10, lu false uu false, arg:3..7 lu true uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 3, upper: 7, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "12-5 Up and Intersects : 0..10, lu false uu false, arg:7..12 lu true uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 7, upper: 12, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "12-6 Boundary and Up : 0..10, lu false uu false, arg:10..12 lu true uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 10, upper: 12, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "12-7 Up and Out: 0..10, lu false uu false, arg:12..15 lu true uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 12, upper: 15, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		// true false true true
		{name: "13-1 Down and out : 0..10, lu true uu false, arg:-5..-1 lu true uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: -1, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "13-2 Down and Boundary : 0..10, lu true uu false, arg:-5..0 lu true uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: 0, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "13-3 Down and Intersects : lu true uu false, arg:-5..5 lu true uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: 5, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "13-4 In : 0..10, lu true uu false, arg:3..7 lu true uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 3, upper: 7, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "13-5 Up and Intersects : 0..10, lu true uu false, arg:7..12 lu true uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 7, upper: 12, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "13-6 Boundary and Up : 0..10, lu true uu false, arg:10..12 lu true uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 10, upper: 12, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "13-7 Up and Out: 0..10, lu true uu false, arg:12..15 lu true uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 12, upper: 15, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		// false true true true
		{name: "14-1 Down and out : 0..10, lu false uu true, arg:-5..-1 lu true uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: -1, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "14-2 Down and Boundary : 0..10, lu false uu true, arg:-5..0 lu true uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: 0, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "14-3 Down and Intersects : lu false uu true, arg:-5..5 lu true uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: 5, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "14-4 In : 0..10, lu false uu true, arg:3..7 lu true uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 3, upper: 7, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "14-5 Up and Intersects : 0..10, lu false uu true, arg:7..12 lu true uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 7, upper: 12, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "14-6 Boundary and Up : 0..10, lu false uu true, arg:10..12 lu true uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 10, upper: 12, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "14-7 Up and Out: 0..10, lu false uu true, arg:12..15 lu true uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 12, upper: 15, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		// true true true true
		{name: "15-1 Down and out : 0..10, lu true uu true, arg:-5..-1 lu true uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: -1, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "15-2 Down and Boundary : 0..10, lu true uu true, arg:-5..0 lu true uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: 0, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "15-3 Down and Intersects : lu true uu true, arg:-5..5 lu true uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: 5, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "15-4 In : 0..10, lu true uu true, arg:3..7 lu true uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 3, upper: 7, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "15-5 Up and Intersects : 0..10, lu true uu true, arg:7..12 lu true uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 7, upper: 12, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "15-6 Boundary and Up : 0..10, lu true uu true, arg:10..12 lu true uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 10, upper: 12, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		{name: "15-7 Up and Out: 0..10, lu true uu true, arg:12..15 lu true uu true, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 12, upper: 15, lowerUnbounded: true, upperUnbounded: true, lowerIncluded: true, upperIncluded: true}},
			want: nil},
		// ===============================
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.IntersectsAsInterVal(tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Intersects() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLowerUnboundedInterval(t *testing.T) {
	type args[T generics.Number] struct {
		upper         T
		UpperIncluded bool
	}
	type testCase[T generics.Number] struct {
		name string
		args args[T]
		want *Interval[T]
	}
	tests := []testCase[int]{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := LowerUnboundedInterval(tt.args.upper, tt.args.UpperIncluded); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LowerUnboundedInterval() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultiplicityInterval_IsMandatory(t *testing.T) {
	type testCase[T generics.Number] struct {
		name string
		mi   MultiplicityInterval[T]
		want bool
	}
	tests := []testCase[int]{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.mi.IsMandatory(); got != tt.want {
				t.Errorf("IsMandatory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultiplicityInterval_IsOpen(t *testing.T) {
	type testCase[T generics.Number] struct {
		name string
		mi   MultiplicityInterval[T]
		want bool
	}
	tests := []testCase[int]{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.mi.IsOpen(); got != tt.want {
				t.Errorf("IsOpen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultiplicityInterval_IsOptional(t *testing.T) {
	type testCase[T generics.Number] struct {
		name string
		mi   MultiplicityInterval[T]
		want bool
	}
	tests := []testCase[int]{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.mi.IsOptional(); got != tt.want {
				t.Errorf("IsOptional() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultiplicityInterval_IsProhibited(t *testing.T) {
	type testCase[T generics.Number] struct {
		name string
		mi   MultiplicityInterval[T]
		want bool
	}
	tests := []testCase[int]{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.mi.IsProhibited(); got != tt.want {
				t.Errorf("IsProhibited() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewInterval(t *testing.T) {
	type testCase[T generics.Number] struct {
		name string
		want *Interval[T]
	}
	tests := []testCase[int]{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewInterval[int](); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInterval() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewIntervalBuilder(t *testing.T) {
	type testCase[T generics.Number] struct {
		name string
		want *IntervalBuilder[T]
	}
	tests := []testCase[int]{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewIntervalBuilder[int](); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIntervalBuilder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewMultiplicityInterval(t *testing.T) {
	type testCase[T generics.Number] struct {
		name string
		want *MultiplicityInterval[T]
	}
	tests := []testCase[int]{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMultiplicityInterval[int](); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMultiplicityInterval() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnboundedInterval(t *testing.T) {
	type args[T generics.Number] struct {
		lower T
		upper T
	}
	type testCase[T generics.Number] struct {
		name string
		args args[T]
		want *Interval[T]
	}
	tests := []testCase[int]{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := UnboundedInterval(tt.args.lower, tt.args.upper); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnboundedInterval() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpperUnboundedInterval(t *testing.T) {
	type args[T generics.Number] struct {
		lower         T
		LowerIncluded bool
	}
	type testCase[T generics.Number] struct {
		name string
		args args[T]
		want *Interval[T]
	}
	tests := []testCase[int]{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := UpperUnboundedInterval(tt.args.lower, tt.args.LowerIncluded); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpperUnboundedInterval() = %v, want %v", got, tt.want)
			}
		})
	}
}
