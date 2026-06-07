package string

import (
	"bytes"
	"izqy.top/common"
	"strings"
)

// 43. 字符串相乘
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
