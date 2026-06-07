package string

import (
	"testing"
)

// 整数范围 输入确保在 1 到 3999 的范围内

func Test_intToRoman(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want string
	}{
		{"1", 3, "III"},
		{"2", 4, "IV"},
		{"3", 9, "IX"},
		{"4", 58, "LVIII"},
		{"5", 1994, "MCMXCIV"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intToRoman(tt.num); got != tt.want {
				t.Errorf("intToRoman() = %v, want %v", got, tt.want)
			}
		})
	}
}
