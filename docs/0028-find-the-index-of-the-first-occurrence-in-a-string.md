# 28. 找出字符串中第一个匹配项的下标

## 题目信息

| 属性 | 值 |
|------|-----|
| 题号 | [28. 找出字符串中第一个匹配项的下标](https://leetcode.cn/problems/find-the-index-of-the-first-occurrence-in-a-string/) |
| 难度 | 简单 |
| 分类 | string |
| 标签 | 双指针、字符串、字符串匹配 |
| 文件 | `string/code_28.go` |

## 题目描述

给你两个字符串 `needle` 和 `haystack`，请你在 `haystack` 字符串中找出 `needle` 字符串的第一个匹配项的下标。如果 `needle` 不是 `haystack` 的一部分，则返回 `-1`。

**示例 1：**
```
输入：haystack = "sadbutsad", needle = "sad"
输出：0
解释："sad" 在下标 0 和 6 处匹配。第一个匹配项的下标是 0。
```

**示例 2：**
```
输入：haystack = "leetcode", needle = "leeto"
输出：-1
```

**提示：**
- `1 <= haystack.length, needle.length <= 10^4`
- `haystack` 和 `needle` 仅由小写英文字符组成

## 解题算法

### 方法：朴素字符串匹配

**核心思路：**
使用双指针逐字符比较：在 haystack 中滑动匹配 needle。当发现不匹配时，将 i 指针回退已匹配的步数，j 指针重置为 0，继续尝试。当 j 走完 needle 时，返回起始位置。

**代码实现：**
```go
func strStr(haystack string, needle string) int {
	var l1 = len(haystack)
	var l2 = len(needle)
	if l2 == 0 {
		return 0
	}
	for i, j := 0, 0; l1 >= l2 && i < l1 && j < l2; i++ {
		if haystack[i] == needle[j] {
			if j == l2-1 {
				return i - l2 + 1
			}
			j++
		} else {
			i = i - j // i 回退已经相同的步数
			j = 0     // j 重置
		}
	}
	return -1
}
```

**复杂度分析：**
- 时间复杂度：O(n * m)，其中 n 是 haystack 长度，m 是 needle 长度
- 空间复杂度：O(1)，只使用了常数个变量
