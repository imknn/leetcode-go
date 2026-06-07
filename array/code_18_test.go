package array

import (
	"reflect"
	"sort"
	"testing"
)

func Test_fourSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   [][]int
	}{
		{"示例1", []int{1, 0, -1, 0, -2, 2}, 0, [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}}},
		{"示例2", []int{2, 2, 2, 2, 2}, 8, [][]int{{2, 2, 2, 2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := fourSum(tt.nums, tt.target)
			for _, s := range got {
				sort.Ints(s)
			}
			sortSlices(got)
			sortSlices(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fourSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
