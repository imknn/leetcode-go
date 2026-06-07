# 4. 寻找两个正序数组的中位数

## 题目信息

| 属性 | 值 |
|------|-----|
| 题号 | [4. 寻找两个正序数组的中位数](https://leetcode.cn/problems/median-of-two-sorted-arrays/) |
| 难度 | 困难 |
| 分类 | binary_search |
| 标签 | 数组、二分查找、分治 |
| 文件 | `binary_search/code_4.go` |

## 题目描述

给定两个大小分别为 m 和 n 的正序（从小到大）数组 `nums1` 和 `nums2`。请你找出并返回这两个正序数组的 **中位数**。

算法的时间复杂度应该为 O(log (m+n))。

**示例 1：**
```
输入：nums1 = [1,3], nums2 = [2]
输出：2.00000
解释：合并数组 = [1,2,3]，中位数 2
```

**示例 2：**
```
输入：nums1 = [1,2], nums2 = [3,4]
输出：2.50000
解释：合并数组 = [1,2,3,4]，中位数 (2 + 3) / 2 = 2.5
```

**提示：**
- `nums1.length == m`
- `nums2.length == n`
- `0 <= m <= 1000`
- `0 <= n <= 1000`
- `1 <= m + n <= 2000`
- `-10^6 <= nums1[i], nums2[i] <= 10^6`

## 解题算法

### 方法：二分查找

**核心思路：**
中位数将合并后的数组分成等长的左右两部分。我们在较短数组上二分查找分割点 `partition1`，使得：
- `maxLeft1 <= minRight2` 且 `maxLeft2 <= minRight1`

**算法步骤：**
1. 确保 `nums1` 是较短的数组（交换如果需要）
2. 在 `nums1` 上二分查找 `partition1`（左半部分个数）
3. 计算 `partition2 = (m+n+1)/2 - partition1`
4. 检查分割是否合法：左半部分最大值 <= 右半部分最小值
5. 根据总长度奇偶计算中位数

**示例推导（nums1=[1,3], nums2=[2]）：**
```
m=2, n=1, 总长度=3（奇数）
partition2 = (3+1)/2 - partition1 = 2 - partition1

partition1=1, partition2=1:
  maxLeft1=1, minRight1=3
  maxLeft2=2, minRight2=∞

验证: 1<=∞ ✓, 2<=3 ✓ → 合法！
中位数 = max(1,2) = 2
```

**代码实现：**
```go
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    if len(nums1) > len(nums2) {
        nums1, nums2 = nums2, nums1
    }
    m, n := len(nums1), len(nums2)
    left, right := 0, m

    for left <= right {
        partition1 := (left + right) / 2
        partition2 := (m+n+1)/2 - partition1

        maxLeft1 := math.MinInt64
        if partition1 > 0 { maxLeft1 = nums1[partition1-1] }
        minRight1 := math.MaxInt64
        if partition1 < m { minRight1 = nums1[partition1] }
        maxLeft2 := math.MinInt64
        if partition2 > 0 { maxLeft2 = nums2[partition2-1] }
        minRight2 := math.MaxInt64
        if partition2 < n { minRight2 = nums2[partition2] }

        if maxLeft1 <= minRight2 && maxLeft2 <= minRight1 {
            if (m+n)%2 == 1 {
                return float64(max(maxLeft1, maxLeft2))
            }
            return float64(max(maxLeft1, maxLeft2)+min(minRight1, minRight2)) / 2.0
        } else if maxLeft1 > minRight2 {
            right = partition1 - 1
        } else {
            left = partition1 + 1
        }
    }
    return 0
}
```

**复杂度分析：**
- 时间复杂度：O(log(min(m,n)))，在较短数组上二分
- 空间复杂度：O(1)，只使用常数额外空间
