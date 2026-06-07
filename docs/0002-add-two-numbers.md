# 2. 两数相加

## 题目信息

| 属性 | 值 |
|------|-----|
| 题号 | [2. 两数相加](https://leetcode.cn/problems/add-two-numbers/) |
| 难度 | Medium |
| 分类 | linked_list |
| 标签 | linked list, math |
| 文件 | `linked_list/code_2.go` |

## 题目描述

给你两个 **非空** 的链表，表示两个非负的整数。它们每位数字都是按照 **逆序** 的方式存储的，并且每个节点只能存储 **一位** 数字。

请你将两个数相加，并以相同形式返回一个表示和的链表。

你可以假设除了数字 `0` 之外，这两个数都不会以 `0` 开头。

**示例 1：**

```
输入：l1 = [2,4,3], l2 = [5,6,4]
输出：[7,0,8]
解释：342 + 465 = 807。
```

**示例 2：**

```
输入：l1 = [0], l2 = [0]
输出：[0]
```

**示例 3：**

```
输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
输出：[8,9,9,9,0,0,0,1]
```

## 解题算法

### 方法：模拟加法

**核心思路：**
同时遍历两个链表，逐位相加并记录进位。每次将两链表当前节点的值与进位相加，对 10 取余得到当前位的值，除以 10 得到新的进位。当两个链表都遍历完后，如果还有进位，则额外添加一个节点。使用哑节点简化头节点的处理。

**代码实现：**
```go
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
```

**复杂度分析：**
- 时间复杂度：O(max(m, n))，其中 m 和 n 分别为两个链表的长度
- 空间复杂度：O(max(m, n))，结果链表的长度最多为 max(m, n) + 1
