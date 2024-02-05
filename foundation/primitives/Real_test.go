package primitives

import (
	"reflect"
	"testing"
)

func TestReal_Add(t *testing.T) {
	type fields struct {
		value float32
	}
	type args struct {
		other INumeric
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   INumeric
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Real{
				value: tt.fields.value,
			}
			if got := p.Add(tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReal_Divide(t *testing.T) {
	type fields struct {
		value float32
	}
	type args struct {
		other INumeric
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   INumeric
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Real{
				value: tt.fields.value,
			}
			if got := p.Divide(tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Divide() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReal_Exponent(t *testing.T) {
	type fields struct {
		value float32
	}
	type args struct {
		other INumeric
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   INumeric
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Real{
				value: tt.fields.value,
			}
			if got := p.Exponent(tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Exponent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReal_GreaterThan(t *testing.T) {
	type fields struct {
		value float32
	}
	type args struct {
		other IOrdered
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Boolean
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Real{
				value: tt.fields.value,
			}
			if got := p.GreaterThan(tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GreaterThan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReal_GreaterThanOrEqual(t *testing.T) {
	type fields struct {
		value float32
	}
	type args struct {
		other IOrdered
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Boolean
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Real{
				value: tt.fields.value,
			}
			if got := p.GreaterThanOrEqual(tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GreaterThanOrEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReal_IsEqual(t *testing.T) {
	type fields struct {
		value float32
	}
	type args struct {
		b IAny
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   IAny
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Real{
				value: tt.fields.value,
			}
			if got := p.IsEqual(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IsEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReal_LessThan(t *testing.T) {
	type fields struct {
		value float32
	}
	type args struct {
		other IOrdered
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Boolean
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Real{
				value: tt.fields.value,
			}
			if got := p.LessThan(tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LessThan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReal_LessThanOrEqual(t *testing.T) {
	type fields struct {
		value float32
	}
	type args struct {
		other IOrdered
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Boolean
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Real{
				value: tt.fields.value,
			}
			if got := p.LessThanOrEqual(tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LessThanOrEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReal_Multiply(t *testing.T) {
	type fields struct {
		value float32
	}
	type args struct {
		other INumeric
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   INumeric
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Real{
				value: tt.fields.value,
			}
			if got := p.Multiply(tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReal_Negative(t *testing.T) {
	type fields struct {
		value float32
	}
	tests := []struct {
		name   string
		fields fields
		want   INumeric
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Real{
				value: tt.fields.value,
			}
			if got := p.Negative(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Negative() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReal_SetValue(t *testing.T) {
	type fields struct {
		value float32
	}
	type args struct {
		value float32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Real{
				value: tt.fields.value,
			}
			p.SetValue(tt.args.value)
		})
	}
}

func TestReal_Subtract(t *testing.T) {
	type fields struct {
		value float32
	}
	type args struct {
		other INumeric
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   INumeric
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Real{
				value: tt.fields.value,
			}
			if got := p.Subtract(tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Subtract() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReal_Value(t *testing.T) {
	type fields struct {
		value float32
	}
	tests := []struct {
		name   string
		fields fields
		want   float32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Real{
				value: tt.fields.value,
			}
			if got := p.Value(); got != tt.want {
				t.Errorf("Value() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReal_ToFixedNumberOfDecimals(t *testing.T) {
	type fields struct {
		value float32
	}
	type args struct {
		precision *Integer
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   IFloat
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Real{
				value: tt.fields.value,
			}
			if got := p.ToFixedNumberOfDecimals(tt.args.precision); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToFixedNumberOfDecimals() = %v, want %v", got, tt.want)
			}
		})
	}
}