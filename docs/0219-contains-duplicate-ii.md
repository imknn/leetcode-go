# 219. 存在重复元素 II

## 题目信息

| 属性 | 值 |
|------|-----|
| 题号 | [219. 存在重复元素 II](https://leetcode.cn/problems/contains-duplicate-ii/) |
| 难度 | Easy |
| 分类 | array |
| 标签 | hash set |
| 文件 | `array/code_219.go` |

## 题目描述

给你一个整数数组 `nums` 和一个整数 `k`，判断数组中是否存在两个 **不同的索引** `i` 和 `j`，满足 `nums[i] == nums[j]` 且 `abs(i - j) <= k`。如果存在，返回 `true`；否则，返回 `false`。

**示例 1：**

```
输入：nums = [1,2,3,1], k = 3
输出：true
```

**示例 2：**

```
输入：nums = [1,0,1,1], k = 1
输出：true
```

**示例 3：**

```
输入：nums = [1,2,3,1,2,3], k = 2
输出：false
```

## 解题算法

### 方法：哈希表记录索引

**核心思路：**
使用哈希表将每个元素值映射到其所有出现的索引列表。遍历完成后，对每个出现多次的元素，检查其索引列表中是否存在两个索引之差不超过 `k`。

**代码实现：**
```go
func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int][]int)
	for idx, num := range nums {
		if _, ok := m[num]; ok {
			m[num] = append(m[num], idx)
		} else {
			n := make([]int, 1)
			n[0] = idx
			m[num] = n
		}
	}
	// 检查重复元素
	for _, v := range m {
		length := len(v)
		if length >= 2 {
			for i := 0; i < length; i++ {
				for j := i + 1; j < length; j++ {
					if v[j]-v[i] <= k {
						return true
					}
				}
			}
		}
	}
	return false
}
```

**复杂度分析：**
- 时间复杂度：O(n^2)（最坏情况下所有元素相同）
- 空间复杂度：O(n)
