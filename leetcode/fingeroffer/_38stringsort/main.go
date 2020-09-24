package main

//输入一个字符串，打印出该字符串中字符的所有排列。
//
//
//
// 你可以以任意顺序返回这个字符串数组，但里面不能有重复元素。
//
//
//
// 示例:
//
// 输入：s = "abc"
//输出：["abc","acb","bac","bca","cab","cba"]
//
//
//
//
// 限制：
//
// 1 <= s 的长度 <= 8
// Related Topics 回溯算法

func core(remain []byte, last []byte) []string {
	if len(remain) == 1 {
		return []string{string(append(last, remain[0]))}
	}

	var ret []string
	var remainBytes []byte
	m := make(map[byte]bool)

	for i := 0; i < len(remain); i++ {
		if m[remain[i]] {
			continue
		}
		m[remain[i]] = true
		remainBytes = append(remainBytes, remain[:i]...)
		remainBytes = append(remainBytes, remain[i+1:]...)
		ret = append(ret, core(remainBytes, append(last, remain[i]))...)
		remainBytes = remainBytes[0:0]
	}
	return ret
}

func permutation(s string) []string {
	if len(s) == 0 {
		return nil
	}
	return core([]byte(s), []byte{})
}

func main() {
	permutation("abc")
}
