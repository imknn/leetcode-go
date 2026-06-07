package string

/* 859. 亲密字符串 */
func buddyStrings(A string, B string) bool {
	diff := make([]int, 0)
	m := len(A)
	if m != len(B) {
		return false
	}
	check, hasSame := 0, false
	for i := 0; i < m; i++ {
		if A[i] != B[i] {
			diff = append(diff, i)
		}
		if !hasSame && ((check>>(A[i]-'a'))&1) == 1 {
			hasSame = true
		}
		if !hasSame {
			check = check | (1 << (A[i] - 'a'))
		}
	}
	n := len(diff)
	if n == 0 {
		return hasSame
	} else if n == 2 {
		return A[diff[0]] == B[diff[1]] && A[diff[1]] == B[diff[0]]
	}
	return false
}
