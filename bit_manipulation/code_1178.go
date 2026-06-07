package bit_manipulation

/* 1178. 猜字谜  */
func findNumOfValidWords(words []string, puzzles []string) []int {
	// 使用位标记的方式设置字符串中包含的字符
	wordsX := make([]uint32, len(words))
	for i, word := range words {
		for _, ch := range word {
			wordsX[i] = wordsX[i] | 1<<(uint32)(ch-'a')
		}
		//fmt.Printf("%b\n", wordsX[i])
	}
	puzzlesX := make([]uint32, len(puzzles))
	for i, puzzle := range puzzles {
		for _, ch := range puzzle {
			puzzlesX[i] = puzzlesX[i] | 1<<(uint32)(ch-'a')
		}
		//fmt.Printf("%b\n", puzzlesX[i])
	}
	answer := make([]int, len(puzzles))
	for i, puzzle := range puzzles {
		answer[i] = 0
		for j, _ := range words {
			// 1. 判断第一个字符
			// 2. 判断字符串的包含关系(异或^检查差异，&检查差异的字符是否属于word)
			if (1<<(uint32)(puzzle[0]-'a'))&wordsX[j] == 0 || wordsX[j]^puzzlesX[i]&wordsX[j] != 0 {
				continue
			}
			answer[i]++
		}
	}
	return answer
}
