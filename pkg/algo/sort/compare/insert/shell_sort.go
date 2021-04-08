package insert

/*
	希尔排序
	思路：
		在插入排序的算法上进行优化，将切分成N份数列按照插入算法思想进行比较(当数据量大时，插入排序比较次数增多，通过切分成小份进行优化)

*/

func ShellInsertSort(arr []int) []int {
	// i 为切分粒度，二分法思维
	for gap := len(arr) / 2; gap > 0; gap /= 2 {
		for i := gap; i < len(arr); i++ {
			curr := arr[i]
			j := i - gap
			for j > 0 && arr[j] < curr {
				arr[j+gap] = arr[j]
				j -= gap
			}
			arr[j+gap] = curr
		}
	}
	return arr
}
