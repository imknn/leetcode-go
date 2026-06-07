# 217. 存在重复元素

## 题目信息

| 属性 | 值 |
|------|-----|
| 题号 | [217. 存在重复元素](https://leetcode.cn/problems/contains-duplicate/) |
| 难度 | Easy |
| 分类 | array |
| 标签 | hash set |
| 文件 | `array/code_217.go` |

## 题目描述

给你一个整数数组 `nums`。如果任一数值在数组中出现 **至少两次**，返回 `true`；如果数组中每个元素都是唯一的，返回 `false`。

**示例 1：**

```
输入：nums = [1,2,3,1]
输出：true
```

**示例 2：**

```
输入：nums = [1,2,3,4]
输出：false
```

**示例 3：**

```
输入：nums = [1,1,1,3,3,4,3,2,4,2]
输出：true
```

## 解题算法

### 方法：哈希集合

**核心思路：**
使用哈希集合记录已遍历过的元素。遍历数组时，检查当前元素是否已在集合中出现过，如果出现过则说明存在重复元素，返回 `true`。否则将当前元素加入集合。遍历结束后仍未找到重复元素，返回 `false`。

**代码实现：**
```go
func containsDuplicate(nums []int) bool {
	m := make(map[int]bool)
	for _, num := range nums {
		if _, ok := m[num]; ok {
			return true
		} else {
			m[num] = true
		}
	}
	return false
}
```

**复杂度分析：**
- 时间复杂度：O(n)
- 空间复杂度：O(n)
