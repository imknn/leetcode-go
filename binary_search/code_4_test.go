package binary_search

import "testing"

func Test_findMedianSortedArrays(t *testing.T) {
	tests := []struct {
		name string
		nums1 []int
		nums2 []int
		want  float64
	}{
		{"示例1", []int{1, 3}, []int{2}, 2.0},
		{"示例2", []int{1, 2}, []int{3, 4}, 2.5},
		{"一个为空", []int{1}, []int{}, 1.0},
		{"另一个为空", []int{}, []int{2}, 2.0},
		{"单元素", []int{1}, []int{2}, 1.5},
		{"相同元素", []int{1, 1}, []int{1, 1}, 1.0},
		{"交错", []int{1, 3, 5}, []int{2, 4, 6}, 3.5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findMedianSortedArrays(tt.nums1, tt.nums2)
			if got != tt.want {
				t.Errorf("findMedianSortedArrays(%v, %v) = %v, want %v", tt.nums1, tt.nums2, got, tt.want)
			}
		})
	}
}
