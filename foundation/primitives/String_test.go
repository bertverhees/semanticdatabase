package primitives

import (
	"reflect"
	"strconv"
	"testing"
)

func TestString_GreaterThan(t *testing.T) {
	type fields struct {
		value string
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
			name:   "Aaaa equals Aaaa",
			fields: fields{value: "Aaaa"},
			args:   args{other: NewString("Aaaa")},
			want:   NewBoolean(false),
		},
		{
			name:   "Aaaa greater than Bbbb",
			fields: fields{value: "Aaaa"},
			args:   args{other: NewString("Bbbb")},
			want:   NewBoolean(false),
		},
		{
			name:   "Bbbb greater than Aaaa",
			fields: fields{value: "Bbbb"},
			args:   args{other: NewString("Aaaa")},
			want:   NewBoolean(true),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &String{
				value: tt.fields.value,
			}
			if got := p.GreaterThan(tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GreaterThan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestString_GreaterThanOrEqual(t *testing.T) {
	type fields struct {
		value string
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
			name:   "Aaaa GreaterThanOrEqual Aaaa",
			fields: fields{value: "Aaaa"},
			args:   args{other: NewString("Aaaa")},
			want:   NewBoolean(true),
		},
		{
			name:   "Aaaa GreaterThanOrEqual Bbbb",
			fields: fields{value: "Aaaa"},
			args:   args{other: NewString("Bbbb")},
			want:   NewBoolean(false),
		},
		{
			name:   "Bbbb GreaterThanOrEqual Aaaa",
			fields: fields{value: "Bbbb"},
			args:   args{other: NewString("Aaaa")},
			want:   NewBoolean(true),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &String{
				value: tt.fields.value,
			}
			if got := p.GreaterThanOrEqual(tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GreaterThanOrEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestString_IsEqual(t *testing.T) {
	type fields struct {
		value string
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
			name:   "Aaaa equals Aaaa",
			fields: fields{value: "Aaaa"},
			args:   args{b: NewString("Aaaa")},
			want:   NewBoolean(true),
		},
		{
			name:   "Aaaa equals Bbbb",
			fields: fields{value: "Aaaa"},
			args:   args{b: NewString("Bbbb")},
			want:   NewBoolean(false),
		},
		{
			name:   "Bbbb equals Aaaa",
			fields: fields{value: "Bbbb"},
			args:   args{b: NewString("Aaaa")},
			want:   NewBoolean(false),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &String{
				value: tt.fields.value,
			}
			if got := p.IsEqual(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IsEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestString_LessThan(t *testing.T) {
	type fields struct {
		value string
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
			name:   "Aaaa less then Aaaa",
			fields: fields{value: "Aaaa"},
			args:   args{other: NewString("Aaaa")},
			want:   NewBoolean(false),
		},
		{
			name:   "Aaaa less than Bbbb",
			fields: fields{value: "Aaaa"},
			args:   args{other: NewString("Bbbb")},
			want:   NewBoolean(true),
		},
		{
			name:   "Bbbb less than Aaaa",
			fields: fields{value: "Bbbb"},
			args:   args{other: NewString("Aaaa")},
			want:   NewBoolean(false),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &String{
				value: tt.fields.value,
			}
			if got := p.LessThan(tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LessThan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestString_LessThanOrEqual(t *testing.T) {
	type fields struct {
		value string
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
			name:   "Aaaa LessThanOrEqual Aaaa",
			fields: fields{value: "Aaaa"},
			args:   args{other: NewString("Aaaa")},
			want:   NewBoolean(true),
		},
		{
			name:   "Aaaa LessThanOrEqual Bbbb",
			fields: fields{value: "Aaaa"},
			args:   args{other: NewString("Bbbb")},
			want:   NewBoolean(true),
		},
		{
			name:   "Bbbb LessThanOrEqual Aaaa",
			fields: fields{value: "Bbbb"},
			args:   args{other: NewString("Aaaa")},
			want:   NewBoolean(false),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &String{
				value: tt.fields.value,
			}
			if got := p.LessThanOrEqual(tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LessThanOrEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestString_ReturnStringFromIOrdered(t *testing.T) {
	type args struct {
		ordered IOrdered
	}
	tests := []struct {
		name string
		args args
		want *String
	}{
		// TODO: Add test cases.
		{
			name: "Double",
			args: args{NewDouble(12.0001)},
			want: NewString("12.0001"),
		},
		{
			name: "Real",
			args: args{NewReal(12.0001)},
			want: NewString("12.0001"),
		},
		{
			name: "Integer",
			args: args{NewInteger(12)},
			want: NewString("12"),
		},
		{
			name: "Integer64",
			args: args{NewInteger64(12)},
			want: NewString("12"),
		},
		{
			name: "Character",
			args: args{NewCharacter('a')},
			want: NewString("a"),
		},
		{
			name: "Octet",
			args: args{NewOctet(126)},
			want: NewString("126"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &String{
				value: "",
			}
			if got := p.ConvertFromIOrdered(tt.args.ordered); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertFromIOrdered() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestString_GreaterThanNumeric(t *testing.T) {
	type fields struct {
		value uint8
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
			name:   "1 equals 2",
			fields: fields{value: 1},
			args:   args{other: NewString(2)},
			want:   NewBoolean(false),
		},
		{
			name:   "2 greater than 1",
			fields: fields{value: 2},
			args:   args{other: NewString(1)},
			want:   NewBoolean(false),
		},
		{
			name:   "2 greater than 1",
			fields: fields{value: 2},
			args:   args{other: ConvertToStringFromIOrdered(uint8(1)},
			want:   NewBoolean(true),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &String{
				value: strconv.FormatInt(int64(tt.fields.value), 10),
			}
			if got := p.GreaterThan(tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GreaterThan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestString_GreaterThanOrEqualNumeric(t *testing.T) {
	type fields struct {
		value string
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
			name:   "Aaaa GreaterThanOrEqual Aaaa",
			fields: fields{value: "Aaaa"},
			args:   args{other: NewString("Aaaa")},
			want:   NewBoolean(true),
		},
		{
			name:   "Aaaa GreaterThanOrEqual Bbbb",
			fields: fields{value: "Aaaa"},
			args:   args{other: NewString("Bbbb")},
			want:   NewBoolean(false),
		},
		{
			name:   "Bbbb GreaterThanOrEqual Aaaa",
			fields: fields{value: "Bbbb"},
			args:   args{other: NewString("Aaaa")},
			want:   NewBoolean(true),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &String{
				value: tt.fields.value,
			}
			if got := p.GreaterThanOrEqual(tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GreaterThanOrEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestString_IsEqualNumeric(t *testing.T) {
	type fields struct {
		value string
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
			name:   "Aaaa equals Aaaa",
			fields: fields{value: "Aaaa"},
			args:   args{b: NewString("Aaaa")},
			want:   NewBoolean(true),
		},
		{
			name:   "Aaaa equals Bbbb",
			fields: fields{value: "Aaaa"},
			args:   args{b: NewString("Bbbb")},
			want:   NewBoolean(false),
		},
		{
			name:   "Bbbb equals Aaaa",
			fields: fields{value: "Bbbb"},
			args:   args{b: NewString("Aaaa")},
			want:   NewBoolean(false),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &String{
				value: tt.fields.value,
			}
			if got := p.IsEqual(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IsEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestString_LessThanNumeric(t *testing.T) {
	type fields struct {
		value string
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
			name:   "Aaaa less then Aaaa",
			fields: fields{value: "Aaaa"},
			args:   args{other: NewString("Aaaa")},
			want:   NewBoolean(false),
		},
		{
			name:   "Aaaa less than Bbbb",
			fields: fields{value: "Aaaa"},
			args:   args{other: NewString("Bbbb")},
			want:   NewBoolean(true),
		},
		{
			name:   "Bbbb less than Aaaa",
			fields: fields{value: "Bbbb"},
			args:   args{other: NewString("Aaaa")},
			want:   NewBoolean(false),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &String{
				value: tt.fields.value,
			}
			if got := p.LessThan(tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LessThan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestString_LessThanOrEqualNumeric(t *testing.T) {
	type fields struct {
		value string
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
			name:   "Aaaa LessThanOrEqual Aaaa",
			fields: fields{value: "Aaaa"},
			args:   args{other: NewString("Aaaa")},
			want:   NewBoolean(true),
		},
		{
			name:   "Aaaa LessThanOrEqual Bbbb",
			fields: fields{value: "Aaaa"},
			args:   args{other: NewString("Bbbb")},
			want:   NewBoolean(true),
		},
		{
			name:   "Bbbb LessThanOrEqual Aaaa",
			fields: fields{value: "Bbbb"},
			args:   args{other: NewString("Aaaa")},
			want:   NewBoolean(false),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &String{
				value: tt.fields.value,
			}
			if got := p.LessThanOrEqual(tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LessThanOrEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}
