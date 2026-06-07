package linked_list

// 19. 删除链表的倒数第 N 个结点
// 快慢指针：快指针先走 n 步，然后同步走，快指针到末尾时慢指针在目标位置
// 时间 O(L)，空间 O(1)
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Val: 0, Next: head}
	fast, slow := dummy, dummy

	// 快指针先走 n+1 步
	for i := 0; i <= n; i++ {
		fast = fast.Next
	}

	// 同步走
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	// 删除 slow.Next
	slow.Next = slow.Next.Next
	return dummy.Next
}
