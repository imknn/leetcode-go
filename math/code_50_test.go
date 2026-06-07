package math

import "testing"

func Test_myPow(t *testing.T) {
	type args struct {
		x float64
		n int
	}
	tests := []struct {
		args args
		want float64
	}{
		{args{2.00000, 10}, 1024.00000},
		//{args{2.10000, 3}, 9.26100},
		//{args{2.00000, -2}, 0.25000},
		{args{2.00000, 4}, 16},
		{args{2.00000, 1}, 2},
		//{args{2.00000, 3}, 8},
		{args{0.00001, 2147483647}, 0},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := myPow(tt.args.x, tt.args.n); got != tt.want {
				t.Errorf("myPow() = %v, want %v", got, tt.want)
			}
		})
	}
}
