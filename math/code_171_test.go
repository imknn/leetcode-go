package math

import "testing"

func Test_titleToNumber(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"1", "A", 1},
		{"2", "AB", 28},
		{"3", "ZY", 701},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := titleToNumber(tt.s); got != tt.want {
				t.Errorf("titleToNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
