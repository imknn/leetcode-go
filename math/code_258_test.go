package math

import "testing"

func Test_addDigits(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want int
	}{
		{"1", 38, 2},
		{"2", 11111, 5},
		{"3", 9999999, 9},
		{"4", 1, 1},
		{"5", 0, 0},
		{"6", 2, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addDigits(tt.num); got != tt.want {
				t.Errorf("addDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
