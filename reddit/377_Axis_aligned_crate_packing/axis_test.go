package axis

import (
	"testing"
)

func Test_fit1(t *testing.T) {
	type args struct {
		x int
		y int
		a int
		b int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				x: 25,
				y: 18,
				a: 6,
				b: 5,
			},
			want: 12,
		},
		{
			args: args{
				x: 10,
				y: 10,
				a: 1,
				b: 1,
			},
			want: 100,
		},
		{
			args: args{
				x: 12,
				y: 34,
				a: 5,
				b: 6,
			},
			want: 10,
		},
		{
			args: args{
				x: 12345,
				y: 678910,
				a: 1112,
				b: 1314,
			},
			want: 5676,
		},
		{
			args: args{
				x: 5,
				y: 100,
				a: 6,
				b: 1,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := fit1(tt.args.x, tt.args.y, tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("fit1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fit2(t *testing.T) {
	type args struct {
		x int
		y int
		a int
		b int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				x: 25,
				y: 18,
				a: 6,
				b: 5,
			},
			want: 15,
		},
		{
			args: args{
				x: 12,
				y: 34,
				a: 5,
				b: 6,
			},
			want: 12,
		},
		{
			args: args{
				x: 12345,
				y: 678910,
				a: 1112,
				b: 1314,
			},
			want: 5676,
		},
		{
			args: args{
				x: 5,
				y: 5,
				a: 3,
				b: 2,
			},
			want: 2,
		},
		{
			args: args{
				x: 5,
				y: 100,
				a: 6,
				b: 1,
			},
			want: 80,
		},
		{
			args: args{
				x: 5,
				y: 5,
				a: 6,
				b: 1,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := fit2(tt.args.x, tt.args.y, tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("fit2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fit3(t *testing.T) {
	type args struct {
		x int
		y int
		z int
		a int
		b int
		c int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				x: 10,
				y: 10,
				z: 10,
				a: 1,
				b: 1,
				c: 1,
			},
			want: 1000,
		},
		{
			args: args{
				x: 12,
				y: 34,
				z: 56,
				a: 7,
				b: 8,
				c: 9,
			},
			want: 32,
		},
		{
			args: args{
				x: 123,
				y: 456,
				z: 789,
				a: 10,
				b: 11,
				c: 12,
			},
			want: 32604,
		},
		{
			args: args{
				x: 1234567,
				y: 89101112,
				z: 13141516,
				a: 171819,
				b: 202122,
				c: 232425,
			},
			want: 174648,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := fit3(tt.args.x, tt.args.y, tt.args.z, tt.args.a, tt.args.b, tt.args.c); got != tt.want {
				t.Errorf("fit3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fitn(t *testing.T) {
	type args struct {
		crate []int
		box   []int
	}
	tests := []struct {
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := fitn(tt.args.crate, tt.args.box); got != tt.want {
				t.Errorf("fitn() = %v, want %v", got, tt.want)
			}
		})
	}
}
