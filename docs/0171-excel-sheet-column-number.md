# 171. Excel 表列序号

## 题目信息

| 属性 | 值 |
|------|-----|
| 题号 | [171. Excel 表列序号](https://leetcode.cn/problems/excel-sheet-column-number/) |
| 难度 | Easy |
| 分类 | math |
| 标签 | 数学、字符串 |
| 文件 | `math/code_171.go` |

## 题目描述

给定一个表示 Excel 表中列名称的字符串，返回其对应的列号。

例如：
```
A -> 1
B -> 2
C -> 3
...
Z -> 26
AA -> 27
AB -> 28
...
```

**示例 1：**
```
输入：columnTitle = "A"
输出：1
```

**示例 2：**
```
输入：columnTitle = "AB"
输出：28
```

**示例 3：**
```
输入：columnTitle = "ZY"
输出：701
```

## 解题算法

### 方法：26进制转换

**核心思路：**
将字符串看作 26 进制数，从右向左遍历每一位，将字符映射为 1-26 的数字，乘以对应的权值（26 的幂次）后累加。这与第 168 题互为逆操作。

**代码实现：**
```go
func titleToNumber(s string) int {
	sum := 0
	ratio := 1
	for i := len(s) - 1; i >= 0; i-- {
		sum += int(s[i]-'A'+1) * ratio
		ratio *= 26
	}
	return sum
}
```

**复杂度分析：**
- 时间复杂度：O(n)，n 为字符串长度，只需遍历一次
- 空间复杂度：O(1)，只使用了常数个额外变量
