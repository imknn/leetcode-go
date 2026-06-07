package bit_manipulation

// 371. 两整数之和
// 不使用 + 和 - 运算符计算两数之和
// 核心思路：异或得无进位和，与运算左移得进位，循环直到进位为0

const mask = 0xFFFFFFFF

func getSum(a int, b int) int {
	for b != 0 {
		carry := (a & b) & mask // 进位（只保留低32位）
		a = (a ^ b) & mask     // 无进位和（只保留低32位）
		b = carry << 1         // 进位左移
	}
	// 处理负数：如果第32位是1，需要扩展符号位
	if a > 0x7FFFFFFF {
		a = ^((^a) & mask) // 转为 Go 的 64 位负数
	}
	return a
}
