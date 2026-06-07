# 12. 整数转罗马数字

## 题目信息

| 属性 | 值 |
|------|-----|
| 题号 | [12. 整数转罗马数字](https://leetcode.cn/problems/integer-to-roman/) |
| 难度 | 中等 |
| 分类 | string |
| 标签 | 哈希表、数学、字符串 |
| 文件 | `string/code_12.go` |

## 题目描述

七个不同的符号代表罗马数字：`I(1)`, `V(5)`, `X(10)`, `L(50)`, `C(100)`, `D(500)`, `M(1000)`。

罗马数字通过从大到小组合符号来表示，包含减法规则如 `IV = 4`, `IX = 9`, `XL = 40`, `XC = 90`, `CD = 400`, `CM = 900`。

给定一个整数，将其转换为罗马数字表示。

**示例 1：**
```
输入：num = 3
输出："III"
```

**示例 2：**
```
输入：num = 58
输出："LVIII"
```

**示例 3：**
```
输入：num = 1994
输出："MCMXCIV"
```

**提示：**
- `1 <= num <= 3999`

## 解题算法

### 方法：贪心

**核心思路：**
将所有可能的罗马数字（包括减法规则的组合）从大到小排列。贪心地从最大值开始，如果当前数字大于等于某个罗马值，则写入对应的符号并减去该值，直到数字变为 0。

**代码实现：**
```go
func intToRoman(num int) string {
	number := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	roman := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	var buffer bytes.Buffer
	for idx := 0; num > 0 && idx < 13; {
		if num >= number[idx] {
			buffer.WriteString(roman[idx])
			num = num - number[idx]
		} else {
			idx++
		}
	}
	return buffer.String()
}
```

**复杂度分析：**
- 时间复杂度：O(1)，最多循环 13 次（固定大小的映射表）
- 空间复杂度：O(1)，只使用缓冲区存储结果
