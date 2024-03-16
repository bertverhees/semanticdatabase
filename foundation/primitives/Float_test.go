package primitives

import (
	"math"
	"reflect"
	"testing"
)

func TestDouble_ToFixedNumberOfDecimals(t *testing.T) {
	type fields struct {
		value float64
	}
	type args struct {
		precision int
	}
	tests := []struct {
		name string
		args args
		want *Double
	}{
		{
			name: "ToFixedNumberOfDecimals Precision 1",
			args: args{1},
			want: NewDouble(3.1),
		},
		{
			name: "ToFixedNumberOfDecimals Precision 2",
			args: args{2},
			want: NewDouble(3.14),
		},
		{
			name: "ToFixedNumberOfDecimals Precision 3",
			args: args{3},
			want: NewDouble(3.142),
		},
		{
			name: "ToFixedNumberOfDecimals Precision 4",
			args: args{4},
			want: NewDouble(3.1416),
		},
		{
			name: "ToFixedNumberOfDecimals Precision 5",
			args: args{5},
			want: NewDouble(3.14159),
		},
		{
			name: "ToFixedNumberOfDecimals Precision 6",
			args: args{6},
			want: NewDouble(3.141593),
		},
		{
			name: "ToFixedNumberOfDecimals Precision 7",
			args: args{7},
			want: NewDouble(3.1415927),
		},
		{
			name: "ToFixedNumberOfDecimals Precision 8",
			args: args{8},
			want: NewDouble(3.14159265),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Double{
				value: math.Pi,
			}
			if got := ToFixedNumberOfDecimals(p.value, tt.args.precision); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToFixedNumberOfDecimals() = %v, want %v", got, tt.want)
			}
		})
	}
}
