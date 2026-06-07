# 16. 最接近的三数之和

## 题目信息

| 属性 | 值 |
|------|-----|
| 题号 | [16. 最接近的三数之和](https://leetcode.cn/problems/3sum-closest/) |
| 难度 | Medium |
| 分类 | array |
| 标签 | two pointers, sort |
| 文件 | `array/code_16.go` |

## 题目描述

给你一个长度为 `n` 的整数数组 `nums` 和一个目标值 `target`。请你从 `nums` 中选出三个整数，使它们的和与 `target` 最接近。

返回这三个数的和。

假定每组输入只存在恰好一个解。

**示例 1：**

```
输入：nums = [-1,2,1,-4], target = 1
输出：2
解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2)。
```

**示例 2：**

```
输入：nums = [0,0,0], target = 1
输出：0
```

## 解题算法

### 方法：排序 + 双指针

**核心思路：**
与三数之和类似，先对数组排序，然后固定第一个数，用双指针在剩余部分寻找最接近目标值的三数之和。

维护一个 `closest` 变量记录当前最接近目标值的和。每次计算三数之和后，如果比当前 `closest` 更接近目标值，则更新 `closest`。如果恰好等于目标值，直接返回。

**代码实现：**
```go
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)
	closest := nums[0] + nums[1] + nums[2]

	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, n-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if abs(sum-target) < abs(closest-target) {
				closest = sum
			}
			if sum == target {
				return target
			} else if sum < target {
				left++
			} else {
				right--
			}
		}
	}
	return closest
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```

**复杂度分析：**
- 时间复杂度：O(n^2)
- 空间复杂度：O(1)
