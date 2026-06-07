package dynamic_programming

import (
	"testing"
)

func Test_fib(t *testing.T) {
	tests := []struct {
		name string
		N    int
		want int
	}{
		{"0", 0, 0},
		{"1", 1, 1},
		{"2", 2, 1},
		{"3", 3, 2},
		{"4", 4, 3},
		{"5", 5, 5},
		{"6", 6, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fib(tt.N); got != tt.want {
				t.Errorf("fib() = %v, want %v", got, tt.want)
			}
		})
	}
}
