package main

import "fmt"

func main() {
	num := 312
	res := intToRoman(num)
	fmt.Println(res)
	fmt.Println("CCCXII" == res)
}

var (
	thousands = []string{"", "M", "MM", "MMM"}
	hundres   = []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	ten       = []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	one       = []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
)

func intToRoman(num int) string {
	if num < 1 || num > 3999 {
		return ""
	}
	return thousands[num/1000] + hundres[num%1000/100] + ten[num%100/10] + one[num%10]
}
