package linked_list

import (
	"reflect"
	"testing"
)

func node(nums []int) *ListNode {
	var root, current *ListNode
	for _, num := range nums {
		var one ListNode
		one.Val = num
		if root == nil {
			root = &one
		} else {
			current.Next = &one
		}
		current = &one
	}
	return root
}

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"1", args{node([]int{2, 4, 3}), node([]int{5, 6, 4})}, node([]int{7, 0, 8})},
		{"2", args{node([]int{0}), node([]int{0})}, node([]int{0})},
		{"3", args{node([]int{9, 9, 9, 9, 9, 9, 9}), node([]int{9, 9, 9, 9})}, node([]int{8, 9, 9, 9, 0, 0, 0, 1})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
