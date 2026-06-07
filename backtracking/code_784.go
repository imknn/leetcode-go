package backtracking

/* 784. 字母大小写全排列 */
//func letterCasePermutation(S string) []string {
//	result := make([]string, 0)
//	for _, s := range S {
//		if s >= 'a' && s <= 'z' {
//			result = append(result, "")
//			appendResult(result, s)
//			appendResult(result, s + 32)
//		} else if s >= 'A' && s <= 'Z' {
//			appendResult(result, s)
//			result = append(result, "")
//			appendResult(result, s + 32)
//		} else {
//			buffer.WriteByte(byte(s))
//		}
//		result = append(result, buffer.String())
//	}
//	return result
//}
//
//func all(result []string, s int32) []string {
//	for idx, s := range result {
//		result[idx] = s + "" // 5465   5645
//
//	}
//	return result
//}
