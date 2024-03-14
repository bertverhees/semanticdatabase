package primitives

import (
	"reflect"
	"testing"
)

func TestCharacter_GreaterThan(t *testing.T) {
	type fields struct {
		value rune
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
			name:   "a equals a GreaterThan",
			fields: fields{value: 'a'},
			args:   args{other: NewCharacter('a')},
			want:   NewBoolean(false),
		},
		{
			name:   "b greater than a GreaterThan",
			fields: fields{value: 'b'},
			args:   args{other: NewCharacter('a')},
			want:   NewBoolean(true),
		},
		{
			name:   "a smaller than b GreaterThan",
			fields: fields{value: 'a'},
			args:   args{other: NewCharacter('b')},
			want:   NewBoolean(false),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Character{
				value: tt.fields.value,
			}
			if got, _ := p.GreaterThan(tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GreaterThan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCharacter_GreaterThanOrEqual(t *testing.T) {
	type fields struct {
		value rune
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
			name:   "a equals a GreaterThanOrEqual",
			fields: fields{value: 'a'},
			args:   args{other: NewCharacter('a')},
			want:   NewBoolean(true),
		},
		{
			name:   "b greater than a GreaterThanOrEqual",
			fields: fields{value: 'b'},
			args:   args{other: NewCharacter('a')},
			want:   NewBoolean(true),
		},
		{
			name:   "a smaller than b GreaterThanOrEqual",
			fields: fields{value: 'a'},
			args:   args{other: NewCharacter('b')},
			want:   NewBoolean(false),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Character{
				value: tt.fields.value,
			}
			if got, _ := p.GreaterThanOrEqual(tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GreaterThanOrEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCharacter_IsEqual(t *testing.T) {
	type fields struct {
		value rune
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
			name:   "a equals a IsEqual",
			fields: fields{value: 'a'},
			args:   args{b: NewCharacter('a')},
			want:   NewBoolean(true),
		},
		{
			name:   "b greater than a IsEqual",
			fields: fields{value: 'b'},
			args:   args{b: NewCharacter('a')},
			want:   NewBoolean(false),
		},
		{
			name:   "a smaller than b IsEqual",
			fields: fields{value: 'a'},
			args:   args{b: NewCharacter('b')},
			want:   NewBoolean(false),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Character{
				value: tt.fields.value,
			}
			if got := p.IsEqual(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IsEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCharacter_LessThan(t *testing.T) {
	type fields struct {
		value rune
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
			name:   "a equals a LessThan",
			fields: fields{value: 'a'},
			args:   args{other: NewCharacter('a')},
			want:   NewBoolean(false),
		},
		{
			name:   "b greater than a LessThan",
			fields: fields{value: 'b'},
			args:   args{other: NewCharacter('a')},
			want:   NewBoolean(false),
		},
		{
			name:   "a smaller than b LessThan",
			fields: fields{value: 'a'},
			args:   args{other: NewCharacter('b')},
			want:   NewBoolean(true),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Character{
				value: tt.fields.value,
			}
			if got, _ := p.LessThan(tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LessThan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCharacter_LessThanOrEqual(t *testing.T) {
	type fields struct {
		value rune
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
			name:   "a equals a LessThanOrEqual",
			fields: fields{value: 'a'},
			args:   args{other: NewCharacter('a')},
			want:   NewBoolean(true),
		},
		{
			name:   "b greater than a LessThanOrEqual",
			fields: fields{value: 'b'},
			args:   args{other: NewCharacter('a')},
			want:   NewBoolean(false),
		},
		{
			name:   "a smaller than b LessThanOrEqual",
			fields: fields{value: 'a'},
			args:   args{other: NewCharacter('b')},
			want:   NewBoolean(true),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Character{
				value: tt.fields.value,
			}
			if got, _ := p.LessThanOrEqual(tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LessThanOrEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCharacter_returnCharacterFromIOrdered(t *testing.T) {
	type fields struct {
		value rune
	}
	type args struct {
		ordered IOrdered
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Character
	}{
		{
			name: "Double",
			args: args{NewDouble(12.0001)},
			want: NewCharacter(12),
		},
		{
			name: "Real",
			args: args{NewReal(12.0001)},
			want: NewCharacter(12),
		},
		{
			name: "Integer",
			args: args{NewInteger(1200)},
			want: NewCharacter(1200),
		},
		{
			name: "Integer64",
			args: args{NewInteger64(65000)},
			want: NewCharacter(65000),
		},
		{
			name: "Octet",
			args: args{NewOctet(127)},
			want: NewCharacter(127),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := ToCharacter(tt.args.ordered); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("returnCharacterFromIOrdered() = %v, want %v", got, tt.want)
			}
		})
	}
}
