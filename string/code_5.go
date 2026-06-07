package string

//
//// f[i,j] = f[i+1, j-1] and si == sj
//
//func longestPalindrome(s string) string {
//	ret, max := s[0:1], 1
//	length := len(s)
//	var result = make([]bool, length*length) // 二维转一维
//	// 坐标是  i*length+1 i*length+j
//	for i := 0; i < length; i++ {
//		result[i*length+i] = true // 回文串长度1
//		if i < length-1 { // 回文串长度2
//			if s[i] == s[i+1] {
//				if max < 2 {
//					max = 2
//					ret = s[i : i+max]
//				}
//				result[i*length+i+1] = true
//			} else {
//				result[i*length+i+1] = false
//			}
//		}
//		for j := 2; j < length; j++ { // j 表示回文串的长度
//			idx:= i*length+j
//			result[idx] = result[(i+1)*length+(j-1)] && s[i] == s[j]
//			if result[idx] && (j-i + 1) > max{
//				max = j-i+1
//
//			}
//
//		}
//	}
//	return ""
//}
