package math

import "izqy.top/common"

// 914. 卡牌分组
// 这是一道计算最大公约数的题目，可以使用 辗转相除法
func hasGroupsSizeX(deck []int) bool {
	m := make(map[int]int)
	for _, v := range deck {
		if _, ok := m[v]; ok {
			m[v] = m[v] + 1
		} else {
			m[v] = 1
		}
	}
	a := 0
	for _, v := range m {
		if a == 0 {
			a = v
			continue
		}
		a = common.Gcd(a, v)
		if a < 2 {
			return false
		}
	}
	return a > 1
}
