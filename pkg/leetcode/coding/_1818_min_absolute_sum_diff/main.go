package main

import (
	"fmt"
)

//leetcode submit region begin(Prohibit modification and deletion)
func minAbsoluteSumDiff(nums1 []int, nums2 []int) int {
	if len(nums1) == len(nums1) {
		max := 0
		maxIndex := 0
		sum := 0
		for i, v := range nums1 {
			x := nums2[i]
			if nums2[i] < 0 {
				x = CalcAbs(x)
			}
			n := CalcAbs(v - x)
			if i == 0 {
				max = n
			} else {
				if n != 0 && max < n {
					sum = sum + max
					max = n
					maxIndex = i
				} else {
					sum += n
				}
			}
		}
		if max > 0 {
			min := 0
			for i, v := range nums1 {
				x := nums2[maxIndex]
				if nums2[maxIndex] < 0 {
					x = CalcAbs(nums2[maxIndex])
				}
				n := CalcAbs(v - x)
				if i == 0 {
					min = n
				} else {
					if min > n && n < max {
						min = n
					}
				}
			}
			return (sum + min) % (10e9 + 7)
		}
	}
	return 0
}

func CalcAbs(a int) (ret int) {
	ret = (a ^ a>>31) - a>>31
	return
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	num1_1 := []int{1, 7, 5}
	num2_1 := []int{2, 3, 5}
	fmt.Println(minAbsoluteSumDiff(num1_1, num2_1))

	num1_2 := []int{2, 4, 6, 8, 10}
	num2_2 := []int{2, 4, 6, 8, 10}
	fmt.Println(minAbsoluteSumDiff(num1_2, num2_2))

	num1_3 := []int{1, 10, 4, 4, 2, 7}
	num2_3 := []int{9, 3, 5, 1, 7, 4}
	fmt.Println(minAbsoluteSumDiff(num1_3, num2_3))

	num1_4 := []int{1, 28, 21}
	num2_4 := []int{9, 21, 20}
	fmt.Println(minAbsoluteSumDiff(num1_4, num2_4))
}
