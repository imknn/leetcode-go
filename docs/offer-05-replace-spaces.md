# 剑指 Offer 05. 替换空格

## 题目信息

| 属性 | 值 |
|------|-----|
| 题号 | [剑指 Offer 05. 替换空格](https://leetcode.cn/problems/ti-huan-kong-ge-lcof/) |
| 难度 | Easy |
| 分类 | other |
| 标签 | string |
| 文件 | `other/code_offer_05.go` |

## 题目描述

请实现一个函数，把字符串 `s` 中的每个空格替换成 `%20`。

**示例 1：**

```
输入：s = "We are happy."
输出："We%20are%20happy."
```

## 解题算法

### 方法：遍历替换

**核心思路：**
遍历字符串中的每个字符，如果遇到空格就向缓冲区写入 `%20`，否则直接写入原字符。使用 `bytes.Buffer` 高效地拼接字符串，避免了字符串拼接的性能开销。

**代码实现：**
```go
func replaceSpace(s string) string {
	var buffer bytes.Buffer
	for _, ch := range s {
		if ch == ' ' {
			buffer.WriteString("%20")
		} else {
			buffer.WriteByte(byte(ch))
		}
	}
	return buffer.String()
}
```

**复杂度分析：**
- 时间复杂度：O(n)，遍历字符串一次
- 空间复杂度：O(n)，最坏情况下所有字符都是空格，结果字符串长度为 3n
