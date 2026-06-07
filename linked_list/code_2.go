package linked_list

// 2. 两数相加
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var root, current *ListNode
	carry := 0 // 进位
	for l1 != nil || l2 != nil {
		var sum ListNode
		sum.Val = 0
		if l1 == nil {
			sum.Val = l2.Val + carry
			l2 = l2.Next
		} else if l2 == nil {
			sum.Val = l1.Val + carry
			l1 = l1.Next
		} else {
			sum.Val = l1.Val + l2.Val + carry
			l1 = l1.Next
			l2 = l2.Next
		}
		if sum.Val >= 10 {
			sum.Val -= 10
			carry = 1
		} else {
			carry = 0
		}
		if root == nil {
			root = &sum
		} else {
			current.Next = &sum
		}
		current = &sum
	}
	if carry > 0 { // 最后仍然有进位
		var sum = ListNode{carry, nil}
		current.Next = &sum
		carry = 0
		current = &sum
	}
	return root
}
