package main

import (
	"fmt"
	"sort"
)

//leetcode submit region begin(Prohibit modification and deletion)
func hIndex(citations []int) int {
	length := len(citations)
	first := 0
	last := length - 1
	for first <= last {
		// 找出当前查询的中位数索引
		midIndex := first + (last-first)/2
		if citations[midIndex] >= length-midIndex {
			// 末尾游标索引左移
			last = midIndex - 1
		} else {
			// 开始游标索引右移
			first = midIndex + 1
		}
	}
	return length - first
}

func hIndex2(citations []int) int {
	n := len(citations)
	return n - sort.Search(n, func(x int) bool { return citations[x] >= n-x })
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	conditions := []int{0, 1, 3, 5, 6}
	res := hIndex(conditions)
	res2 := hIndex2(conditions)
	fmt.Println(res)
	fmt.Println(res2)
}
