package array

import (
	"reflect"
	"testing"
)

func Test_generate(t *testing.T) {
	tests := []struct {
		numRows int
		want    [][]int
	}{
		{1, [][]int{{1}}},
		{2, [][]int{{1}, {1, 1}}},
		{5, [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}}},
	}
	for _, tt := range tests {
		t.Run("118", func(t *testing.T) {
			if got := generate(tt.numRows); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generate() = %v, want %v", got, tt.want)
			}
		})
	}
}
