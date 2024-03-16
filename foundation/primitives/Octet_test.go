package primitives

import (
	"reflect"
	"testing"
)

func TestOctet_GreaterThan(t *testing.T) {
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
			name:   "GreaterThen 5 , 4",
			fields: fields{5},
			args:   args{NewOctet(4)},
			want:   NewBoolean(5 > 4),
		},
		{
			name:   "GreaterThen 5 , 5",
			fields: fields{5},
			args:   args{NewOctet(5)},
			want:   NewBoolean(5 > 5),
		},
		{
			name:   "GreaterThen 4 , 5",
			fields: fields{4},
			args:   args{NewOctet(5)},
			want:   NewBoolean(4 > 5),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Octet{
				value: tt.fields.value,
			}
			if got, _ := p.GreaterThan(tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GreaterThan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOctet_GreaterThanOrEqual(t *testing.T) {
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
			name:   "GreaterThanOrEqual 5 , 4",
			fields: fields{5},
			args:   args{NewOctet(4)},
			want:   NewBoolean(5 >= 4),
		},
		{
			name:   "GreaterThanOrEqual 5 , 5",
			fields: fields{5},
			args:   args{NewOctet(5)},
			want:   NewBoolean(5 >= 5),
		},
		{
			name:   "GreaterThanOrEqual 4 , 5",
			fields: fields{4},
			args:   args{NewOctet(5)},
			want:   NewBoolean(4 >= 5),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Octet{
				value: tt.fields.value,
			}
			if got, _ := p.GreaterThanOrEqual(tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GreaterThanOrEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOctet_IsEqual(t *testing.T) {
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
			name:   "IsEqual 5 , 4",
			fields: fields{5},
			args:   args{NewOctet(4)},
			want:   NewBoolean(5 == 4),
		},
		{
			name:   "IsEqual 5 , 5",
			fields: fields{5},
			args:   args{NewOctet(5)},
			want:   NewBoolean(5 == 5),
		},
		{
			name:   "IsEqual 4 , 5",
			fields: fields{4},
			args:   args{NewOctet(5)},
			want:   NewBoolean(4 == 5),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Octet{
				value: tt.fields.value,
			}
			if got := p.IsEqual(tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IsEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOctet_LessThan(t *testing.T) {
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
			name:   "LessThan 5 , 4",
			fields: fields{5},
			args:   args{NewOctet(4)},
			want:   NewBoolean(5 < 4),
		},
		{
			name:   "LessThan 5 , 5",
			fields: fields{5},
			args:   args{NewOctet(5)},
			want:   NewBoolean(5 < 5),
		},
		{
			name:   "LessThan 4 , 5",
			fields: fields{4},
			args:   args{NewOctet(5)},
			want:   NewBoolean(4 < 5),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Octet{
				value: tt.fields.value,
			}
			if got, _ := p.LessThan(tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LessThan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOctet_LessThanOrEqual(t *testing.T) {
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
			name:   "LessThanOrEqual 5 , 4",
			fields: fields{5},
			args:   args{NewOctet(4)},
			want:   NewBoolean(5 <= 4),
		},
		{
			name:   "LessThanOrEqual 5 , 5",
			fields: fields{5},
			args:   args{NewOctet(5)},
			want:   NewBoolean(5 <= 5),
		},
		{
			name:   "LessThanOrEqual 4 , 5",
			fields: fields{4},
			args:   args{NewOctet(5)},
			want:   NewBoolean(4 <= 5),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Octet{
				value: tt.fields.value,
			}
			if got, _ := p.LessThanOrEqual(tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LessThanOrEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}
