package array

import (
	"fmt"
	"testing"
)

func Test_rotate(t *testing.T) {
	tests := []struct {
		matrix [][]int
	}{
		{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}},
	}
	for _, tt := range tests {
		t.Run("48", func(t *testing.T) {
			rotate(tt.matrix)
			fmt.Println(tt)
		})
	}
}
