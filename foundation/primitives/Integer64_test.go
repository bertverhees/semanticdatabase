package primitives

import (
	"reflect"
	"testing"
)

func TestInteger64_Add(t *testing.T) {
	type args struct {
		other INumeric
	}
	tests := []struct {
		name string
		args args
		want INumeric
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			in := Integer64{}
			if got, _ := in.Add(tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInteger64_Divide(t *testing.T) {
	type args struct {
		other INumeric
	}
	tests := []struct {
		name string
		args args
		want INumeric
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			in := Integer64{}
			if got, _ := in.Divide(tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Divide() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInteger64_Exponent(t *testing.T) {
	type args struct {
		other INumeric
	}
	tests := []struct {
		name string
		args args
		want INumeric
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			in := Integer64{}
			if got, _ := in.Exponent(tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Exponent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInteger64_GreaterThan(t *testing.T) {
	type args struct {
		other IOrdered
	}
	tests := []struct {
		name string
		args args
		want *Boolean
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			in := Integer64{}
			if got, _ := in.GreaterThan(tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GreaterThan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInteger64_GreaterThanOrEqual(t *testing.T) {
	type args struct {
		other IOrdered
	}
	tests := []struct {
		name string
		args args
		want *Boolean
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			in := Integer64{}
			if got, _ := in.GreaterThanOrEqual(tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GreaterThanOrEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInteger64_IsEqual(t *testing.T) {
	type args struct {
		any IAny
	}
	tests := []struct {
		name string
		args args
		want IAny
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			in := Integer64{}
			if got := in.IsEqual(tt.args.any); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IsEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInteger64_LessThan(t *testing.T) {
	type args struct {
		other IOrdered
	}
	tests := []struct {
		name string
		args args
		want *Boolean
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			in := Integer64{}
			if got, _ := in.LessThan(tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LessThan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInteger64_LessThanOrEqual(t *testing.T) {
	type args struct {
		other IOrdered
	}
	tests := []struct {
		name string
		args args
		want *Boolean
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			in := Integer64{}
			if got, _ := in.LessThanOrEqual(tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LessThanOrEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInteger64_Multiply(t *testing.T) {
	type args struct {
		other INumeric
	}
	tests := []struct {
		name string
		args args
		want INumeric
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			in := Integer64{}
			if got, _ := in.Multiply(tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInteger64_Negative(t *testing.T) {
	tests := []struct {
		name string
		want INumeric
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			in := Integer64{}
			if got, _ := in.Negative(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Negative() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInteger64_Subtract(t *testing.T) {
	type args struct {
		other INumeric
	}
	tests := []struct {
		name string
		args args
		want INumeric
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			in := Integer64{}
			if got, _ := in.Subtract(tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Subtract() = %v, want %v", got, tt.want)
			}
		})
	}
}
