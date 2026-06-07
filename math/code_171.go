package math

func titleToNumber(s string) int {
	sum := 0
	ratio := 1
	for i := len(s) - 1; i >= 0; i-- {
		sum += int(s[i]-'A'+1) * ratio
		ratio *= 26
	}
	return sum
}
