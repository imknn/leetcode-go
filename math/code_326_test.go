package math

import "testing"

func Test_isPowerOfThree(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want bool
	}{
		{"1", 27, true},
		{"2", 0, false},
		{"3", 9, true},
		{"4", 45, false},
		{"5", 1, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPowerOfThree(tt.n); got != tt.want {
				t.Errorf("isPowerOfThree() = %v, want %v", got, tt.want)
			}
		})
	}
}
