package julius

import "testing"

func Test_leaps(t *testing.T) {
	type args struct {
		year1 uint64
		year2 uint64
	}
	tests := []struct {
		args args
		want uint64
	}{

		{
			args: args{
				year1: 2016,
				year2: 2017,
			},
			want: 1,
		},
		{
			args: args{
				year1: 2019,
				year2: 2020,
			},
			want: 0,
		},
		{
			args: args{
				year1: 1900,
				year2: 1901,
			},
			want: 0,
		},
		{
			args: args{
				year1: 2000,
				year2: 2001,
			},
			want: 1,
		},
		{
			args: args{
				year1: 2800,
				year2: 2801,
			},
			want: 0,
		},
		{
			args: args{
				year1: 123456,
				year2: 123456,
			},
			want: 0,
		},
		{
			args: args{
				year1: 1234,
				year2: 5678,
			},
			want: 1077,
		},
		{
			args: args{
				year1: 123456,
				year2: 7891011,
			},
			want: 1881475,
		},
		{
			args: args{
				year1: 123456789101112,
				year2: 1314151617181920,
			},
			want: 288412747246240,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := leaps(tt.args.year1, tt.args.year2); got != tt.want {
				t.Errorf("leaps() = %v, want %v", got, tt.want)
			}
		})
	}
}
