# 1185. 一周中的第几天

## 题目信息

| 属性 | 值 |
|------|-----|
| 题号 | [1185. 一周中的第几天](https://leetcode.cn/problems/day-of-the-week/) |
| 难度 | Easy |
| 分类 | math |
| 标签 | 数学 |
| 文件 | `math/code_1185.go` |

## 题目描述

给你一个日期，请你用字符串表示该日期是星期几。

**示例 1：**
```
输入：day = 31, month = 8, year = 2019
输出："Saturday"
```

**示例 2：**
```
输入：day = 18, month = 7, year = 1999
输出："Sunday"
```

**示例 3：**
```
输入：day = 15, month = 8, year = 1993
输出："Sunday"
```

## 解题算法

### 方法：基姆拉尔森计算公式

**核心思路：**
使用基姆拉尔森计算公式（Kim Larsen calculation formula）直接计算给定日期是星期几。该公式考虑了年、月、日以及闰年的修正。对于 1 月和 2 月，将其视为上一年的 13 月和 14 月来处理，以简化闰年判断的逻辑。

**代码实现：**
```go
func dayOfTheWeek(day int, month int, year int) string {
	week := [...]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	if month <= 2 {
		month += 12
		year--
	}
	t := (day + 2*month + 3*(month+1)/5 + year + year/4 - year/100 + year/400 + 1) % 7
	return week[t]
}
```

**复杂度分析：**
- 时间复杂度：O(1)，只涉及常数次算术运算
- 空间复杂度：O(1)，只使用了常数个额外变量
