package bit_manipulation

/* 136. 只出现一次的数字
- 交换律：a ^ b ^ c <=> a ^ c ^ b
- 任何数于0异或为任何数 0 ^ n => n
- 相同的数异或为0: n ^ n => 0
*/
func singleNumber(nums []int) int {
	n, length := nums[0], len(nums)
	for i := 1; i < length; i++ {
		n ^= nums[i]
	}
	return n
}
