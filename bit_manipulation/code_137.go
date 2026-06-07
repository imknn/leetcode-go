package bit_manipulation

/* 137. 只出现一次的数字 II
1. 方法一： 将数字转成32位二进制，统计各个位上1的数量，不是3的倍数的列写1，其他列写0，就找到了这个出现 1 次的数。
2. 方法二：使用两个变量统计每一位的数量，n1,n0 ,如果n1=1,n0=0,代表10，即有两位相同。具体表示：00 (0或3位相同)，01(1位相同) 10 (2位相同)
3. 方法三：使用map统计每个num的数量
*/

func singleNumber3(nums []int) int {
	m, n := make(map[int]int), nums[0]
	for _, num := range nums {
		if _, ok := m[num]; !ok {
			m[num] = 1
		} else {
			m[num]++
		}
	}
	for k, v := range m {
		if v == 1 {
			return k
		}
	}
	return n
}

//func singleNumber2(nums []int) int {
//	n1, n0 := 0, 0
//	for _, num := range nums {
//		n1 = (n1 ^ num) & ^n0
//		n0 = (n0 ^ num) & ^n1
//	}
//	return n1
//}
