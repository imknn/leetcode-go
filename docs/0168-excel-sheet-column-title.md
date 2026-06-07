# 168. Excel 表列名称

## 题目信息

| 属性 | 值 |
|------|-----|
| 题号 | [168. Excel 表列名称](https://leetcode.cn/problems/excel-sheet-column-title/) |
| 难度 | Easy |
| 分类 | math |
| 标签 | 数学、字符串 |
| 文件 | `math/code_168.go` |

## 题目描述

给你一个整数 columnNumber，返回它在 Excel 表中相对应的列名称。

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
输入：columnNumber = 1
输出："A"
```

**示例 2：**
```
输入：columnNumber = 28
输出："AB"
```

**示例 3：**
```
输入：columnNumber = 701
输出："ZY"
```

## 解题算法

### 方法：进制转换（26进制）

**核心思路：**
将列号转换为 26 进制表示。与普通进制转换不同的是，Excel 列名没有 0 的概念（A=1 而不是 A=0），因此每次取模前先将列号减 1。取模得到当前位的字符，除以 26 后继续处理高位。最后将结果反转得到正确顺序。

**代码实现：**
```go
func convertToTitle(columnNumber int) string {
	var buffer bytes.Buffer
	for columnNumber > 0 {
		columnNumber--
		buffer.WriteByte(byte('A' + (columnNumber % 26)))
		columnNumber /= 26
	}

	common.Reverse(&buffer)
	return buffer.String()
}
```

**复杂度分析：**
- 时间复杂度：O(log₂₆(n))，循环次数与列号的 26 进制位数成正比
- 空间复杂度：O(log₂₆(n))，用于存储结果字符串
