package main

import (
	"fmt"
)

//给你一个数组 nums，对于其中每个元素 nums[i]，请你统计数组中比它小的所有数字的数目。
//
// 换而言之，对于每个 nums[i] 你必须计算出有效的 j 的数量，其中 j 满足 j != i 且 nums[j] < nums[i] 。
//
// 以数组形式返回答案。
//
//
//
// 示例 1：
//
// 输入：nums = [8,1,2,2,3]
//输出：[4,0,1,1,3]
//解释：
//对于 nums[0]=8 存在四个比它小的数字：（1，2，2 和 3）。
//对于 nums[1]=1 不存在比它小的数字。
//对于 nums[2]=2 存在一个比它小的数字：（1）。
//对于 nums[3]=2 存在一个比它小的数字：（1）。
//对于 nums[4]=3 存在三个比它小的数字：（1，2 和 2）。
//
//
// 示例 2：
//
// 输入：nums = [6,5,4,8]
//输出：[2,1,0,3]
//
//
// 示例 3：
//
// 输入：nums = [7,7,7,7]
//输出：[0,0,0,0]
//
//
//
//
// 提示：
//
//
// 2 <= nums.length <= 500
// 0 <= nums[i] <= 100
//
// Related Topics 数组 哈希表
// 👍 96 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func smallerNumbersThanCurrent1(nums []int) []int {
	res := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[j] < nums[i] {
				res[i]++
			}
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

//leetcode submit region begin(Prohibit modification and deletion)
func smallerNumbersThanCurrent(nums []int) []int {
	res := make([]int, len(nums))

	numsRepl := make([]int, len(nums))
	copy(numsRepl, nums)

	sortSmaller := make(map[int]int, len(nums))

	for i := 0; i < len(numsRepl); i++ {
		for j := 0; j < i; j++ {
			if numsRepl[i] == numsRepl[j] {
				continue
			}
			if numsRepl[i] < numsRepl[j] {
				swap := numsRepl[j]
				numsRepl[j] = numsRepl[i]
				numsRepl[i] = swap
			}
		}
	}
	fmt.Println("sort:", numsRepl)

	for _x := range numsRepl {
		if (_x == 0 || sortSmaller[numsRepl[_x]] == 0) && (_x != 0 && numsRepl[_x] != numsRepl[_x-1]) {
			sortSmaller[numsRepl[_x]] = _x
		}
	}

	for x, y := range nums {
		res[x] = sortSmaller[y]
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	a := []int{8, 1, 2, 2, 3}
	fmt.Println(smallerNumbersThanCurrent1(a))
	b := []int{6, 5, 4, 8}
	fmt.Println(smallerNumbersThanCurrent1(b))
	c := []int{7, 7, 7, 7}
	fmt.Println(smallerNumbersThanCurrent1(c))
	d := []int{10, 2, 4, 5, 8}
	fmt.Println(smallerNumbersThanCurrent1(d))
	e := []int{6, 2, 8, 0, 3, 8}
	fmt.Println(smallerNumbersThanCurrent1(e))
}
