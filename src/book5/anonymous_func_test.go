package book5

import (
	"sort"
	"testing"
)

// 记录每个课程的前置课程
var prereqs = map[string][]string{
	"算法课": {"数据结构"},
	"微积分": {"线性代数"},
	"编译原理": {
		"数据结构",
		"形式语言",
		"计算机组成原理",
	},
	"数据结构": {"离散数学"},
	"数据库":  {"数据结构"},
	"离散数学": {"编程入门"},
	"形式语言": {"离散数学"},
	"网络":   {"操作系统"},
	"操作系统": {"数据结构", "计算机组成原理"},
	"程序语言": {"数据结构", "计算机组成原理"},
}

// 深度优先遍历
func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)

	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}
	var keys []string
	for key := range m{
		keys = append(keys,key)
	}
	sort.Strings(keys)
	visitAll(keys)
	return order
}

func TestAnonymousFunc(t *testing.T) {

}
