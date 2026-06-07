# 58. 最后一个单词的长度

## 题目信息

| 属性 | 值 |
|------|-----|
| 题号 | [58. 最后一个单词的长度](https://leetcode.cn/problems/length-of-last-word/) |
| 难度 | 简单 |
| 分类 | string |
| 标签 | 字符串 |
| 文件 | `string/code_58.go` |

## 题目描述

给你一个字符串 `s`，由若干单词组成，单词前后用一些空格隔开。返回字符串中最后一个单词的长度。

如果不存在最后一个单词，请返回 `0`。

单词是指仅由字母（不包含空格）组成的非空字符串。

**示例 1：**
```
输入：s = "Hello World"
输出：5
```

**示例 2：**
```
输入：s = "   fly me   to   the moon  "
输出：4
```

**示例 3：**
```
输入：s = "luffy is still joyboy"
输出：6
```

**提示：**
- `1 <= s.length <= 10^4`
- `s` 仅有英文字母和空格 `' '` 组成
- `s` 中至少存在一个单词

## 解题算法

### 方法：从尾部倒序遍历

**核心思路：**
从字符串末尾向前遍历，跳过尾部空格，然后统计连续非空字符的个数，即为最后一个单词的长度。

**代码实现：**
```go
func lengthOfLastWord(s string) int {
	count := 0
	length := len(s)
	for i := length - 1; i >= 0; i-- {
		if s[i] != ' ' {
			count++
		} else if count > 0 {
			return count
		}
	}
	return count
}
```

**复杂度分析：**
- 时间复杂度：O(n)，其中 n 是字符串长度，最多遍历一次
- 空间复杂度：O(1)，只使用了常数个变量
