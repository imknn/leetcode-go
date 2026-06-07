# 290. 单词规律

## 题目信息

| 属性 | 值 |
|------|-----|
| 题号 | [290. 单词规律](https://leetcode.cn/problems/word-pattern/) |
| 难度 | 简单 |
| 分类 | string |
| 标签 | 哈希表、字符串 |
| 文件 | `string/code_290.go` |

## 题目描述

给定一种规律 `pattern` 和一个字符串 `str`，判断 `str` 是否遵循相同的规律。

遵循指完全匹配，即 `pattern` 中的每个字符与 `str` 中的每个单词之间存在双射关系。

**示例 1：**
```
输入：pattern = "abba", str = "dog cat cat dog"
输出：true
```

**示例 2：**
```
输入：pattern = "abba", str = "dog cat cat fish"
输出：false
```

**示例 3：**
```
输入：pattern = "aaaa", str = "dog cat cat dog"
输出：false
```

**提示：**
- `1 <= pattern.length <= 300`
- `pattern` 仅包含小写英文字母
- `1 <= str.length <= 3000`
- `str` 仅由小写英文字母和空格组成
- `str` 中每个单词之间恰好有一个空格，且 `str` 的开头和结尾没有空格

## 解题算法

### 方法：双向哈希映射

**核心思路：**
与同构字符串类似，维护两个映射：字符到单词的映射和单词到字符的映射。遍历时同时检查双向映射的一致性，如果不匹配则返回 false。

**代码实现：**
```go
func wordPattern(pattern string, str string) bool {
	arr := strings.Split(str, " ")
	len1, len2 := len(pattern), len(arr)
	if len1 != len2 {
		return false
	}
	m := make(map[int32]string)
	n := make(map[string]int32)
	for idx, ch := range pattern {
		if item, ok := m[ch]; ok {
			if item != arr[idx] {
				return false
			}
		} else {
			m[ch] = arr[idx]
		}
		if item, ok := n[arr[idx]]; ok {
			if item != ch {
				return false
			}
		} else {
			n[arr[idx]] = ch
		}
	}
	return true
}
```

**复杂度分析：**
- 时间复杂度：O(n)，其中 n 是 pattern 的长度
- 空间复杂度：O(n)，存储两个映射表
