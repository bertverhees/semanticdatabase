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
			i: *NewIntervalBuilder[int]().setLower(0).setUpper(1).setUpperIncluded(true).setLowerIncluded(true),
			want: &Interval[int]{
				lower:          0,
				upper:          1,
				lowerUnbounded: false,
				upperUnbounded: false,
				lowerIncluded:  true,
				upperIncluded:  true,
			}},
		{name: "Test for setted Unbounded booleans",
			i: *NewIntervalBuilder[int]().setLower(0).setUpper(1).setLowerUnbounded(true).setUpperUnbounded(true),
			want: &Interval[int]{
				lower:          0,
				upper:          1,
				lowerUnbounded: true,
				upperUnbounded: true,
				lowerIncluded:  false,
				upperIncluded:  false,
			}},
		{name: "Test for setted Included and unbounded booleans",
			i: *NewIntervalBuilder[int]().setLower(0).setUpper(1).setLowerIncluded(true).setUpperIncluded(true).setUpperUnbounded(true).setLowerUnbounded(true),
			want: &Interval[int]{
				lower:          0,
				upper:          1,
				lowerUnbounded: true,
				upperUnbounded: true,
				lowerIncluded:  true,
				upperIncluded:  true,
			}},
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
		{name: "Should contain with included true and unbounded false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 5, upper: 9, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: true},
		{name: "Should NOT contain with included true and unbounded false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -1, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: false},
		{name: "Should NOT contain with included true and unbounded false and upper 11",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 5, upper: 11, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: false},
		{name: "Should NOT contain with upper-included false and unbounded false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: false},
			args: args[int]{other: &Interval[int]{lower: 5, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: false},
		{name: "Should NOT contain with lower-included false and unbounded false",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: false, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: false},
		{name: "Should contain with lower-included false and upper-unbounded true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: false, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: 0, upper: 11, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: false},
		{name: "Should contain with lower-included false and lower-unbounded true",
			i:    Interval[int]{lower: 0, upper: 10, lowerUnbounded: false, upperUnbounded: true, lowerIncluded: false, upperIncluded: true},
			args: args[int]{other: &Interval[int]{lower: -1, upper: 9, lowerUnbounded: false, upperUnbounded: false, lowerIncluded: true, upperIncluded: true}},
			want: false},
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
		// TODO: Add test cases.
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
	tests := []testCase[int /* TODO: Insert concrete types here */]{
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
	tests := []testCase[int /* TODO: Insert concrete types here */]{
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
	tests := []testCase[int /* TODO: Insert concrete types here */]{
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
	tests := []testCase[int /* TODO: Insert concrete types here */]{
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
	tests := []testCase[int /* TODO: Insert concrete types here */]{
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
	tests := []testCase[int /* TODO: Insert concrete types here */]{
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
	tests := []testCase[int /* TODO: Insert concrete types here */]{
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
	tests := []testCase[int /* TODO: Insert concrete types here */]{
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
	tests := []testCase[int /* TODO: Insert concrete types here */]{
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
	tests := []testCase[int /* TODO: Insert concrete types here */]{
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
