package main

import (
	"fmt"
	"strings"
)

func main() {
	stdStr := "PAYPALISHIRING"
	num := 4
	res := convert(stdStr, num)
	fmt.Println(res)
	fmt.Println("PINALSIGYAHRPI" == res)
}

func convert(s string, numRows int) string {

	if len(s) <= 1 || numRows <= 1 {
		return s
	}
	// map 存每一行的元素，补位用""
	rowMap := make(map[int][]string, 0)

	eachSize := 2*numRows - 2
	rowNum := 0
	for index, _ := range s {
		ele := s[index]
		// 一共多少轮
		line := index % eachSize

		//fmt.Println(eachSize, rounds, line, rowNum)

		if line < numRows && rowNum < numRows {
			if line == 0 {
				rowNum = 0
			}
			eleList := rowMap[rowNum]
			if eleList == nil {
				eleList = make([]string, 0)
			}
			eleList = append(eleList, string(ele))
			rowMap[rowNum] = eleList
			rowNum++
		} else {
			rowNum--
			eleList := rowMap[rowNum-1]
			if eleList == nil {
				eleList = make([]string, 0)
			}
			eleList = append(eleList, string(ele))
			rowMap[rowNum-1] = eleList
		}

	}

	res := ""
	for index := 0; index < len(rowMap); index++ {
		fmt.Println(strings.Join(rowMap[index], ""))
		res += strings.Join(rowMap[index], "")
	}
	return res
}
