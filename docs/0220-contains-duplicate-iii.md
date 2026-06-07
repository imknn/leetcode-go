# 220. 存在重复元素 III

## 题目信息

| 属性 | 值 |
|------|-----|
| 题号 | [220. 存在重复元素 III](https://leetcode.cn/problems/contains-duplicate-iii/) |
| 难度 | 困难 |
| 分类 | hash |
| 标签 | 有序集合、数组、哈希表、滑动窗口 |
| 文件 | `hash/code_220.go` |

## 题目描述

给你一个整数数组 `nums` 和两个整数 `k` 和 `t`。请你判断是否存在两个不同下标 `i` 和 `j`，使得 `abs(nums[i] - nums[j]) <= t`，同时又满足 `abs(i - j) <= k`。

如果存在则返回 `true`，不存在返回 `false`。

**示例 1：**
```
输入：nums = [1,2,3,1], k = 3, t = 0
输出：true
```

**示例 2：**
```
输入：nums = [1,0,1,1], k = 1, t = 2
输出：true
```

**示例 3：**
```
输入：nums = [1,5,9,1,5,9], k = 2, t = 3
输出：false
```

**提示：**
- `0 <= nums.length <= 2 * 10^4`
- `-2^31 <= nums[i] <= 2^31 - 1`
- `0 <= k <= 10^4`
- `0 <= t <= 2^31 - 1`

## 解题算法

### 方法：暴力枚举

**核心思路：**
双重循环遍历所有满足距离条件 `j - i <= k` 的下标对 `(i, j)`，检查它们的值之差是否满足 `abs(nums[i] - nums[j]) <= t`。

**代码实现：**
```go
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	length := len(nums)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if j-i <= k {
				abs := nums[j] - nums[i]
				if abs < 0 {
					abs = -abs
				}
				if abs <= t {
					return true
				}
			}
		}
	}
	return false
}
```

**复杂度分析：**
- 时间复杂度：O(n^2)，双重循环遍历数组
- 空间复杂度：O(1)，只使用常量额外空间
