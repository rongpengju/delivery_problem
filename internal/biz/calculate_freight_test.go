package biz

import "testing"

func TestCourier_Calculate(t *testing.T) {
	type fields struct {
		BasicPrice    int8
		OverPrice     int8
		BasicWeight   int8
		InsuranceRate int8
	}
	type args struct {
		weight float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
	}{
		{
			name: "Not more than 1kg",
			fields: fields{
				BasicPrice:    18,
				OverPrice:     5,
				BasicWeight:   1,
				InsuranceRate: 1,
			},
			args: args{weight: 0.8},
			want: 18,
		},
		{
			name: "Change base price",
			fields: fields{
				BasicPrice:    12,
				OverPrice:     5,
				BasicWeight:   1,
				InsuranceRate: 1,
			},
			args: args{weight: 0.8},
			want: 12,
		},
		{
			name: "Over 1kg",
			fields: fields{
				BasicPrice:    18,
				OverPrice:     5,
				BasicWeight:   1,
				InsuranceRate: 1,
			},
			args: args{weight: 1.3},
			want: 23,
		},
		{
			name: "Over 10kg",
			fields: fields{
				BasicPrice:    18,
				OverPrice:     5,
				BasicWeight:   1,
				InsuranceRate: 1,
			},
			args: args{weight: 10.7},
			want: 69,
		},
		{
			name: "Less than 0kg",
			fields: fields{
				BasicPrice:    18,
				OverPrice:     5,
				BasicWeight:   1,
				InsuranceRate: 1,
			},
			args: args{weight: -1.8},
			want: 0,
		},
		{
			name: "More than 100kg",
			fields: fields{
				BasicPrice:    18,
				OverPrice:     5,
				BasicWeight:   1,
				InsuranceRate: 1,
			},
			args: args{weight: 109.0},
			want: 0,
		},
		{
			name: "Happens to be an integer",
			fields: fields{
				BasicPrice:    18,
				OverPrice:     5,
				BasicWeight:   1,
				InsuranceRate: 1,
			},
			args: args{weight: 2.0},
			want: 23,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Courier{
				BasicPrice:    tt.fields.BasicPrice,
				OverPrice:     tt.fields.OverPrice,
				BasicWeight:   tt.fields.BasicWeight,
				InsuranceRate: tt.fields.InsuranceRate,
			}
			if got := c.Calculate(tt.args.weight); got != tt.want {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
