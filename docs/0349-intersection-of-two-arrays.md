# 349. 两个数组的交集

## 题目信息

| 属性 | 值 |
|------|-----|
| 题号 | [349. 两个数组的交集](https://leetcode.cn/problems/intersection-of-two-arrays/) |
| 难度 | Easy |
| 分类 | array |
| 标签 | hash set |
| 文件 | `array/code_349.go` |

## 题目描述

给定两个整数数组 `nums1` 和 `nums2`，返回它们的交集。输出结果中的每个元素一定是 **唯一** 的。我们可以不考虑输出结果的顺序。

**示例 1：**

```
输入：nums1 = [1,2,2,1], nums2 = [2,2]
输出：[2]
```

**示例 2：**

```
输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
输出：[9,4]
```

## 解题算法

### 方法：哈希集合

**核心思路：**
使用两个哈希集合。先将 `nums1` 的所有元素加入第一个集合。然后遍历 `nums2`，如果元素在第一个集合中存在，则加入第二个集合（自动去重）。最后将第二个集合的元素转换为结果数组。

**代码实现：**
```go
func intersection(nums1 []int, nums2 []int) []int {
	m := make(map[int]bool)
	for _, i := range nums1 {
		m[i] = true
	}
	n := make(map[int]bool)
	for _, i := range nums2 {
		if _, ok := m[i]; ok {
			n[i] = true
		}
	}
	result, idx := make([]int, len(n)), 0
	for k, _ := range n {
		result[idx] = k
		idx++
	}
	return result
}
```

**复杂度分析：**
- 时间复杂度：O(m + n)
- 空间复杂度：O(m + n)
