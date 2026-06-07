package math

import "testing"

func Test_countPrimes(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"1", 10, 4},
		{"2", 10000, 1229},
		{"3", 499979, 41537},
		{"4", 1500000, 114155},
		{"5", 999983, 78497},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPrimes(tt.n); got != tt.want {
				t.Errorf("countPrimes() = %v, want %v", got, tt.want)
			}
		})
	}
}
