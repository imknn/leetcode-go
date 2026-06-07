# 345. 反转字符串中的元音字母

## 题目信息

| 属性 | 值 |
|------|-----|
| 题号 | [345. 反转字符串中的元音字母](https://leetcode.cn/problems/reverse-vowels-of-a-string/) |
| 难度 | 简单 |
| 分类 | string |
| 标签 | 双指针、字符串 |
| 文件 | `string/code_345.go` |

## 题目描述

给你一个字符串 `s`，仅反转字符串中的所有元音字母，并返回结果字符串。

元音字母包括 `'a'`、`'e'`、`'i'`、`'o'`、`'u'`，且可能以大写和小写两种形式出现。

**示例 1：**
```
输入：s = "hello"
输出："holle"
```

**示例 2：**
```
输入：s = "leetcode"
输出："leotcede"
```

**提示：**
- `1 <= s.length <= 3 * 10^5`
- `s` 由可打印的 ASCII 字符组成

## 解题算法

### 方法：双指针

**核心思路：**
使用左右两个指针从字符串两端向中间移动，左指针找到第一个元音字母，右指针找到从右边数第一个元音字母，交换两个元音字母，然后继续向中间移动，直到两指针相遇。

**代码实现：**
```go
func reverseVowels(s string) string {
	length := len(s)
	b := []byte(s)
	x, y := 0, length-1
	for {
		for x < y {
			if b[x] == 'a' || b[x] == 'e' || b[x] == 'i' || b[x] == 'o' || b[x] == 'u' ||
				b[x] == 'A' || b[x] == 'E' || b[x] == 'I' || b[x] == 'O' || b[x] == 'U' {
				break
			} else {
				x++
			}
		}
		for x < y {
			if b[y] == 'a' || b[y] == 'e' || b[y] == 'i' || b[y] == 'o' || b[y] == 'u' ||
				b[y] == 'A' || b[y] == 'E' || b[y] == 'I' || b[y] == 'O' || b[y] == 'U' {
				break
			} else {
				y--
			}
		}
		if x < y {
			b[x], b[y] = b[y], b[x]
			x++
			y--
		} else {
			break
		}
	}
	return string(b)
}
```

**复杂度分析：**
- 时间复杂度：O(n)，左右指针各遍历一次字符串
- 空间复杂度：O(n)，转换为字节数组进行操作
