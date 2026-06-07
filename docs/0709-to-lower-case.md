# 709. 转换成小写字母

## 题目信息

| 属性 | 值 |
|------|-----|
| 题号 | [709. 转换成小写字母](https://leetcode.cn/problems/to-lower-case/) |
| 难度 | 简单 |
| 分类 | string |
| 标签 | 字符串 |
| 文件 | `string/code_709.go` |

## 题目描述

给你一个字符串 `s`，将该字符串中的大写字母转换成相同的小写字母，返回新的字符串。

**示例 1：**
```
输入：s = "Hello"
输出："hello"
```

**示例 2：**
```
输入：s = "here"
输出："here"
```

**示例 3：**
```
输入：s = "LOVELY"
输出："lovely"
```

**提示：**
- `1 <= s.length <= 100`
- `s` 由可打印的 ASCII 字符组成

## 解题算法

### 方法：逐字符转换

**核心思路：**
遍历字符串，对每个字符判断是否为大写字母（A-Z），如果是则转换为对应的小写字母（通过 ASCII 码差值运算），否则保持不变。

**代码实现：**
```go
func toLowerCase(str string) string {
	var buffer bytes.Buffer
	for _, ch := range str {
		if ch >= 'A' && ch <= 'Z' {
			buffer.WriteByte(byte(ch - 'A' + 'a'))
		} else {
			buffer.WriteByte(byte(ch))
		}
	}
	return buffer.String()
}
```

**复杂度分析：**
- 时间复杂度：O(n)，遍历字符串的每个字符
- 空间复杂度：O(n)，用于构建结果字符串
