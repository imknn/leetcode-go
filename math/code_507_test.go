package math

import "testing"

func Test_checkPerfectNumber(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want bool
	}{
		{"1", 28, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkPerfectNumber(tt.num); got != tt.want {
				t.Errorf("checkPerfectNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
