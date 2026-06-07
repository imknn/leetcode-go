# 242. 有效的字母异位词

## 题目信息

| 属性 | 值 |
|------|-----|
| 题号 | [242. 有效的字母异位词](https://leetcode.cn/problems/valid-anagram/) |
| 难度 | 简单 |
| 分类 | string |
| 标签 | 哈希表、字符串、排序 |
| 文件 | `string/code_242.go` |

## 题目描述

给定两个字符串 `s` 和 `t`，判断 `t` 是否是 `s` 的字母异位词。

字母异位词是指两个字符串包含相同的字符，且每个字符出现的次数也相同（不考虑顺序）。

**示例 1：**
```
输入：s = "anagram", t = "nagaram"
输出：true
```

**示例 2：**
```
输入：s = "rat", t = "car"
输出：false
```

**提示：**
- `1 <= s.length, t.length <= 5 * 10^4`
- `s` 和 `t` 仅包含小写字母

## 解题算法

### 方法：计数数组

**核心思路：**
使用一个长度为 26 的数组统计字符频率。遍历 s 时对每个字符计数加 1，遍历 t 时对每个字符计数减 1。最后检查数组中所有值是否为 0，若全为 0 则是字母异位词。

**代码实现：**
```go
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := make([]int, 26)
	for _, c := range s {
		m[c-'a']++
	}
	for _, c := range t {
		m[c-'a']--
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}
```

**复杂度分析：**
- 时间复杂度：O(n)，其中 n 是字符串长度
- 空间复杂度：O(1)，固定大小为 26 的计数数组
