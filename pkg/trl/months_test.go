package trl

import (
	"fmt"
	"testing"
)

func TestMonthsShift(t *testing.T) {
	type args struct {
		i     int
		shift int
	}
	tests := []struct {
		// name string
		args args
		want int
	}{
		{args{5, 1}, 6},
		{args{2, -1}, 1},

		{args{11, 2}, 1},
		{args{1, 25}, 2},
		{args{12, 23}, 11},

		{args{2, -3}, 11},
		{args{5, -3}, 2},
		{args{8, -3}, 5},
		{args{11, -3}, 8},

		{args{2, -2}, 12},
		{args{7, -8}, 11},
		{args{7, -18}, 1},

		{args{7, -19}, 12},
		{args{7, -20}, 11},
		{args{7, -21}, 10},
		{args{7, -22}, 9},

		{args{7, -23}, 8},
		{args{7, -24}, 7},
		{args{7, -25}, 6},

		{args{7, -48}, 7},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("t%v", tt.args), func(t *testing.T) {
			if got := MonthsShift(tt.args.i, tt.args.shift); got != tt.want {
				t.Errorf("\nMonthsShift(%v,%v) = %v, want %v", tt.args.i, tt.args.shift, got, tt.want)
			}
		})
	}
}
