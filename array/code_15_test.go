package array

import (
	"reflect"
	"sort"
	"testing"
)

func sortSlices(slices [][]int) {
	for _, s := range slices {
		sort.Ints(s)
	}
	sort.Slice(slices, func(i, j int) bool {
		for k := 0; k < len(slices[i]) && k < len(slices[j]); k++ {
			if slices[i][k] != slices[j][k] {
				return slices[i][k] < slices[j][k]
			}
		}
		return len(slices[i]) < len(slices[j])
	})
}

func Test_threeSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{"示例1", []int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{"示例2", []int{0, 1, 1}, nil},
		{"示例3", []int{0, 0, 0}, [][]int{{0, 0, 0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := threeSum(tt.nums)
			sortSlices(got)
			sortSlices(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("threeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
