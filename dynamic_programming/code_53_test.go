package dynamic_programming

import (
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	if x := maxSubArray(nums); x != 6 {
		t.Fail()
	}

	nums = []int{-2}
	if x := maxSubArray(nums); x != -2 {
		t.Fail()
	}
}

func Test_maxSubArray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"1", []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{"2", []int{-2}, -2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubArray(tt.nums); got != tt.want {
				t.Errorf("maxSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
