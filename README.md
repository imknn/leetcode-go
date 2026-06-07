# leetcode-go

LeetCode 刷题记录，按算法分类整理。

## 分类目录

| 分类 | 目录 | 题数 |
|------|------|------|
| 数组 | [array/](array/) | 16 |
| 字符串 | [string/](string/) | 21 |
| 链表 | [linked_list/](linked_list/) | 5 |
| 栈/队列 | [stack_queue/](stack_queue/) | 2 |
| 二分查找 | [binary_search/](binary_search/) | 2 |
| 动态规划 | [dynamic_programming/](dynamic_programming/) | 3 |
| 位运算 | [bit_manipulation/](bit_manipulation/) | 13 |
| 数学 | [math/](math/) | 17 |
| 哈希表 | [hash/](hash/) | 4 |
| 回溯 | [backtracking/](backtracking/) | 3 |
| 排序 | [sorting/](sorting/) | 1 |
| 随机化 | [random/](random/) | 1 |
| 其他 | [other/](other/) | 2 |
| Shell | [shell/](shell/) | 3 |
| **合计** | | **93** |

## 公共工具

[common/](common/) - 提供 `Gcd`、`Reverse` 等公共函数。

## 运行测试

```bash
# 运行所有测试
go test ./...

# 运行某个分类的测试
go test ./array/

# 运行单个测试
go test ./array/ -run Test_twoSum
```
