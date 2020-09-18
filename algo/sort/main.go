package main

import (
	"algo/algo/sort/compare/insert"
	"algo/algo/sort/compare/selector"
	"fmt"
)

func main() {
	arr := []int{1, 3, 5, 6, 3, 7, 4, 9, 2, 8}
	// 选择排序
	fmt.Println("======选择排序======")
	selectRes := selector.SelectSort(arr)
	for x := range selectRes {
		fmt.Printf("%d", arr[x])
	}
	fmt.Println()
	// 插入排序
	fmt.Println("======插入排序======")
	insertRes := insert.InsertSort(arr)
	for x := range insertRes {
		fmt.Printf("%d", arr[x])
	}
	fmt.Println()
	// 希尔排序
	fmt.Println("======希尔排序======")
	shellInsertRes := insert.ShellInsertSort(arr)
	for x := range shellInsertRes {
		fmt.Printf("%d", arr[x])
	}
}
