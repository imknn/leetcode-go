package bit_manipulation

/* 338. 比特位计数 */

func countBits(num int) []int {
	result, t := make([]int, num+1), 0
	// 判断上一个数字的最后一位
	// 如果是0，那就 + 1
	// 如果是0， 向前遍历每一位，找到1就-1，直到找到0 +1 结束
	for i := 1; i <= num; i++ {
		result[i] = result[i-1]
		for j := 0; j < 32; j++ {
			if t&1 == 1 {
				result[i]--
				t >>= 1
			} else {
				result[i]++
				break
			}
		}
		t = i
	}
	return result
}
