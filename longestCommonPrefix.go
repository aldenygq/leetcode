package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	count := len(strs)
	fmt.Printf("prefix:%v\n",prefix)
	fmt.Printf("count:%v\n",count)
	fmt.Printf("==============start===============\n")
	for i := 1; i < count; i++ {
		prefix = lcp(prefix, strs[i])
		fmt.Printf("prefix:%v\n",prefix)
		fmt.Printf("index:%v\n",i)
		if len(prefix) == 0 {
			fmt.Printf("len(prefix):%v\n",len(prefix))
			break
		}
	}
	return prefix
}

func lcp(str1, str2 string) string {
	length := min(len(str1), len(str2))
	index := 0
	for index < length && str1[index] == str2[index] {
		index++
	}
	return str1[:index]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	s := []string{"abcef","abcfg","abcdf","abgh"}
	fmt.Printf("str:%v\n",longestCommonPrefix(s))
}
