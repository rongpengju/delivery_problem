package biz

import "testing"

func TestCalculate(t *testing.T) {
	type args struct {
		weight float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Not more than 1kg",
			args: args{weight: 0.8},
			want: 18,
		},
		{
			name: "Over 1kg",
			args: args{weight: 1.3},
			want: 23,
		},
		{
			name: "Over 10kg",
			args: args{weight: 10.7},
			want: 69,
		},
		{
			name: "Less than 0kg",
			args: args{weight: -1.8},
			want: 0,
		},
		{
			name: "More than 100kg",
			args: args{weight: 109.0},
			want: 0,
		},
		{
			name: "Happens to be an integer",
			args: args{weight: 2.0},
			want: 23,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Calculate(tt.args.weight); got != tt.want {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
