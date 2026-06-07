# 1178. 猜字谜

## 题目信息

| 属性 | 值 |
|------|-----|
| 题号 | [1178. 猜字谜](https://leetcode.cn/problems/number-of-valid-words-for-each-puzzle/) |
| 难度 | 困难 |
| 分类 | bit_manipulation |
| 标签 | 位运算、数组、哈希表、字符串 |
| 文件 | `bit_manipulation/code_1178.go` |

## 题目描述

外国友人玩游戏，给你一个若干**唯一**单词的字典和一个猜测字符串的列表。你需要根据猜测字符串返回匹配的单词列表。每次猜测由 7 个字母组成的字符串（仅限小写字母）。

每个猜测的匹配规则：
1. 猜测字符串中的每个字母必须在单词中出现
2. 猜测字符串的**第一个字母**必须在单词中出现

返回一个答案列表，列表中的元素为满足上述条件的匹配单词。

**示例 1：**
```
输入：words = ["mobile","mouse","moneypot","monitor","mousepad"], puzzles = ["mouse","keyboard","mouse"]
输出：[3,1,3]
解释：
对于谜题 "mouse"，匹配的单词有："mobile"、"mouse"、"mousepad"。
```

**示例 2：**
```
输入：words = ["apple","pleas","please"], puzzles = ["aelwxyz","aelpxyz","aelpsxy","saelpxy","xaelpsy"]
输出：[0,1,3,2,0]
```

**提示：**
- `1 <= words.length <= 10^5`
- `4 <= words[i].length <= 50`
- `1 <= puzzles.length <= 10^4`
- `puzzles[i].length == 7`
- `words[i]` 和 `puzzles[i]` 只包含小写英文字母
- `words` 中的所有字符串互不相同

## 解题算法

### 方法：位掩码标记 + 暴力匹配

**核心思路：**
1. 将每个单词和谜题都转换为一个 26 位的位掩码（bitset），其中第 i 位表示字母表中第 i 个字母是否存在
2. 对于每个谜题，遍历所有单词，检查两个条件：
   - 单词必须包含谜题的第一个字母
   - 单词中不能包含谜题中没有的字母（通过异或和与运算检查）

**位掩码构造：**
- 对于字符串中的每个字符 `ch`，将位掩码的第 `ch-'a'` 位置为 1

**匹配检查：**
- `wordsX[j] ^ puzzlesX[i] & wordsX[j]`：检查单词中的字母是否都在谜题中
- 如果异或后与自身做与运算不为 0，说明单词中有谜题没有的字母

**代码实现：**
```go
func findNumOfValidWords(words []string, puzzles []string) []int {
	// 使用位标记的方式设置字符串中包含的字符
	wordsX := make([]uint32, len(words))
	for i, word := range words {
		for _, ch := range word {
			wordsX[i] = wordsX[i] | 1<<(uint32)(ch-'a')
		}
	}
	puzzlesX := make([]uint32, len(puzzles))
	for i, puzzle := range puzzles {
		for _, ch := range puzzle {
			puzzlesX[i] = puzzlesX[i] | 1<<(uint32)(ch-'a')
		}
	}
	result := make([]int, len(puzzles))
	for i, puzzle := range puzzles {
		result[i] = 0
		for j, _ := range words {
			// 1. 判断第一个字符
			// 2. 判断字符串的包含关系(异或^检查差异，&检查差异的字符是否属于word)
			if (1<<(uint32)(puzzle[0]-'a'))&wordsX[j] == 0 || wordsX[j]^puzzlesX[i]&wordsX[j] != 0 {
				continue
			}
			result[i]++
		}
	}
	return result
}
```

**复杂度分析：**
- 时间复杂度：O(n * m)，其中 n 为 puzzles 数量，m 为 words 数量
- 空间复杂度：O(n + m)，存储位掩码数组
