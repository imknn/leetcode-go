package math

import "testing"

func Test_smallestEvenMultiple(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"1", 5, 10},
		{"2", 6, 6},
		{"3", 8, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestEvenMultiple(tt.n); got != tt.want {
				t.Errorf("smallestEvenMultiple() = %v, want %v", got, tt.want)
			}
		})
	}
}
