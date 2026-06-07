package math

import "testing"

func Test_trailingZeroes(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"1", 3, 0},
		{"2", 5, 1},
		{"3", 100, 24},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trailingZeroes(tt.n); got != tt.want {
				t.Errorf("trailingZeroes() = %v, want %v", got, tt.want)
			}
		})
	}
}
