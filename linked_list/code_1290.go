package linked_list

func getDecimalValue(head *ListNode) int {
	num := head.Val
	for head.Next != nil {
		head = head.Next
		num = num<<1 + head.Val
	}
	return num
}
