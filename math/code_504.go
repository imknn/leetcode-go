package math

import (
	"bytes"
)

/* 504. 七进制数 */
func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}
	isNe := false
	if num < 0 {
		num = -num
		isNe = true
	}
	var buffer bytes.Buffer
	for num > 0 {
		buffer.WriteByte(byte('0' + (num % 7)))
		num /= 7
	}
	if isNe {
		buffer.WriteByte('-')
	}
	// 逆转
	s := buffer.Bytes()
	length := len(s)
	half := length / 2
	for i := 0; i < half; i++ {
		s[i], s[length-1-i] = s[length-1-i], s[i]
	}
	return buffer.String()
}
