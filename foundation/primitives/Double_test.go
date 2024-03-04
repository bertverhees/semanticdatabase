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
			if got := p.Add(tt.args.other).(*Double).ToFixedNumberOfDecimals(NewInteger(1)); !reflect.DeepEqual(got, tt.want) {
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
			want:   NewDouble(5.1 / 4.8).ToFixedNumberOfDecimals(NewInteger(2)).(*Double),
		},
		{
			name:   "Divide 5.1 / 4.8 Real",
			fields: fields{5.1},
			args:   args{NewReal(4.8)},
			want:   NewDouble(5.1 / 4.8).ToFixedNumberOfDecimals(NewInteger(2)).(*Double),
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
			if got := p.Divide(tt.args.other).(*Double).ToFixedNumberOfDecimals(NewInteger(2)); !reflect.DeepEqual(got, tt.want) {
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
			want:   NewDouble(math.Pow(5.1, 4.8)).ToFixedNumberOfDecimals(NewInteger(2)).(*Double),
		},
		{
			name:   "Exponent 5.1 ^ 4.8 Real",
			fields: fields{5.1},
			args:   args{NewReal(4.8)},
			want:   NewDouble(math.Pow(5.1, 4.8)).ToFixedNumberOfDecimals(NewInteger(2)).(*Double),
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
			if got := p.Exponent(tt.args.other).(*Double).ToFixedNumberOfDecimals(NewInteger(2)); !reflect.DeepEqual(got, tt.want) {
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
			if got := p.GreaterThan(tt.args.other); !reflect.DeepEqual(got, tt.want) {
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
			if got := p.GreaterThanOrEqual(tt.args.other); !reflect.DeepEqual(got, tt.want) {
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
			if got := p.LessThan(tt.args.other); !reflect.DeepEqual(got, tt.want) {
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
			if got := p.LessThanOrEqual(tt.args.other); !reflect.DeepEqual(got, tt.want) {
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
			want:   NewDouble(5.1 * 4.8).ToFixedNumberOfDecimals(NewInteger(2)).(*Double),
		},
		{
			name:   "Multiply 5.1 * 4.8 Real",
			fields: fields{5.1},
			args:   args{NewReal(4.8)},
			want:   NewDouble(5.1 * 4.8).ToFixedNumberOfDecimals(NewInteger(2)).(*Double),
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
			if got := p.Multiply(tt.args.other).(*Double).ToFixedNumberOfDecimals(NewInteger(2)); !reflect.DeepEqual(got, tt.want) {
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
			if got := p.Negative(); !reflect.DeepEqual(got, tt.want) {
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
			if got := p.Subtract(tt.args.other).(*Double).ToFixedNumberOfDecimals(NewInteger(1)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Subtract() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDouble_ConvertFromIOrdered(t *testing.T) {
	type args struct {
		ordered IOrdered
	}
	tests := []struct {
		name string
		args args
		want *Double
	}{
		{
			name: "ConvertFromIOrdered 4.8 Double",
			args: args{NewDouble(4.8)},
			want: NewDouble(4.8),
		},
		{
			name: "ConvertFromIOrdered 4.8 Real",
			args: args{NewReal(4.8)},
			want: NewDouble(4.8),
		},
		{
			name: "ConvertFromIOrdered 4 Integer",
			args: args{NewInteger(4)},
			want: NewDouble(4),
		},
		{
			name: "ConvertFromIOrdered 4 Integer64",
			args: args{NewInteger64(4)},
			want: NewDouble(4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Double{
				value: 0,
			}
			if got := p.ConvertFromIOrdered(tt.args.ordered).(*Double).ToFixedNumberOfDecimals(NewInteger(2)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertFromIOrdered() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDouble_ToFixedNumberOfDecimals(t *testing.T) {
	type fields struct {
		value float64
	}
	type args struct {
		precision *Integer
	}
	tests := []struct {
		name string
		args args
		want IFloat
	}{
		{
			name: "ToFixedNumberOfDecimals Precision 1",
			args: args{NewInteger(1)},
			want: NewDouble(3.1),
		},
		{
			name: "ToFixedNumberOfDecimals Precision 2",
			args: args{NewInteger(2)},
			want: NewDouble(3.14),
		},
		{
			name: "ToFixedNumberOfDecimals Precision 3",
			args: args{NewInteger(3)},
			want: NewDouble(3.142),
		},
		{
			name: "ToFixedNumberOfDecimals Precision 4",
			args: args{NewInteger(4)},
			want: NewDouble(3.1416),
		},
		{
			name: "ToFixedNumberOfDecimals Precision 5",
			args: args{NewInteger(5)},
			want: NewDouble(3.14159),
		},
		{
			name: "ToFixedNumberOfDecimals Precision 6",
			args: args{NewInteger(6)},
			want: NewDouble(3.141593),
		},
		{
			name: "ToFixedNumberOfDecimals Precision 7",
			args: args{NewInteger(7)},
			want: NewDouble(3.1415927),
		},
		{
			name: "ToFixedNumberOfDecimals Precision 8",
			args: args{NewInteger(8)},
			want: NewDouble(3.14159265),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Double{
				value: math.Pi,
			}
			if got := p.ToFixedNumberOfDecimals(tt.args.precision); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToFixedNumberOfDecimals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDouble_ConvertFromINumeric(t *testing.T) {
	type fields struct {
		value float64
	}
	type args struct {
		ordered INumeric
	}
	tests := []struct {
		name string
		args args
		want *Double
	}{
		{
			name: "ConvertFromINumeric 4.8 Double",
			args: args{NewDouble(4.8)},
			want: NewDouble(4.8),
		},
		{
			name: "ConvertFromINumeric 4.8 Real",
			args: args{NewReal(4.8)},
			want: NewDouble(4.8),
		},
		{
			name: "ConvertFromINumeric 4 Integer",
			args: args{NewInteger(4)},
			want: NewDouble(4),
		},
		{
			name: "ConvertFromINumeric 4 Integer64",
			args: args{NewInteger64(4)},
			want: NewDouble(4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Double{
				value: 0,
			}
			if got := p.ConvertFromINumeric(tt.args.ordered).(*Double).ToFixedNumberOfDecimals(NewInteger(2)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertFromINumeric() = %v, want %v", got, tt.want)
			}
		})
	}
}
