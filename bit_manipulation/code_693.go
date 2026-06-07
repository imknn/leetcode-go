package bit_manipulation

/* 693. 交替位二进制数 */

func hasAlternatingBits(n int) bool {
	t := n % 2
	for n >>= 1; n > 0; n >>= 1 {
		k := n % 2
		if t == k {
			return false
		}
		t = k
	}
	return true
}
