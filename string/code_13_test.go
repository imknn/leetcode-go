package string

import (
	"testing"
)

func Test_romanToInt(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"1", "III", 3},
		{"2", "IV", 4},
		{"3", "IX", 9},
		{"4", "LVIII", 58},
		{"5", "MCMXCIV", 1994},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := romanToInt(tt.s); got != tt.want {
				t.Errorf("romanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
