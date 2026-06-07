package math

import "testing"

func Test_missingNumber(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"1", []int{3, 0, 1}, 2},
		{"2", []int{9, 6, 4, 2, 3, 5, 7, 0, 1}, 8},
		{"3", []int{9, 6, 4, 2, 3, 7, 8, 0, 1, 10}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := missingNumber(tt.nums); got != tt.want {
				t.Errorf("missingNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
