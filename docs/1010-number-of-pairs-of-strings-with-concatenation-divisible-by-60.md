# 1010. 总持续时间可被 60 整除的歌曲

## 题目信息

| 属性 | 值 |
|------|-----|
| 题号 | [1010. 总持续时间可被 60 整除的歌曲](https://leetcode.cn/problems/number-of-pairs-of-strings-with-concatenation-divisible-by-60/) |
| 难度 | Easy |
| 分类 | array |
| 标签 | hash map |
| 文件 | `array/code_1010.go` |

## 题目描述

在歌曲列表中，第 `i` 首歌曲的持续时间为 `time[i]` 秒。

返回其总持续时间（以秒为单位）可被 `60` 整除的歌曲对的数量。形式上，我们希望下标 `i` 和 `j` 满足 `i != j` 且 `time[i] + time[j]` 是 `60` 的倍数。

**示例 1：**

```
输入：time = [30,20,150,100,40]
输出：3
解释：总持续时间可被 60 整除的歌曲对为 (30, 150)、(20, 100) 和 (40, 150)。
```

**示例 2：**

```
输入：time = [60,60,60]
输出：3
```

## 解题算法

### 方法：哈希表（余数配对）

**核心思路：**
利用模运算的性质：如果 `(a + b) % 60 == 0`，则 `(a % 60 + b % 60) % 60 == 0`。因此只需要统计每首歌曲持续时间除以 60 的余数分布。

对于余数 `idx`：
- 如果 `idx == 0` 或 `idx == 30`，它们自身可以配对，配对数为 `C(t, 2) = t * (t-1) / 2`
- 对于其他余数 `idx`，需要找到余数为 `60 - idx` 的歌曲进行配对，配对数为 `t * dir[60-idx]`

为了避免重复计数，只遍历余数 0 到 29。

**代码实现：**
```go
func numPairsDivisibleBy60(time []int) int {
	dir := make([]int, 60)
	for _, t := range time {
		dir[t%60]++
	}
	count := 0
	for idx, t := range dir {
		if t > 0 {
			if idx == 0 || idx == 30 {
				count += (t * (t - 1)) / 2
			} else {
				count += t * dir[60-idx]
			}
		}
		if idx >= 30 {
			break
		}
	}
	return count
}
```

**复杂度分析：**
- 时间复杂度：O(n)
- 空间复杂度：O(1)（固定大小为 60 的数组）
