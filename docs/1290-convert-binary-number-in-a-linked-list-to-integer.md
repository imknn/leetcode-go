# 1290. 二进制链表转整数

## 题目信息

| 属性 | 值 |
|------|-----|
| 题号 | [1290. 二进制链表转整数](https://leetcode.cn/problems/convert-binary-number-in-a-linked-list-to-integer/) |
| 难度 | Easy |
| 分类 | linked_list |
| 标签 | linked list, bit manipulation |
| 文件 | `linked_list/code_1290.go` |

## 题目描述

给你一个单链表的引用结点 `head`。链表中每个结点的值不是 0 就是 1。已知此链表是一个整数数字的二进制表示形式。

请你返回该链表所表示数字的 **十进制值** 。

**示例 1：**

```
输入：head = [1,0,1]
输出：5
解释：二进制数 (101) 转化为十进制数 (5)
```

**示例 2：**

```
输入：head = [0]
输出：0
```

**示例 3：**

```
输入：head = [1,0,0,1,0,0,1,1,1,0,0,0,0,0,0]
输出：18880
```

## 解题算法

### 方法：位运算

**核心思路：**
遍历链表，利用位运算将二进制转换为十进制。从高位到低位依次处理，每移动到下一个节点时，将当前结果左移一位（相当于乘以 2），然后加上当前节点的值。这样逐位累加即可得到最终的十进制数。

**代码实现：**
```go
func getDecimalValue(head *ListNode) int {
	num := head.Val
	for head.Next != nil {
		head = head.Next
		num = num<<1 + head.Val
	}
	return num
}
```

**复杂度分析：**
- 时间复杂度：O(n)，其中 n 是链表的长度
- 空间复杂度：O(1)，只使用了常数级别的额外空间
