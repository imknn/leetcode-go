package array

import "testing"

func Test_containsDuplicate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{"1", []int{1, 2, 3, 1}, true},
		{"2", []int{1, 2, 3, 4}, false},
		{"3", []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsDuplicate(tt.nums); got != tt.want {
				t.Errorf("containsDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
