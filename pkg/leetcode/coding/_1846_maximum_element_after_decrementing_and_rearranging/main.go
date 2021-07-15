package main

import (
	"fmt"
	"sort"
)

//leetcode submit region begin(Prohibit modification and deletion)
func maximumElementAfterDecrementingAndRearranging(arr []int) int {
	// 排序数组，从小到大
	sort.Ints(arr)
	l := 0
	for index, item := range arr {
		l++
		// 保证第一个数是1
		if index == 0 && item != 1 {
			arr[index] = 1
		}
		// 如果right比left大于1，进行替换
		if index > 0 && item-arr[index-1] > 1 {
			arr[index] = arr[index-1] + 1
		}
	}
	// 最终末尾即为结果
	return arr[l-1]
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	arr := []int{2, 2, 1, 2, 1}
	fmt.Println(maximumElementAfterDecrementingAndRearranging(arr))
	arr2 := []int{100, 1, 1000}
	fmt.Println(maximumElementAfterDecrementingAndRearranging(arr2))
	arr3 := []int{1, 2, 3, 4, 5}
	fmt.Println(maximumElementAfterDecrementingAndRearranging(arr3))
}
