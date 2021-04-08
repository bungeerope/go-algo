package main

import "fmt"

func uniqueOccurrences(arr []int) bool {
	countMap := make(map[int]int)

	for _, ele := range arr {
		countMap[ele]++
	}
	times := make(map[int]int)
	for _, c := range countMap {
		times[c]++
	}

	return len(times) == len(countMap)
}

func main() {
	fmt.Println(uniqueOccurrences([]int{1, 2, 2, 1, 1, 3}))
	//fmt.Println(uniqueOccurrences([]int{1, 2}))
	fmt.Println(uniqueOccurrences([]int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}))
}
