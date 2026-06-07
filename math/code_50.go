package math

func myPow(x float64, n int) float64 {
	var i, total = n, float64(1)
	for i != 0 {
		if i&1 == 1 {
			total *= x
		}
		x *= x
		i /= 2
	}
	if n > 0 {
		return total
	} else {
		return 1 / total
	}
}
