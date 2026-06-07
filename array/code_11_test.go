package array

import "testing"

func Test_maxArea(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{"示例1", []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{"示例2", []int{1, 1}, 1},
		{"三个元素", []int{1, 2, 1}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxArea(tt.height); got != tt.want {
				t.Errorf("maxArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
