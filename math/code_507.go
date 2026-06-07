package math

func checkPerfectNumber(num int) bool {
	if num < 2 {
		return false
	}
	// 先找出所有的因子
	var sum = 1
	for a := 2; a*a < num; a++ {
		if num%a == 0 {
			sum = sum + a + num/a
		}
	}
	return sum == num
}
