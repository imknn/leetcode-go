package string

import "math"

// 8. 字符串转换整数 (atoi)
// 模拟实现 atoi 函数，处理前导空格、符号、数字和溢出
// 时间 O(n)，空间 O(1)
func myAtoi(s string) int {
	i := 0
	n := len(s)

	// 1. 跳过前导空格
	for i < n && s[i] == ' ' {
		i++
	}

	// 如果已经到达字符串末尾，返回 0
	if i == n {
		return 0
	}

	// 2. 处理符号
	sign := 1
	if s[i] == '+' {
		i++
	} else if s[i] == '-' {
		sign = -1
		i++
	}

	// 3. 读取数字并处理溢出
	result := 0
	for i < n && s[i] >= '0' && s[i] <= '9' {
		digit := int(s[i] - '0')

		// 检查溢出：result * 10 + digit > math.MaxInt32
		if result > (math.MaxInt32-digit)/10 {
			if sign == 1 {
				return math.MaxInt32
			}
			return math.MinInt32
		}

		result = result*10 + digit
		i++
	}

	return sign * result
}
