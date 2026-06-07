package binary_search

import "testing"

func Test_mySqrt(t *testing.T) {
	tests := []struct {
		name string
		x    int
		want int
	}{
		{"0", 1, 1},
		{"1", 4, 2},
		{"2", 8, 2},
		{"3", 16, 4},
		{"4", 17, 4},
		{"5", 100, 10},
		{"6", 99, 9},
		{"7", 101, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mySqrt(tt.x); got != tt.want {
				t.Errorf("mySqrt() = %v, want %v", got, tt.want)
			}
		})
	}
}
