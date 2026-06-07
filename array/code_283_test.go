package array

import (
	"fmt"
	"testing"
)

func Test_moveZeroes(t *testing.T) {
	tests := []struct {
		name string
		nums []int
	}{
		{"1", []int{0, 1, 0, 3, 12}},
		{"2", []int{0, 0, 0, 3, 12}},
		{"3", []int{1, 2, 0, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moveZeroes(tt.nums)
			fmt.Println(tt.nums)
		})
	}
}
