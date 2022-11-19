package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	// 哈希集合，记录每个字符是否出现过
	m := map[byte]int{}
	n := len(s)
	// 右指针，初始值为 -1，相当于我们在字符串的左边界的左侧，还没有开始移动
	rk, ans := -1, 0
	for i := 0; i < n; i++ {
		fmt.Printf("m start:%v\n",m)
		fmt.Printf("index start:%v\n",i)
		fmt.Printf("rk start:%v\n",rk)
		fmt.Printf("ans start:%v\n",ans)
		if i != 0 {
			// 左指针向右移动一格，移除一个字符
			delete(m, s[i-1])
		}
		for rk + 1 < n && m[s[rk+1]] == 0 {
			// 不断地移动右指针
			m[s[rk+1]]++
			rk++
		}
		// 第 i 到 rk 个字符是一个极长的无重复字符子串
		ans = max(ans, rk - i + 1)
		fmt.Printf("m end:%v\n",m)
		fmt.Printf("index end:%v\n",i)
		fmt.Printf("rk end:%v\n",rk)
		fmt.Printf("ans end:%v\n",ans)
		fmt.Printf("====================")
	}
	return ans
}

func max(x, y int) int {
	fmt.Printf("x:%v,y:%v\n",x,y)
	if x < y {
		return y
	}
	return x
}
func main() {
	s := "abcabcbb"
	fmt.Printf("len:%v\n",lengthOfLongestSubstring(s))
}
