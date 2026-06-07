# 21. 合并两个有序链表

## 题目信息

| 属性 | 值 |
|------|-----|
| 题号 | [21. 合并两个有序链表](https://leetcode.cn/problems/merge-two-sorted-lists/) |
| 难度 | Easy |
| 分类 | linked_list |
| 标签 | linked list, recursion |
| 文件 | `linked_list/code_21.go` |

## 题目描述

将两个升序链表合并为一个新的 **升序** 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

**示例 1：**

```
输入：l1 = [1,2,4], l2 = [1,3,4]
输出：[1,1,2,3,4,4]
```

**示例 2：**

```
输入：l1 = [], l2 = []
输出：[]
```

**示例 3：**

```
输入：l1 = [], l2 = [0]
输出：[0]
```

## 解题算法

### 方法：迭代合并

**核心思路：**
创建一个哑节点作为结果链表的头部，然后同时遍历两个链表，每次将较小的节点接到结果链表的末尾。当其中一个链表遍历完毕后，将另一个链表的剩余部分直接接到结果链表末尾。

**代码实现：**
```go
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
```

**复杂度分析：**
- 时间复杂度：O(m + n)，其中 m 和 n 分别为两个链表的长度
- 空间复杂度：O(1)，只使用了常数级别的额外空间
