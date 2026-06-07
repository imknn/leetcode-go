# 83. 删除排序链表中的重复元素

## 题目信息

| 属性 | 值 |
|------|-----|
| 题号 | [83. 删除排序链表中的重复元素](https://leetcode.cn/problems/remove-duplicates-from-sorted-list/) |
| 难度 | Easy |
| 分类 | linked_list |
| 标签 | linked list |
| 文件 | `linked_list/code_83.go` |

## 题目描述

给定一个已排序的链表的头 `head`，删除所有重复的元素，使每个元素只出现一次。返回已排序的链表。

**示例 1：**

```
输入：head = [1,1,2]
输出：[1,2]
```

**示例 2：**

```
输入：head = [1,1,2,3,3]
输出：[1,2,3]
```

## 解题算法

### 方法：一次遍历

**核心思路：**
由于链表已经排序，重复元素一定是相邻的。只需要一次遍历链表，对于每个节点，比较它和下一个节点的值：如果相同则跳过下一个节点（删除重复），否则移动到下一个节点。

**代码实现：**
```go
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	current := head
	for current.Next != nil {
		if current.Val == current.Next.Val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}
	return head
}
```

**复杂度分析：**
- 时间复杂度：O(n)，其中 n 是链表的长度
- 空间复杂度：O(1)，只使用了常数级别的额外空间
