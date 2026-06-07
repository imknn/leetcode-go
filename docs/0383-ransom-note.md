# 383. 赎金笔记

## 题目信息

| 属性 | 值 |
|------|-----|
| 题号 | [383. 赎金笔记](https://leetcode.cn/problems/ransom-note/) |
| 难度 | 简单 |
| 分类 | string |
| 标签 | 哈希表、字符串、计数 |
| 文件 | `string/code_383.go` |

## 题目描述

给你两个字符串 `ransomNote` 和 `magazine`。判断 `ransomNote` 能否由 `magazine` 构成。

如果可以，返回 `true`；否则返回 `false`。

`ransomNote` 中的每个字母在 `magazine` 中只能使用一次。

**示例 1：**
```
输入：ransomNote = "a", magazine = "b"
输出：false
```

**示例 2：**
```
输入：ransomNote = "aa", magazine = "ab"
输出：false
```

**示例 3：**
```
输入：ransomNote = "aa", magazine = "aab"
输出：true
```

**提示：**
- `1 <= ransomNote.length, magazine.length <= 10^5`
- `ransomNote` 和 `magazine` 仅由小写英文字母组成

## 解题算法

### 方法：计数数组

**核心思路：**
用一个长度为 26 的数组统计 magazine 中每个字母的出现次数。然后遍历 ransomNote，如果某个字母的剩余计数不足（为 0），则无法构成，返回 false。

**代码实现：**
```go
func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}
	m := make([]int, 26)
	for _, ch := range magazine {
		m[ch-'a']++
	}
	for _, ch := range ransomNote {
		if m[ch-'a'] > 0 {
			m[ch-'a']--
		} else {
			return false
		}
	}
	return true
}
```

**复杂度分析：**
- 时间复杂度：O(n + m)，其中 n 和 m 分别是 ransomNote 和 magazine 的长度
- 空间复杂度：O(1)，固定大小为 26 的计数数组
