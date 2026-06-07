# 125. 验证回文串

## 题目信息

| 属性 | 值 |
|------|-----|
| 题号 | [125. 验证回文串](https://leetcode.cn/problems/valid-palindrome/) |
| 难度 | 简单 |
| 分类 | string |
| 标签 | 双指针、字符串 |
| 文件 | `string/code_125.go` |

## 题目描述

如果在将所有大写字符转换为小写字符、并移除所有非字母数字字符之后，短语正着读和反着读都一样。则可以认为该字符串是一个回文串。

给定一个字符串 `s`，返回 `true` 如果它是回文的，或 `false` 否则。

**示例 1：**
```
输入：s = "A man, a plan, a canal: Panama"
输出：true
解释："amanaplanacanalpanama" 是回文串。
```

**示例 2：**
```
输入：s = "race a car"
输出：false
解释："raceacar" 不是回文串。
```

**示例 3：**
```
输入：s = " "
输出：true
解释：在移除所有非字母数字字符之后，字符串为空，是回文串。
```

**提示：**
- `1 <= s.length <= 2 * 10^5`
- `s` 仅可打印的 ASCII 字符

## 解题算法

### 方法：预处理 + 双指针

**核心思路：**
先预处理字符串，只保留字母和数字字符并统一转为小写，然后使用双指针从两端向中间比较，判断是否为回文。

**代码实现：**
```go
func isPalindrome(s string) bool {
	t := make([]byte, 0)
	for _, c := range s {
		if c >= 'A' && c <= 'Z' {
			c = c - 'A' + 'a'
			t = append(t, byte(c))
		} else if (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') {
			t = append(t, byte(c))
		}
	}
	for l, i := len(t), 0; i < l/2; i++ {
		if t[i] != t[l-1-i] {
			return false
		}
	}
	return true
}
```

**复杂度分析：**
- 时间复杂度：O(n)，遍历字符串进行预处理和双指针比较
- 空间复杂度：O(n)，存储预处理后的新字符串
