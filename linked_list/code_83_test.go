package linked_list

import (
	"reflect"
	"testing"
)

func Test_deleteDuplicates(t *testing.T) {
	tests := []struct {
		name string
		args *ListNode
		want *ListNode
	}{
		{"1", node([]int{1, 1, 2}), node([]int{1, 2})},
		{"1", node([]int{1, 1, 2, 3, 3}), node([]int{1, 2, 3})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicates(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
