package string

import "bytes"

/* 12. 整数转罗马数字 */
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
