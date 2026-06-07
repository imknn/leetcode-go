package linked_list

/* 21. 合并两个有序链表*/
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var header = new(ListNode) // 创建头节点，new 返回指针
	current := header
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}
		current = current.Next
	}
	if l1 == nil {
		current.Next = l2
	} else {
		current.Next = l1
	}
	return header.Next
}
