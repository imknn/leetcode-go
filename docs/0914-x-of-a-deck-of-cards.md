# 914. 卡牌分组

## 题目信息

| 属性 | 值 |
|------|-----|
| 题号 | [914. 卡牌分组](https://leetcode.cn/problems/x-of-a-deck-of-cards/) |
| 难度 | Easy |
| 分类 | math |
| 标签 | 数学、哈希表 |
| 文件 | `math/code_914.go` |

## 题目描述

给定一副每张都有整数的牌，每张牌的正面都标有一个数字。你需要将这些牌分成若干组，每组的大小都是 X，且每组中的数字都相同。

返回 true 当且仅当可以这样分组。

**示例 1：**
```
输入：deck = [1,2,3,4,4,3,2,1]
输出：true
解释：可以分成 [1,1]、[2,2]、[3,3]、[4,4] 四组，每组大小为 2。
```

**示例 2：**
```
输入：deck = [1,1,1,2,2,2,3,3]
输出：false
解释：没有一种分组大小满足条件。
```

## 解题算法

### 方法：最大公约数

**核心思路：**
统计每种卡牌出现的次数，然后求所有频次的最大公约数（GCD）。如果最大公约数大于等于 2，说明可以分组；否则无法分组。使用辗转相除法依次求所有频次的 GCD，如果中途 GCD 变为 1 则提前返回 false。

**代码实现：**
```go
func hasGroupsSizeX(deck []int) bool {
	m := make(map[int]int)
	for _, v := range deck {
		if _, ok := m[v]; ok {
			m[v] = m[v] + 1
		} else {
			m[v] = 1
		}
	}
	a := 0
	for _, v := range m {
		if a == 0 {
			a = v
			continue
		}
		a = common.Gcd(a, v)
		if a < 2 {
			return false
		}
	}
	return a > 1
}
```

**复杂度分析：**
- 时间复杂度：O(n log C)，n 为牌的数量，C 为最大频次，GCD 计算时间与 log C 成正比
- 空间复杂度：O(n)，哈希表存储每种卡牌的频次
