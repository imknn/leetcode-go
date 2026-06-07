# 43. 字符串相乘

## 题目信息

| 属性 | 值 |
|------|-----|
| 题号 | [43. 字符串相乘](https://leetcode.cn/problems/multiply-strings/) |
| 难度 | 中等 |
| 分类 | string |
| 标签 | 数学、字符串、模拟 |
| 文件 | `string/code_43.go` |

## 题目描述

给定两个以字符串形式表示的非负整数 `num1` 和 `num2`，返回 `num1` 和 `num2` 的乘积，也用字符串表示。

你不能使用任何内置的 BigInteger 库或直接将输入转换为整数。

**示例 1：**
```
输入：num1 = "2", num2 = "3"
输出："6"
```

**示例 2：**
```
输入：num1 = "123", num2 = "456"
输出："56088"
```

**提示：**
- `1 <= num1.length, num2.length < 200`
- `num1` 和 `num2` 仅包含数字 `0-9`
- `num1` 和 `num2` 均不以零开头，除非该值本身为 0

## 解题算法

### 方法：模拟竖式乘法

**核心思路：**
模拟小学竖式乘法：用 `num2` 的每一位去乘以整个 `num1`，每次结果后面补相应数量的 0，然后将所有中间结果逐位相加。其中"单数字乘以字符串"和"字符串相加"分别由辅助函数实现。

**代码实现：**
```go
func multiply(num1 string, num2 string) string {
	s := ""
	for i, zero := len(num2)-1, 0; i >= 0; i, zero = i-1, zero+1 {
		r := multiplyOne(num1, int(num2[i]-'0'))
		s = sum(s, r+strings.Repeat("0", zero))
	}
	// 处理 1000 * 0 出现 000 的问题
	findUnZero := false
	var buffer bytes.Buffer
	for i := 0; i < len(s); i++ {
		if findUnZero {
			buffer.WriteByte(s[i])
		} else if s[i] != '0' {
			findUnZero = true
			buffer.WriteByte(s[i])
		}
	}
	if !findUnZero {
		buffer.WriteByte('0')
	}
	return buffer.String()
}

// 长字符串乘以一个数字
func multiplyOne(num1 string, num int) string {
	if num == 0 {
		return "0"
	} else if num == 1 {
		return num1
	}
	var buffer bytes.Buffer
	carry := 0 // 进位
	for i := len(num1) - 1; i >= 0; i-- {
		result := int(num1[i]-'0')*num + carry
		carry = result / 10
		result = result % 10
		buffer.WriteByte(byte(result + '0'))
	}
	if carry > 0 {
		buffer.WriteByte(byte(carry + '0'))
	}
	return common.Reverse(&buffer)
}

// 字符串相加
func sum(num1 string, num2 string) string {
	m, n := len(num1), len(num2)
	i := m
	if m < n { // 取大的
		i = n
	}
	var buffer bytes.Buffer
	carry := 0 // 进位
	for x := 0; x < i; x++ {
		if m-x-1 < 0 {
			n2 := int(num2[n-x-1] - '0')
			buffer.WriteByte(byte((n2+carry)%10 + '0'))
			carry = int(n2+carry) / 10
		} else if n-x-1 < 0 {
			n1 := int(num1[m-x-1] - '0')
			buffer.WriteByte(byte((n1+carry)%10 + '0'))
			carry = int(n1+carry) / 10
		} else {
			n1 := int(num1[m-x-1] - '0')
			n2 := int(num2[n-x-1] - '0')
			buffer.WriteByte(byte((n1+n2+carry)%10 + '0'))
			carry = int(n1+n2+carry) / 10
		}
	}
	if carry > 0 {
		buffer.WriteByte(byte(carry + '0'))
	}
	return common.Reverse(&buffer)
}
```

**复杂度分析：**
- 时间复杂度：O(m * n)，其中 m 和 n 分别是两个字符串的长度
- 空间复杂度：O(m + n)，存储中间结果和最终结果
