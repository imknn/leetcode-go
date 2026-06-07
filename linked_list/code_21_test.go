package linked_list

import (
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	var l1, next *ListNode // 创建头节点，new 返回指针
	next = new(ListNode)
	l1 = next
	nums := []int{1, 2, 4}
	for x := 0; x < len(nums); x++ {
		next.Val = nums[x]
		if x == len(nums)-1 {
			next.Next = nil
		} else {
			next.Next = new(ListNode)
			next = next.Next
		}
	}
	printList(t, l1)
	t.Log("===========================")

	var l2 *ListNode // 创建头节点，new 返回指针
	next = new(ListNode)
	l2 = next
	nums = []int{1, 3, 4}
	for x := 0; x < len(nums); x++ {
		next.Val = nums[x]
		if x == len(nums)-1 {
			next.Next = nil
		} else {
			next.Next = new(ListNode)
			next = next.Next
		}
	}
	printList(t, l2)
	t.Log("===========================")

	printList(t, mergeTwoLists(l1, l2))
}

func printList(t *testing.T, head *ListNode) {
	for head != nil {
		t.Log(head.Val)
		head = head.Next
	}
}
