# 771. 宝石与石头

## 题目信息

| 属性 | 值 |
|------|-----|
| 题号 | [771. 宝石与石头](https://leetcode.cn/problems/jewels-and-stones/) |
| 难度 | 简单 |
| 分类 | hash |
| 标签 | 哈希表、字符串 |
| 文件 | `hash/code_771.go` |

## 题目描述

给你一个字符串 `jewels` 代表石头中宝石的类型，另有一个字符串 `stones` 代表你拥有的石头。`stones` 中每个字符代表了一种你拥有的石头的类型，你想知道你拥有的石头中有多少是宝石。

字母区分大小写，因此 `a` 和 `A` 是不同类型的石头。

**示例 1：**
```
输入：jewels = "aA", stones = "aAAbbbb"
输出：3
```

**示例 2：**
```
输入：jewels = "z", stones = "ZZ"
输出：0
```

**提示：**
- `1 <= jewels.length, stones.length <= 50`
- `jewels` 和 `stones` 仅由英文字母组成
- `jewels` 中的所有字符都是 **唯一的**

## 解题算法

### 方法：哈希表

**核心思路：**
1. 将宝石字符串 `jewels` 中的字符存入哈希表（集合）
2. 遍历石头字符串 `stones`，检查每个字符是否在哈希表中，统计匹配数量

**代码实现：**
```go
func numJewelsInStones(J string, S string) int {
	num := 0
	m := make(map[int32]bool)
	for _, ch := range J {
		m[ch] = true
	}
	for _, ch := range S {
		if m[ch] {
			num++
		}
	}
	return num
}
```

**复杂度分析：**
- 时间复杂度：O(m + n)，其中 m 为 `jewels` 长度，n 为 `stones` 长度
- 空间复杂度：O(m)，哈希表存储宝石类型
