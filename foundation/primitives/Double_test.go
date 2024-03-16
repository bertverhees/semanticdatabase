package primitives

import (
	"math"
	"reflect"
	"testing"
)

func TestDouble_Add(t *testing.T) {
	type fields struct {
		value float64
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
		{
			name:   "Add 5.1 + 4.8 Double",
			fields: fields{5.1},
			args:   args{NewDouble(4.8)},
			want:   NewDouble(5.1 + 4.8),
		},
		{
			name:   "Add 5.1 + 4.8 Real",
			fields: fields{5.1},
			args:   args{NewReal(4.8)},
			want:   NewDouble(5.1 + 4.8),
		},
		{
			name:   "Add 5.0 + 4 Integer",
			fields: fields{5.0},
			args:   args{NewInteger(4)},
			want:   NewDouble(5.0 + 4),
		},
		{
			name:   "Add 5.0 + 4 Integer64",
			fields: fields{5.0},
			args:   args{NewInteger64(4)},
			want:   NewDouble(5.0 + 4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Double{
				value: tt.fields.value,
			}
			if got, _ := p.Add(NewDouble(ToFixedNumberOfDecimals(tt.args.other.(*Double).value, 1))); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDouble_Divide(t *testing.T) {
	type fields struct {
		value float64
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
		{
			name:   "Divide 5.1 / 4.8 Double",
			fields: fields{5.1},
			args:   args{NewDouble(4.8)},
			want:   NewDouble(ToFixedNumberOfDecimals(5.1/4.8, 2)),
		},
		{
			name:   "Divide 5.1 / 4.8 Real",
			fields: fields{5.1},
			args:   args{NewReal(4.8)},
			want:   NewDouble(ToFixedNumberOfDecimals(5.1/4.8, 2)),
		},
		{
			name:   "Divide 5.0 / 4 Integer",
			fields: fields{5.0},
			args:   args{NewInteger(4)},
			want:   NewDouble(5.0 / 4),
		},
		{
			name:   "Divide 5.0 / 4 Integer64",
			fields: fields{5.0},
			args:   args{NewInteger64(4)},
			want:   NewDouble(5.0 / 4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Double{
				value: tt.fields.value,
			}
			if got, _ := p.Divide(NewDouble(ToFixedNumberOfDecimals(tt.args.other.(*Double).value, 2))); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Divide() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDouble_Exponent(t *testing.T) {
	type fields struct {
		value float64
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
		{
			name:   "Exponent 5.1 ^ 4.8 Double",
			fields: fields{5.1},
			args:   args{NewDouble(4.8)},
			want:   NewDouble(ToFixedNumberOfDecimals(math.Pow(5.1, 4.8), 2)),
		},
		{
			name:   "Exponent 5.1 ^ 4.8 Real",
			fields: fields{5.1},
			args:   args{NewReal(4.8)},
			want:   NewDouble(ToFixedNumberOfDecimals(math.Pow(5.1, 4.8), 2)),
		},
		{
			name:   "Exponent 5.0 ^ 4 Integer",
			fields: fields{5.0},
			args:   args{NewInteger(4)},
			want:   NewDouble(math.Pow(5.0, 4)),
		},
		{
			name:   "Exponent 5.0 ^ 4 Integer64",
			fields: fields{5.0},
			args:   args{NewInteger64(4)},
			want:   NewDouble(math.Pow(5.0, 4)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Double{
				value: tt.fields.value,
			}
			if got, _ := p.Exponent(NewDouble(ToFixedNumberOfDecimals(tt.args.other.(*Double).value, 2))); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Exponent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDouble_GreaterThan(t *testing.T) {
	type fields struct {
		value float64
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
		{
			name:   "GreaterThan 5.1 > 4.8 Double",
			fields: fields{5.1},
			args:   args{NewDouble(4.8)},
			want:   NewBoolean(5.1 > 4.8),
		},
		{
			name:   "GreaterThan 4.8 > 5.1 Double",
			fields: fields{4.8},
			args:   args{NewDouble(5.1)},
			want:   NewBoolean(4.8 > 5.1),
		},
		{
			name:   "GreaterThan 5.1 > 5.1 Double",
			fields: fields{5.1},
			args:   args{NewDouble(5.1)},
			want:   NewBoolean(5.1 > 5.1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Double{
				value: tt.fields.value,
			}
			if got, _ := p.GreaterThan(tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GreaterThan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDouble_GreaterThanOrEqual(t *testing.T) {
	type fields struct {
		value float64
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
		{
			name:   "GreaterThanOrEqual 5.1 >= 4.8 Double",
			fields: fields{5.1},
			args:   args{NewDouble(4.8)},
			want:   NewBoolean(5.1 >= 4.8),
		},
		{
			name:   "GreaterThanOrEqual 4.8 >= 5.1 Double",
			fields: fields{4.8},
			args:   args{NewDouble(5.1)},
			want:   NewBoolean(4.8 >= 5.1),
		},
		{
			name:   "GreaterThanOrEqual 5.1 >= 5.1 Double",
			fields: fields{5.1},
			args:   args{NewDouble(5.1)},
			want:   NewBoolean(5.1 >= 5.1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Double{
				value: tt.fields.value,
			}
			if got, _ := p.GreaterThanOrEqual(tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GreaterThanOrEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDouble_IsEqual(t *testing.T) {
	type fields struct {
		value float64
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
		{
			name:   "IsEqual 4.8 == 5.1 Double",
			fields: fields{4.8},
			args:   args{NewDouble(5.1)},
			want:   NewBoolean(4.8 == 5.1),
		},
		{
			name:   "IsEqual 5.1 == 5.1 Double",
			fields: fields{5.1},
			args:   args{NewDouble(5.1)},
			want:   NewBoolean(5.1 == 5.1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Double{
				value: tt.fields.value,
			}
			if got := p.IsEqual(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IsEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDouble_LessThan(t *testing.T) {
	type fields struct {
		value float64
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
		{
			name:   "LessThan 5.1 <= 4.8 Double",
			fields: fields{5.1},
			args:   args{NewDouble(4.8)},
			want:   NewBoolean(5.1 < 4.8),
		},
		{
			name:   "LessThan 4.8 <= 5.1 Double",
			fields: fields{4.8},
			args:   args{NewDouble(5.1)},
			want:   NewBoolean(4.8 < 5.1),
		},
		{
			name:   "LessThan 5.1 <= 5.1 Double",
			fields: fields{5.1},
			args:   args{NewDouble(5.1)},
			want:   NewBoolean(5.1 < 5.1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Double{
				value: tt.fields.value,
			}
			if got, _ := p.LessThan(tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LessThan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDouble_LessThanOrEqual(t *testing.T) {
	type fields struct {
		value float64
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
		{
			name:   "LessThanOrEqual 5.1 <= 4.8 Double",
			fields: fields{5.1},
			args:   args{NewDouble(4.8)},
			want:   NewBoolean(5.1 <= 4.8),
		},
		{
			name:   "LessThanOrEqual 4.8 <= 5.1 Double",
			fields: fields{4.8},
			args:   args{NewDouble(5.1)},
			want:   NewBoolean(4.8 <= 5.1),
		},
		{
			name:   "LessThanOrEqual 5.1 <= 5.1 Double",
			fields: fields{5.1},
			args:   args{NewDouble(5.1)},
			want:   NewBoolean(5.1 <= 5.1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Double{
				value: tt.fields.value,
			}
			if got, _ := p.LessThanOrEqual(tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LessThanOrEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDouble_Multiply(t *testing.T) {
	type fields struct {
		value float64
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
		{
			name:   "Multiply 5.1 * 4.8 Double",
			fields: fields{5.1},
			args:   args{NewDouble(4.8)},
			want:   NewDouble(ToFixedNumberOfDecimals(5.1*4.8, 2)),
		},
		{
			name:   "Multiply 5.1 * 4.8 Real",
			fields: fields{5.1},
			args:   args{NewReal(4.8)},
			want:   NewDouble(ToFixedNumberOfDecimals(5.1*4.8, 2)),
		},
		{
			name:   "Multiply 5.0 * 4 Integer",
			fields: fields{5.0},
			args:   args{NewInteger(4)},
			want:   NewDouble(5.0 * 4),
		},
		{
			name:   "Multiply 5.0 * 4 Integer64",
			fields: fields{5.0},
			args:   args{NewInteger64(4)},
			want:   NewDouble(5.0 * 4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Double{
				value: tt.fields.value,
			}
			if got, _ := p.Multiply(NewDouble(ToFixedNumberOfDecimals(tt.args.other.(*Double).value, 2))); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDouble_Negative(t *testing.T) {
	type fields struct {
		value float64
	}
	tests := []struct {
		name   string
		fields fields
		want   INumeric
	}{
		{
			name:   "Negative -5.1",
			fields: fields{-5.1},
			want:   NewDouble(5.1),
		},
		{
			name:   "Negative 5.1",
			fields: fields{5.1},
			want:   NewDouble(-5.1),
		},
		{
			name:   "Negative 0",
			fields: fields{0},
			want:   NewDouble(0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Double{
				value: tt.fields.value,
			}
			if got, _ := p.Negative(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Negative() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDouble_Subtract(t *testing.T) {
	type fields struct {
		value float64
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
		{
			name:   "Subtract 5.1 - 4.8 Double",
			fields: fields{5.1},
			args:   args{NewDouble(4.8)},
			want:   NewDouble(5.1 - 4.8),
		},
		{
			name:   "Subtract 5.1 - 4.8 Real",
			fields: fields{5.1},
			args:   args{NewReal(4.8)},
			want:   NewDouble(5.1 - 4.8),
		},
		{
			name:   "Subtract 5.0 - 4 Integer",
			fields: fields{5.0},
			args:   args{NewInteger(4)},
			want:   NewDouble(5.0 - 4),
		},
		{
			name:   "Subtract 5.0 - 4 Integer64",
			fields: fields{5.0},
			args:   args{NewInteger64(4)},
			want:   NewDouble(5.0 - 4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Double{
				value: tt.fields.value,
			}
			if got, _ := p.Subtract(NewDouble(ToFixedNumberOfDecimals(tt.args.other.(*Double).value, 1))); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Subtract() = %v, want %v", got, tt.want)
			}
		})
	}
}
