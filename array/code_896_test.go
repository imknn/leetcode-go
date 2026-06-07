package array

import "testing"

func Test_isMonotonic(t *testing.T) {
	tests := []struct {
		name string
		A    []int
		want bool
	}{
		{"1", []int{1, 2, 2, 3}, true},
		{"2", []int{6, 5, 4, 4}, true},
		{"3", []int{1, 3, 2}, false},
		{"4", []int{1, 1, 1}, true},
		{"5", []int{}, true},
		{"6", []int{1}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMonotonic(tt.A); got != tt.want {
				t.Errorf("isMonotonic() = %v, want %v", got, tt.want)
			}
		})
	}
}
