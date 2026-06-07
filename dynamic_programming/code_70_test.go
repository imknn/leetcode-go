package dynamic_programming

import "testing"

func Test_climbStairs(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"1", 1, 1},
		{"2", 2, 2},
		{"3", 3, 3},
		{"4", 16, 1597},
		{"5", 45, 1836311903},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairs(tt.n); got != tt.want {
				t.Errorf("climbStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
