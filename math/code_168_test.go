package math

import "testing"

func Test_convertToTitle(t *testing.T) {
	tests := []struct {
		name string
		args int
		want string
	}{
		{"1", 1, "A"},
		{"2", 26, "Z"},
		{"3", 27, "AA"},
		{"4", 28, "AB"},
		{"5", 52, "AZ"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertToTitle(tt.args); got != tt.want {
				t.Errorf("convertToTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}
