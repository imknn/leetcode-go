# 20. 有效的括号

## 题目信息

| 属性 | 值 |
|------|-----|
| 题号 | [20. 有效的括号](https://leetcode.cn/problems/valid-parentheses/) |
| 难度 | Easy |
| 分类 | stack_queue |
| 标签 | stack, string |
| 文件 | `stack_queue/code_20.go` |

## 题目描述

给定一个只包含 `'('`、`')'`、`'{'`、`'}'`、`'['`、`']'` 的字符串 `s`，判断字符串是否有效。

有效字符串需满足：

1. 左括号必须用相同类型的右括号闭合。
2. 左括号必须以正确的顺序闭合。
3. 每个右括号都有一个对应的相同类型的左括号。

**示例 1：**

```
输入：s = "()"
输出：true
```

**示例 2：**

```
输入：s = "()[]{}"
输出：true
```

**示例 3：**

```
输入：s = "(]"
输出：false
```

## 解题算法

### 方法：栈

**核心思路：**
使用栈来匹配括号。遍历字符串，遇到左括号就入栈，遇到右括号就检查栈顶是否是对应的左括号。如果匹配则弹出栈顶，否则返回 false。遍历结束后，如果栈为空则说明所有括号都正确匹配。

在实现中使用数组模拟栈，通过 `top` 索引跟踪栈顶位置，避免了 Go 语言切片操作的额外开销。

**代码实现：**
```go
func isValid(s string) bool {
	length := len(s)
	if length == 0 {
		return true
	}
	if length%2 != 0 {
		return false
	}
	// 栈
	top := -1 // 栈顶索引
	save := make([]byte, length)
	for i := 0; i < length; i++ {
		switch s[i] {
		case '(', '[', '{':
			// 入栈
			top++
			save[top] = s[i]
		case ')':
			if top < 0 || save[top] != '(' {
				return false
			} else {
				top--
			}
		case ']':
			if top < 0 || save[top] != '[' {
				return false
			} else {
				top--
			}
		case '}':
			if top < 0 || save[top] != '{' {
				return false
			} else {
				top--
			}
		}
	}
	return top == -1
}
```

**复杂度分析：**
- 时间复杂度：O(n)，其中 n 是字符串的长度
- 空间复杂度：O(n)，最坏情况下所有字符都是左括号
