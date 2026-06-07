package bit_manipulation

import "testing"

func Test_hasAlternatingBits(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want bool
	}{
		{"1", 5, true},
		{"2", 7, false},
		{"3", 11, false},
		{"4", 2, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasAlternatingBits(tt.n); got != tt.want {
				t.Errorf("hasAlternatingBits() = %v, want %v", got, tt.want)
			}
		})
	}
}
