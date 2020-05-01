package selector

/*
	选择排序
 */

func SelectSort(arr []int) []int {
	// 遍历整个数组
	for i := 0; i < len(arr)-1; i++ {
		// 记录每一趟中最值的下标
		pos := i
		for j := i + 1; j < len(arr) ; j++ {
			// 遍历剩下未排序数值
			if arr[pos] > arr[j] {
				// 更新下标为未排序中的最值
				pos = j
			}
		}
		// 进行值的换位移动
		arr[i], arr[pos] = arr[pos], arr[i]
	}
	return arr
}
