# 38. 外观数列

## 题目信息

| 属性 | 值 |
|------|-----|
| 题号 | [38. 外观数列](https://leetcode.cn/problems/count-and-say/) |
| 难度 | 中等 |
| 分类 | string |
| 标签 | 字符串、模拟 |
| 文件 | `string/code_38.go` |

## 题目描述

外观数列是一个整数序列，其规律为：从数字 1 开始，每一项是对前一项的描述。

前几项如下：
- 1
- 11 （1 个 1）
- 21 （2 个 1）
- 1211 （1 个 2，1 个 1）
- 111221 （1 个 1，1 个 2，2 个 1）

给定一个正整数 `n`，返回外观数列的第 `n` 个元素。

**示例 1：**
```
输入：n = 1
输出："1"
```

**示例 2：**
```
输入：n = 4
输出："1211"
```

**提示：**
- `1 <= n <= 30`

## 解题算法

### 方法：模拟

**核心思路：**
从 `"1"` 开始，对当前字符串逐个字符统计连续相同字符的个数，将"数量+字符"拼接到结果中。重复此过程 n-1 次即可得到第 n 项。

**代码实现：**
```go
func countAndSay(n int) string {
	var result = "1"
	var buffer bytes.Buffer
	for ; n > 1; n-- {
		length := len(result)
		var num = 1
		for i := 0; i < length; i++ {
			if i == length-1 || result[i] != result[i+1] {
				// 开始打印
				buffer.WriteByte(byte(num + '0'))
				buffer.WriteByte(result[i])
				num = 1
			} else {
				num++
			}
		}
		result = buffer.String()
		buffer.Reset()
	}
	return result
}
```

**复杂度分析：**
- 时间复杂度：O(n * m)，其中 n 是给定的项数，m 是每一步字符串的长度（m 增长约 2 倍）
- 空间复杂度：O(m)，用于存储当前项和下一项的字符串
