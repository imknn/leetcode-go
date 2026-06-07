package linked_list

/* 237. 删除链表中的节点*/

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

/* 什么鬼，就是写一个处理逻辑 */

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
