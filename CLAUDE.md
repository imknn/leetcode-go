# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## 项目概述

LeetCode 刷题记录仓库，使用 Go 语言（模块名 `izqy.top`，Go 1.16）。少量 Shell 题目。

## 目录结构

项目按算法分类组织：

- `array/` — 数组
- `string/` — 字符串
- `linked_list/` — 链表
- `stack_queue/` — 栈/队列
- `binary_search/` — 二分查找
- `dynamic_programming/` — 动态规划
- `bit_manipulation/` — 位运算
- `math/` — 数学
- `hash/` — 哈希表
- `backtracking/` — 回溯
- `sorting/` — 排序
- `random/` — 随机化
- `other/` — 其他
- `shell/` — Shell 题目
- `common/` — 公共工具函数

## 命名与代码规范

### 题目文件命名
- Go 题目：`code_<题号>.go` + `code_<题号>_test.go`（如 `code_1.go`、`code_1_test.go`）
- Shell 题目：`code_<题号>.sh`（仅有解题脚本，无测试）
- 特殊题号用 `code_lcp_<题号>.go`（力扣杯题目）或 `code_offer_<题号>.go`（剑指 Offer）

### 包名
每个目录的包名即目录名（如 `package array`、`package linked_list`），遵循 Go 的命名约定。

### 测试格式
使用表驱动测试 + 子测试模式：
```go
tests := []struct {
    name string
    args args
    want <type>
}{...}
for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) { ... })
}
```

## 常用命令

```bash
# 运行所有测试
go test ./...

# 运行某个分类的测试
go test ./array/

# 运行单个测试函数
go test ./array/ -run Test_twoSum

# 运行公共工具测试
go test ./common/
```

## 新增题目流程

1. 根据题目类型，在对应分类目录下创建 `code_<题号>.go`（包名与目录一致）
2. 创建 `code_<题号>_test.go`，使用表驱动测试格式
3. 更新该分类目录 `README.md` 中的题目表格
4. 更新根目录 `README.md` 中的题数统计
