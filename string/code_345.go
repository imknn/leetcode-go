package string

func reverseVowels(s string) string {
	length := len(s)
	b := []byte(s)
	x, y := 0, length-1
	for {
		for x < y {
			if b[x] == 'a' || b[x] == 'e' || b[x] == 'i' || b[x] == 'o' || b[x] == 'u' ||
				b[x] == 'A' || b[x] == 'E' || b[x] == 'I' || b[x] == 'O' || b[x] == 'U' {
				break
			} else {
				x++
			}
		}
		for x < y {
			if b[y] == 'a' || b[y] == 'e' || b[y] == 'i' || b[y] == 'o' || b[y] == 'u' ||
				b[y] == 'A' || b[y] == 'E' || b[y] == 'I' || b[y] == 'O' || b[y] == 'U' {
				break
			} else {
				y--
			}
		}
		if x < y {
			b[x], b[y] = b[y], b[x]
			x++
			y--
		} else {
			break
		}
	}
	return string(b)
}
