package base

import (
	"reflect"
	"semanticdatabase/generics"
	"testing"
)

func TestIntervalBuilder_build_errors(t *testing.T) {
	type testCase[T generics.Number] struct {
		name       string
		i          IntervalBuilder[T]
		want       *Interval[T]
		er_size    int
		er_message string
	}
	tests := []testCase[int]{
		{name: "Test for Default Items", i: *NewIntervalBuilder[int](), er_size: 1, er_message: "Impossible interval constellation with lower: 0 == upper: 0 and lowerincluded or upperincluded being false"},
		{name: "Test for Default included and unbounded booleans", i: *NewIntervalBuilder[int]().setLower(0).setUpper(0), want: &Interval[int]{
			lower:          0,
			upper:          0,
			lowerUnbounded: false,
			upperUnbounded: false,
			lowerIncluded:  true,
			upperIncluded:  true,
		}},
		{name: "Test for setted Unbounded booleans", i: *NewIntervalBuilder[int]().setLower(0).setUpper(0).setLowerUnbounded(true).setUpperUnbounded(true), want: &Interval[int]{
			lower:          0,
			upper:          0,
			lowerUnbounded: true,
			upperUnbounded: true,
			lowerIncluded:  true,
			upperIncluded:  true,
		}},
		//{name: "Test for setted Included booleans", i: *NewIntervalBuilder[int]().setLower(0).setUpper(1).setLowerIncluded(true).setUpperIncluded(true), want: &Interval[int]{
		//	lower:          0,
		//	upper:          1,
		//	lowerUnbounded: false,
		//	upperUnbounded: false,
		//	lowerIncluded:  true,
		//	upperIncluded:  true,
		//}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, er := tt.i.Build()
			if len(er) != tt.er_size {
				t.Errorf("Build()-length errors = %v, want %v", len(er), tt.er_size)
			}
			if er[0].Error() != tt.er_message {
				t.Errorf("Build()-error-message = %v, want: %v", er[0].Error(), tt.er_message)
			}
		})
	}
}

func TestIntervalBuilder_build(t *testing.T) {
	type testCase[T generics.Number] struct {
		name string
		i    IntervalBuilder[T]
		want *Interval[T]
	}
	tests := []testCase[int]{
		{name: "Test for Default Items", i: *NewIntervalBuilder[int](), want: &Interval[int]{
			lower:          0,
			upper:          0,
			lowerUnbounded: false,
			upperUnbounded: false,
			lowerIncluded:  true,
			upperIncluded:  true,
		}},
		{name: "Test for Default included and unbounded booleans", i: *NewIntervalBuilder[int]().setLower(0).setUpper(0), want: &Interval[int]{
			lower:          0,
			upper:          0,
			lowerUnbounded: false,
			upperUnbounded: false,
			lowerIncluded:  true,
			upperIncluded:  true,
		}},
		{name: "Test for setted Unbounded booleans", i: *NewIntervalBuilder[int]().setLower(0).setUpper(0).setLowerUnbounded(true).setUpperUnbounded(true), want: &Interval[int]{
			lower:          0,
			upper:          0,
			lowerUnbounded: true,
			upperUnbounded: true,
			lowerIncluded:  true,
			upperIncluded:  true,
		}},
		{name: "Test for setted Included booleans", i: *NewIntervalBuilder[int]().setLower(0).setUpper(0).setLowerIncluded(true).setUpperIncluded(true), want: &Interval[int]{
			lower:          0,
			upper:          0,
			lowerUnbounded: true,
			upperUnbounded: true,
			lowerIncluded:  true,
			upperIncluded:  true,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := tt.i.Build(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Build() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntervalBuilder_setLower(t *testing.T) {
	type args[T generics.Number] struct {
		lower T
	}
	type testCase[T generics.Number] struct {
		name string
		i    IntervalBuilder[T]
		args args[T]
		want *IntervalBuilder[T]
	}
	tests := []testCase[int /* TODO: Insert concrete types here */]{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.setLower(tt.args.lower); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("setLower() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntervalBuilder_setLowerIncluded(t *testing.T) {
	type args struct {
		lowerIncluded bool
	}
	type testCase[T generics.Number] struct {
		name string
		i    IntervalBuilder[T]
		args args
		want *IntervalBuilder[T]
	}
	tests := []testCase[int /* TODO: Insert concrete types here */]{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.setLowerIncluded(tt.args.lowerIncluded); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("setLowerIncluded() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntervalBuilder_setLowerUnbounded(t *testing.T) {
	type args struct {
		lowerUnbounded bool
	}
	type testCase[T generics.Number] struct {
		name string
		i    IntervalBuilder[T]
		args args
		want *IntervalBuilder[T]
	}
	tests := []testCase[int /* TODO: Insert concrete types here */]{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.setLowerUnbounded(tt.args.lowerUnbounded); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("setLowerUnbounded() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntervalBuilder_setUpper(t *testing.T) {
	type args[T generics.Number] struct {
		upper T
	}
	type testCase[T generics.Number] struct {
		name string
		i    IntervalBuilder[T]
		args args[T]
		want *IntervalBuilder[T]
	}
	tests := []testCase[int /* TODO: Insert concrete types here */]{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.setUpper(tt.args.upper); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("setUpper() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntervalBuilder_setUpperIncluded(t *testing.T) {
	type args struct {
		upperIncluded bool
	}
	type testCase[T generics.Number] struct {
		name string
		i    IntervalBuilder[T]
		args args
		want *IntervalBuilder[T]
	}
	tests := []testCase[int /* TODO: Insert concrete types here */]{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.setUpperIncluded(tt.args.upperIncluded); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("setUpperIncluded() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntervalBuilder_setUpperUnbounded(t *testing.T) {
	type args struct {
		upperUnbounded bool
	}
	type testCase[T generics.Number] struct {
		name string
		i    IntervalBuilder[T]
		args args
		want *IntervalBuilder[T]
	}
	tests := []testCase[int /* TODO: Insert concrete types here */]{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.setUpperUnbounded(tt.args.upperUnbounded); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("setUpperUnbounded() = %v, want %v", got, tt.want)
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
	tests := []testCase[int /* TODO: Insert concrete types here */]{
		// TODO: Add test cases.
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
