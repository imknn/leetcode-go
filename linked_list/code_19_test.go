package linked_list

import (
	"reflect"
	"testing"
)

func sliceToList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	head := &ListNode{Val: vals[0]}
	cur := head
	for i := 1; i < len(vals); i++ {
		cur.Next = &ListNode{Val: vals[i]}
		cur = cur.Next
	}
	return head
}

func listToSlice(head *ListNode) []int {
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

func Test_removeNthFromEnd(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		n    int
		want []int
	}{
		{"示例1", []int{1, 2, 3, 4, 5}, 2, []int{1, 2, 3, 5}},
		{"示例2", []int{1}, 1, nil},
		{"示例3", []int{1, 2}, 1, []int{1}},
		{"删除头节点", []int{1, 2}, 2, []int{2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := sliceToList(tt.vals)
			got := removeNthFromEnd(head, tt.n)
			gotSlice := listToSlice(got)
			if !reflect.DeepEqual(gotSlice, tt.want) {
				t.Errorf("removeNthFromEnd() = %v, want %v", gotSlice, tt.want)
			}
		})
	}
}
