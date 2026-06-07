package string

import (
	"strings"
)

/* 13. 罗马数字转整数*/
func romanToInt(s string) int {
	number := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	roman := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	var num int
	for idx := 0; len(s) > 0 && idx < 13; {
		if strings.HasPrefix(s, roman[idx]) {
			num = num + number[idx]
			s = s[len(roman[idx]):]
		} else {
			idx++
		}
	}
	return num
}
