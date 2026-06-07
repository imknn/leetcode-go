package math

import (
	"math"
)

/* 7. 整数反转 */

func reverse(x int) int {
	// -2187483648 ~ 2187483647
	isNe, sum := int64(1), int64(0)
	if x < 0 {
		isNe = -1
		x = -x
	}
	for x > 0 {
		sum = sum*10 + int64(x%10)
		x /= 10
		if sum > math.MaxInt32 || (sum*isNe < math.MinInt32) {
			return 0
		}
	}
	return int(sum * isNe)
}
