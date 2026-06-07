package bit_manipulation

/* 342. 4的幂 */

// 4的幂和2的幂一样，只会出现一位1。但是，4的1总是出现在奇数位。
// 0x5 = 0101b可以用来校验奇数位上的1。

func isPowerOfFour(num int) bool {
	if num > 0 && (num&(num-1) == 0) {
		return num&0x55555555 > 0
	}
	return false
}
