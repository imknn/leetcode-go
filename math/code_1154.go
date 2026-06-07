package math

import (
	"strconv"
)

/* 1154. 一年中的第几天  */
func dayOfYear(date string) int {
	year, _ := strconv.Atoi(date[:4])
	month, _ := strconv.Atoi(date[5:7])
	day, _ := strconv.Atoi(date[8:10])
	count := [...]int{0, 31, 59, 90, 120, 151, 181, 212, 243, 273, 304, 334}
	// 判断闰年
	if ((year%4 == 0 && year%100 != 0) || year%400 == 0) && month >= 3 {
		return count[month-1] + day + 1
	}
	return count[month-1] + day
}
