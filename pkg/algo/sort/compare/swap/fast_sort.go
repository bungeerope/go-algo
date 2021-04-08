package swap

import (
	"math/rand"
	"time"
)

/*
	快速排序
*/

func FastSort(arr []int, left, right int) []int {
	if left >= right {
		return arr
	}
	i := left
	j := right
	rand.Seed(time.Now().Unix())
	r := rand.Intn(right-left) + left
	arr[i], arr[r] = arr[r], arr[i]

	curr := arr[i]
	for i < j {
		for i < j && arr[j] >= curr {
			j--
		}
		arr[i] = arr[j]
		for i < j && arr[i] <= curr {
			i++
		}
		arr[j] = arr[i]
	}
	arr[i] = curr
	arr = FastSort(arr, left, i-1)
	arr = FastSort(arr, i+1, right)
	return arr
}
