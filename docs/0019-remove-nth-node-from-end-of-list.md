# 19. 删除链表的倒数第 N 个结点

## 题目信息

| 属性 | 值 |
|------|-----|
| 题号 | [19. 删除链表的倒数第 N 个结点](https://leetcode.cn/problems/remove-nth-node-from-end-of-list/) |
| 难度 | Medium |
| 分类 | linked_list |
| 标签 | linked list, two pointers |
| 文件 | `linked_list/code_19.go` |

## 题目描述

给你一个链表，删除链表的倒数第 `n` 个结点，并且返回链表的头结点。

**示例 1：**

```
输入：head = [1,2,3,4,5], n = 2
输出：[1,2,3,5]
```

**示例 2：**

```
输入：head = [1], n = 1
输出：[]
```

**示例 3：**

```
输入：head = [1,2], n = 1
输出：[1]
```

## 解题算法

### 方法：快慢指针

**核心思路：**
使用快慢指针，快指针先走 `n+1` 步，然后两个指针同步前进。当快指针到达链表末尾时（`nil`），慢指针恰好指向待删除节点的前一个节点。这样就可以通过 `slow.Next = slow.Next.Next` 删除目标节点。使用哑节点处理头节点被删除的情况。

**代码实现：**
```go
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
```

**复杂度分析：**
- 时间复杂度：O(L)，其中 L 是链表的长度，只需要一次遍历
- 空间复杂度：O(1)，只使用了常数级别的额外空间
