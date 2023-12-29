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
		{name: "0-1 Down and out container: 0..10, included: true and unbounded: false,  arg:-5..-1 included: false, unbounded: false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -5, upper: -1, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: false},
		{name: "0-2 Down and intersecting: container: 0..10, included: true and unbounded: false,  arg:-2..2 included: false, unbounded: false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -2, upper: 2, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: false},
		{name: "0-3 Down Limit and In: container: 0..10, included: true and unbounded: false,  arg:0..6 included: false, unbounded: false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 0, upper: 6, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: true},
		{name: "0-4 In: container: 0..10, included: true and unbounded: false,  arg:2..6 included: false, unbounded: false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 2, upper: 6, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: true},
		{name: "0-5 Up Limit and In: container: 0..10, included: true and unbounded: false,  arg:8..10 included: false, unbounded: false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 0, upper: 6, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: true},
		{name: "0-6 Up and intersecting: container: 0..10, included: true and unbounded: false,  arg:8..12 included: false, unbounded: false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 8, upper: 12, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: false},
		{name: "0-7 Up and out: container: 0..10, included: true and unbounded: false,  arg:12..16 included: false, unbounded: false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 12, upper: 16, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: false},
		// container included upper false included lower true unbounded false, arg included false unbounded false
		{name: "0-8 Down and out container: 0..10, included:upper: false lower: true and unbounded: false,  arg:-5..-1 included: false, unbounded: false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: -5, upper: -1, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: false},
		{name: "0-9 Down and intersecting: container: 0..10, included:upper: false lower: true and unbounded: false,  arg:-2..2 included: false, unbounded: false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: -2, upper: 2, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: false},
		{name: "0-10 Down Limit and In: container: 0..10, included:upper: false lower: true and unbounded: false,  arg:0..6 included: false, unbounded: false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: 0, upper: 6, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: true},
		{name: "0-11 In: container: 0..10, included:upper: false lower: true and unbounded: false,  arg:2..6 included: false, unbounded: false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: 2, upper: 6, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: true},
		{name: "0-12 Up Limit and In: container: 0..10, included:upper: false lower: true and unbounded: false,  arg:8..10 included: false, unbounded: false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: 8, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: false},
		{name: "0-13 Up and intersecting: container: 0..10, included:upper: false lower: true and unbounded: false,  arg:8..12 included: false, unbounded: false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: 8, upper: 12, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: false},
		{name: "0-14 Up and out: container: 0..10, included:upper: false lower: true and unbounded: false,  arg:12..16 included: false, unbounded: false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: 12, upper: 16, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: false},
		// container included false unbounded false, arg included false unbounded false
		{name: "0-15 Down and out container: 0..10, included false and unbounded: false,  arg:-5..-1 included: false, unbounded: false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: -5, upper: -1, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: false},
		{name: "0-16 Down and intersecting: container: 0..10, included false and unbounded: false,  arg:-2..2 included: false, unbounded: false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: -2, upper: 2, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: false},
		{name: "0-17 Down Boundary and In: container: 0..10, included false and unbounded: false,  arg:0..6 included: false, unbounded: false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: 0, upper: 6, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: false},
		{name: "0-18 In: container: 0..10, included:upper: included false and unbounded: false,  arg:2..6 included: false, unbounded: false, result: true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: 2, upper: 6, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: true},
		{name: "0-19 Up Boundary and In: container: 0..10, included false and unbounded: false,  arg:8..10 included: false, unbounded: false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: 8, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: false},
		{name: "0-20 Up and intersecting: container: 0..10, included false and unbounded: false,  arg:8..12 included: false, unbounded: false, result: false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: 8, upper: 12, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false}},
			want: false},
		{name: "0-21 Up and out: container: 0..10, included false and unbounded: false,  arg:12..16 included: false, unbounded: false, result: false",
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
	tests := []testCase[int /* TODO: Insert concrete types here */]{
		{name: "Value 1, unbounded false, included false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false},
			args: args[int]{value: 1},
			want: true},
		{name: "Value -1, unbounded false, included false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false},
			args: args[int]{value: -1},
			want: false},
		{name: "Value 10, unbounded false, included false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false},
			args: args[int]{value: 10},
			want: false},
		{name: "Value 0, unbounded false, included false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false},
			args: args[int]{value: 10},
			want: false},
		{name: "Value 11, unbounded false, included false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: false},
			args: args[int]{value: 11},
			want: false},
		{name: "Value 10, unbounded false, included true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{value: 10},
			want: true},
		{name: "Value 0, unbounded false, included true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{value: 10},
			want: true},
		{name: "Value -1, lower-unbounded true, included true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{value: -1},
			want: true},
		{name: "Value 11, lower-unbounded true, included true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: true, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{value: 11},
			want: false},
		{name: "Value -1, upper-unbounded true, included true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{value: -1},
			want: false},
		{name: "Value 11, upper-unbounded true, included true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: true, upperIncluded: true},
			args: args[int]{value: 11},
			want: true},
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
	tests := []testCase[int /* TODO: Insert concrete types here */]{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.Intersects(tt.args.other); got != tt.want {
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
