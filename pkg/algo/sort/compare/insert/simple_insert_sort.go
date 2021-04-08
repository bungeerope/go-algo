package insert

/*
	插入排序
	思路：
		默认数列中第一个元素是已经完成排序(即下标从1开始循环)
		获取下一个元素，与已经完成排序的部分依次进行比较
		当找到合适位置将元素进行插入
		依次循环列表中未排序元素直至结束
*/
func InsertSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		curr := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > curr {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = curr
	}
	return arr
}
