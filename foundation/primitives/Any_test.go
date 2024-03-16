package primitives

import (
	"reflect"
	"testing"
)

func TestAny_AsBoolean(t *testing.T) {
	tests := []struct {
		name    string
		want    *Boolean
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Any{}
			got, err := a.AsBoolean()
			if (err != nil) != tt.wantErr {
				t.Errorf("AsBoolean() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AsBoolean() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAny_AsCharacter(t *testing.T) {
	tests := []struct {
		name    string
		want    *Character
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Any{}
			got, err := a.AsCharacter()
			if (err != nil) != tt.wantErr {
				t.Errorf("AsCharacter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AsCharacter() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAny_AsDouble(t *testing.T) {
	tests := []struct {
		name    string
		want    *Double
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Any{}
			got, err := a.AsDouble()
			if (err != nil) != tt.wantErr {
				t.Errorf("AsDouble() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AsDouble() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAny_AsInteger(t *testing.T) {
	tests := []struct {
		name    string
		want    *Integer
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Any{}
			got, err := a.AsInteger()
			if (err != nil) != tt.wantErr {
				t.Errorf("AsInteger() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AsInteger() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAny_AsInteger64(t *testing.T) {
	tests := []struct {
		name    string
		want    *Integer64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Any{}
			got, err := a.AsInteger64()
			if (err != nil) != tt.wantErr {
				t.Errorf("AsInteger64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AsInteger64() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAny_AsOctet(t *testing.T) {
	tests := []struct {
		name    string
		want    *Octet
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Any{}
			got, err := a.AsOctet()
			if (err != nil) != tt.wantErr {
				t.Errorf("AsOctet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AsOctet() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAny_AsReal(t *testing.T) {
	tests := []struct {
		name    string
		want    *Real
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Any{}
			got, err := a.AsReal()
			if (err != nil) != tt.wantErr {
				t.Errorf("AsReal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AsReal() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAny_AsString(t *testing.T) {
	tests := []struct {
		name    string
		want    *String
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Any{}
			got, err := a.AsString()
			if (err != nil) != tt.wantErr {
				t.Errorf("AsString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AsString() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAny_AsUri(t *testing.T) {
	tests := []struct {
		name    string
		want    *Uri
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Any{}
			got, err := a.AsUri()
			if (err != nil) != tt.wantErr {
				t.Errorf("AsUri() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AsUri() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAny_IsEqual(t *testing.T) {
	type args struct {
		any IAny
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
			a := &Any{}
			if got := a.IsEqual(tt.args.any); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IsEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAny_NotEqual(t *testing.T) {
	type args struct {
		any IAny
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
			a := &Any{}
			if got := a.NotEqual(tt.args.any); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NotEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInstanceOf(t *testing.T) {
	type args struct {
		aType String
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
			if got := InstanceOf(tt.args.aType); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InstanceOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToBoolean(t *testing.T) {
	type args struct {
		ordered IAny
	}
	tests := []struct {
		name    string
		args    args
		want    *Boolean
		wantErr bool
	}{
		{
			name:    "Boolean to boolean",
			args:    args{NewBoolean(true)},
			want:    NewBoolean(true),
			wantErr: false,
		},
		{
			name:    "Boolean to boolean",
			args:    args{NewBoolean(false)},
			want:    NewBoolean(false),
			wantErr: false,
		},
		{
			name:    "Character to boolean",
			args:    args{NewCharacter('t')},
			want:    NewBoolean(true),
			wantErr: false,
		},
		{
			name:    "Character to boolean",
			args:    args{NewCharacter('f')},
			want:    NewBoolean(false),
			wantErr: false,
		},
		{
			name:    "Character to boolean",
			args:    args{NewCharacter('q')},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Double to boolean",
			args:    args{NewDouble(1)},
			want:    NewBoolean(true),
			wantErr: false,
		},
		{
			name:    "Double to boolean",
			args:    args{NewDouble(0)},
			want:    NewBoolean(false),
			wantErr: false,
		},
		{
			name:    "Double to boolean",
			args:    args{NewDouble(-1)},
			want:    NewBoolean(false),
			wantErr: false,
		},
		{
			name:    "Integer to boolean",
			args:    args{NewInteger(1)},
			want:    NewBoolean(true),
			wantErr: false,
		},
		{
			name:    "Integer to boolean",
			args:    args{NewInteger(0)},
			want:    NewBoolean(false),
			wantErr: false,
		},
		{
			name:    "Integer to boolean",
			args:    args{NewInteger(-1)},
			want:    NewBoolean(false),
			wantErr: false,
		},
		{
			name:    "Integer64 to boolean",
			args:    args{NewInteger64(1)},
			want:    NewBoolean(true),
			wantErr: false,
		},
		{
			name:    "Integer64 to boolean",
			args:    args{NewInteger64(0)},
			want:    NewBoolean(false),
			wantErr: false,
		},
		{
			name:    "Integer64 to boolean",
			args:    args{NewInteger64(-1)},
			want:    NewBoolean(false),
			wantErr: false,
		},
		{
			name:    "Octet to boolean",
			args:    args{NewOctet(1)},
			want:    NewBoolean(true),
			wantErr: false,
		},
		{
			name:    "Octet to boolean",
			args:    args{NewOctet(0)},
			want:    NewBoolean(false),
			wantErr: false,
		},
		{
			name:    "Real to boolean",
			args:    args{NewReal(1)},
			want:    NewBoolean(true),
			wantErr: false,
		},
		{
			name:    "Real to boolean",
			args:    args{NewReal(0)},
			want:    NewBoolean(false),
			wantErr: false,
		},
		{
			name:    "Real to boolean",
			args:    args{NewReal(-1)},
			want:    NewBoolean(false),
			wantErr: false,
		},
		{
			name:    "String to boolean",
			args:    args{NewString("true")},
			want:    NewBoolean(true),
			wantErr: false,
		},
		{
			name:    "String to boolean",
			args:    args{NewString("false")},
			want:    NewBoolean(false),
			wantErr: false,
		},
		{
			name:    "String to boolean",
			args:    args{NewString("t")},
			want:    NewBoolean(true),
			wantErr: false,
		},
		{
			name:    "String to boolean",
			args:    args{NewString("f")},
			want:    NewBoolean(false),
			wantErr: false,
		},
		{
			name:    "String to boolean",
			args:    args{NewString("something else")},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Uri to boolean",
			args:    args{NewUri("http://blabla.bla")},
			want:    NewBoolean(false),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToBoolean(tt.args.ordered)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToBoolean() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToBoolean() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToCharacter(t *testing.T) {
	type args struct {
		ordered IAny
	}
	tests := []struct {
		name    string
		args    args
		want    *Character
		wantErr bool
	}{
		{
			name:    "Boolean to Character",
			args:    args{NewBoolean(true)},
			want:    NewCharacter('t'),
			wantErr: false,
		},
		{
			name:    "Boolean to Character",
			args:    args{NewBoolean(false)},
			want:    NewCharacter('f'),
			wantErr: false,
		},
		{
			name:    "Boolean to Character",
			args:    args{NewBoolean(true)},
			want:    NewCharacter('f'),
			wantErr: false,
		},
		{
			name:    "Boolean to Character",
			args:    args{NewBoolean(false)},
			want:    NewCharacter('t'),
			wantErr: false,
		},
		{
			name:    "Boolean to Character",
			args:    args{NewBoolean(false)},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Character to Character",
			args:    args{NewCharacter('t')},
			want:    NewCharacter('t'),
			wantErr: false,
		},
		{
			name:    "Character to Character",
			args:    args{NewCharacter('f')},
			want:    NewCharacter('f'),
			wantErr: false,
		},
		{
			name:    "Character to Character",
			args:    args{NewCharacter('q')},
			want:    NewCharacter('q'),
			wantErr: true,
		},
		{
			name:    "Double to Character",
			args:    args{NewDouble(200)},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Double to Character",
			args:    args{NewDouble(1)},
			want:    NewCharacter(rune(1)),
			wantErr: false,
		},
		{
			name:    "Double to Character",
			args:    args{NewDouble(0)},
			want:    NewCharacter(rune(0)),
			wantErr: false,
		},
		{
			name:    "Double to Character",
			args:    args{NewDouble(-1)},
			want:    NewCharacter(rune(-1)),
			wantErr: false,
		},
		{
			name:    "Double to Character",
			args:    args{NewDouble(-200)},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Integer to Character",
			args:    args{NewInteger(200)},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Integer to Character",
			args:    args{NewInteger(1)},
			want:    NewCharacter(rune(1)),
			wantErr: false,
		},
		{
			name:    "Integer to Character",
			args:    args{NewInteger(0)},
			want:    NewCharacter(rune(0)),
			wantErr: false,
		},
		{
			name:    "Integer to Character",
			args:    args{NewInteger(-1)},
			want:    NewCharacter(rune(-1)),
			wantErr: false,
		},
		{
			name:    "Integer to Character",
			args:    args{NewInteger(-200)},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Integer64 to Character",
			args:    args{NewInteger64(200)},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Integer64 to Character",
			args:    args{NewInteger64(1)},
			want:    NewCharacter(rune(1)),
			wantErr: false,
		},
		{
			name:    "Integer64 to Character",
			args:    args{NewInteger64(0)},
			want:    NewCharacter(rune(0)),
			wantErr: false,
		},
		{
			name:    "Integer64 to Character",
			args:    args{NewInteger64(-1)},
			want:    NewCharacter(rune(-1)),
			wantErr: false,
		},
		{
			name:    "Integer64 to Character",
			args:    args{NewInteger64(-200)},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Octet to Character",
			args:    args{NewOctet(200)},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Octet to Character",
			args:    args{NewOctet(1)},
			want:    NewCharacter(rune(1)),
			wantErr: false,
		},
		{
			name:    "Octet to Character",
			args:    args{NewOctet(0)},
			want:    NewCharacter(rune(0)),
			wantErr: false,
		},
		{
			name:    "Real to Character",
			args:    args{NewReal(200)},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Real to Character",
			args:    args{NewReal(1)},
			want:    NewCharacter(rune(1)),
			wantErr: false,
		},
		{
			name:    "Real to Character",
			args:    args{NewReal(0)},
			want:    NewCharacter(rune(0)),
			wantErr: false,
		},
		{
			name:    "Real to Character",
			args:    args{NewReal(-1)},
			want:    NewCharacter(rune(-1)),
			wantErr: false,
		},
		{
			name:    "Real to Character",
			args:    args{NewReal(-200)},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "String to Character",
			args:    args{NewString("true")},
			want:    NewCharacter('t'),
			wantErr: false,
		},
		{
			name:    "String to Character",
			args:    args{NewString("false")},
			want:    NewCharacter('f'),
			wantErr: false,
		},
		{
			name:    "String to Character",
			args:    args{NewString("t")},
			want:    NewCharacter('t'),
			wantErr: false,
		},
		{
			name:    "String to Character",
			args:    args{NewString("f")},
			want:    NewCharacter('f'),
			wantErr: false,
		},
		{
			name:    "String to Character",
			args:    args{NewString("something else")},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Uri to Character",
			args:    args{NewUri("http://blabla.bla")},
			want:    NewCharacter(' '),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToCharacter(tt.args.ordered)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToCharacter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToCharacter() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToDouble(t *testing.T) {
	type args struct {
		ordered IAny
	}
	tests := []struct {
		name    string
		args    args
		want    *Double
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToDouble(tt.args.ordered)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToDouble() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToDouble() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToInteger(t *testing.T) {
	type args struct {
		ordered IAny
	}
	tests := []struct {
		name    string
		args    args
		want    *Integer
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToInteger(tt.args.ordered)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToInteger() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToInteger() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToInteger64(t *testing.T) {
	type args struct {
		ordered IAny
	}
	tests := []struct {
		name    string
		args    args
		want    *Integer64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToInteger64(tt.args.ordered)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToInteger64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToInteger64() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToOctet(t *testing.T) {
	type args struct {
		ordered IAny
	}
	tests := []struct {
		name    string
		args    args
		want    *Octet
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToOctet(tt.args.ordered)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToOctet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToOctet() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToReal(t *testing.T) {
	type args struct {
		ordered IAny
	}
	tests := []struct {
		name    string
		args    args
		want    *Real
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToReal(tt.args.ordered)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToReal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToReal() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToString(t *testing.T) {
	type args struct {
		ordered IAny
	}
	tests := []struct {
		name    string
		args    args
		want    *String
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToString(tt.args.ordered)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToString() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToUri(t *testing.T) {
	type args struct {
		ordered IAny
	}
	tests := []struct {
		name    string
		args    args
		want    *Uri
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToUri(tt.args.ordered)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToUri() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToUri() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTypeOf(t *testing.T) {
	type args struct {
		anObject IAny
	}
	tests := []struct {
		name string
		args args
		want *String
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TypeOf(tt.args.anObject); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TypeOf() = %v, want %v", got, tt.want)
			}
		})
	}
}
