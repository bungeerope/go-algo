package main

import (
	"fmt"
	"sort"
)

//给定两个数组，编写一个函数来计算它们的交集。
//
//
//
// 示例 1：
//
// 输入：nums1 = [1,2,2,1], nums2 = [2,2]
//输出：[2]
//
//
// 示例 2：
//
// 输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
//输出：[9,4]
//
//
//
// 说明：
//
//
// 输出结果中的每个元素一定是唯一的。
// 我们可以不考虑输出结果的顺序。
//
// Related Topics 排序 哈希表 双指针 二分查找
// 👍 288 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func intersection(nums1 []int, nums2 []int) []int {

	// 先排序，让两个数组有序
	sort.Ints(nums1)
	sort.Ints(nums2)

	var x []int
	// 通过两个指针进行比较,两个指针有一个越界就退出循环
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		m, n := nums1[i], nums2[j]

		if m == n {

			if len(x) == 0 || m > x[len(x)-1] {
				x = append(x, m)
			}
			// 两个相等，指针同时右移
			i++
			j++
		} else if m < n {
			i++
		} else {
			j++
		}
	}

	return x
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums11 := []int{4, 9, 5}
	nums2 := []int{2, 2}
	nums22 := []int{9, 4, 9, 8, 4}

	fmt.Println(intersection(nums1, nums2))
	fmt.Println(intersection(nums11, nums22))
}
