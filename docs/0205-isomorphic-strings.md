# 205. 同构字符串

## 题目信息

| 属性 | 值 |
|------|-----|
| 题号 | [205. 同构字符串](https://leetcode.cn/problems/isomorphic-strings/) |
| 难度 | 简单 |
| 分类 | string |
| 标签 | 哈希表、字符串 |
| 文件 | `string/code_205.go` |

## 题目描述

给定两个字符串 `s` 和 `t`，判断它们是否是同构的。

如果 `s` 中的字符可以被替换得到 `t`，则两个字符串是同构的。所有出现的字符必须用另一个字符替换，同时保留字符的顺序。没有两个字符可以映射到同一个字符，但一个字符可以映射到自身。

**示例 1：**
```
输入：s = "egg", t = "add"
输出：true
```

**示例 2：**
```
输入：s = "foo", t = "bar"
输出：false
```

**示例 3：**
```
输入：s = "paper", t = "title"
输出：true
```

**提示：**
- `1 <= s.length <= 5 * 10^4`
- `t.length == s.length`
- `s` 和 `t` 由任意有效的 ASCII 字符组成

## 解题算法

### 方法：双向哈希映射

**核心思路：**
维护两个映射关系：s->t 和 t->s。遍历字符串时，如果 s[i] 已经映射到其他字符或 t[i] 已经被其他字符映射，则不是同构字符串。双向映射确保一一对应。

**代码实现：**
```go
func isIsomorphic(s string, t string) bool {
	var l = len(s)
	if l != len(t) {
		return false
	}
	var m = make(map[byte]byte)
	var n = make(map[byte]byte)
	for i := 0; i < l; i++ {
		if _, ok := m[s[i]]; ok && m[s[i]] != t[i] {
			return false
		}
		if _, ok := n[t[i]]; ok && n[t[i]] != s[i] {
			return false
		}
		m[s[i]] = t[i]
		n[t[i]] = s[i]
	}
	return true
}
```

**复杂度分析：**
- 时间复杂度：O(n)，遍历一次字符串
- 空间复杂度：O(字符集大小)，存储两个映射表
