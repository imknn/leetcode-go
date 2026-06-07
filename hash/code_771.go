package hash

/* 771. 宝石与石头*/
func numJewelsInStones(J string, S string) int {
	num := 0
	m := make(map[int32]bool)
	for _, ch := range J {
		m[ch] = true
	}
	for _, ch := range S {
		if m[ch] {
			num++
		}
	}
	return num
}
