package bit_manipulation

import "testing"

func Test_isPowerOfFour(t *testing.T) {
	tests := []struct {
		name string
		args int
		want bool
	}{
		{"1", 16, true},
		{"2", 5, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPowerOfFour(tt.args); got != tt.want {
				t.Errorf("isPowerOfFour() = %v, want %v", got, tt.want)
			}
		})
	}
}
