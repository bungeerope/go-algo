package main

import "fmt"

func main() {
	s := "LVIII"
	res := romanToInt(s)
	fmt.Println(res)
	fmt.Println(58 == res)
}

var dicts = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romanToInt(s string) int {
	pre, res := "", 0
	for index, ele := range s {
		next := string(ele)
		if index == 0 {
			pre = next
			continue
		}
		if dicts[pre] < dicts[next] {
			res -= dicts[pre]
		} else {
			res += dicts[pre]
		}
		pre = next
	}
	return res + dicts[pre]
}
