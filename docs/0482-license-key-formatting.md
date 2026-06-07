# 482. 密钥格式化

## 题目信息

| 属性 | 值 |
|------|-----|
| 题号 | [482. 密钥格式化](https://leetcode.cn/problems/license-key-formatting/) |
| 难度 | 中等 |
| 分类 | string |
| 标签 | 字符串 |
| 文件 | `string/code_482.go` |

## 题目描述

给定一个许可密钥字符串 `s`，仅由字母、数字和破折号 `'-'` 组成，将其重新格式化，使得每个分组恰好包含 `k` 个字符，用破折号分隔。

格式化后的密钥应全部转为大写。

**示例 1：**
```
输入：s = "5F3Z-2e-9-w", k = 4
输出："5F3Z-2E9W"
解释：分组分别是 "5F3Z" 和 "2E9W"，注意第二个分组比第一个短 1 个字符是允许的。
```

**示例 2：**
```
输入：s = "2-5g-3-J", k = 2
输出："2-5G-3J"
解释：分组分别是 "25"、"G3" 和 "J"，注意第一个分组和最后一个分组可以比其他分组短。
```

**提示：**
- `1 <= s.length <= 10^5`
- `s` 由字母、数字和破折号 `'-'` 组成
- `1 <= k <= 10^4`

## 解题算法

### 方法：去分隔符后重新格式化

**核心思路：**
先去除所有破折号，然后从后往前（或统计长度后从前往后）按每 k 个字符一组插入破折号，同时将小写字母转为大写。

**代码实现：**
```go
func licenseKeyFormatting(s string, k int) string {
	var buffer bytes.Buffer
	for _, ch := range s {
		if ch == '-' {
			continue
		}
		buffer.WriteByte(byte(ch))
	}

	i, s := buffer.Len(), buffer.String()
	first := i % k
	buffer.Reset()
	if first == 0 {
		first = k
	}
	for j := 0; j < i; j++ {
		ch := s[j]
		if ch >= 'a' && ch <= 'z' {
			ch = ch + 'A' - 'a'
		}
		buffer.WriteByte(ch)
		first--
		if first == 0 && j < i-1 {
			buffer.WriteByte('-')
			first = k
		}
	}
	return buffer.String()
}
```

**复杂度分析：**
- 时间复杂度：O(n)，其中 n 是字符串长度
- 空间复杂度：O(n)，用于存储结果字符串
