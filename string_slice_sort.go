//字符串数组排序，字符串包含字母和数字
//输入：["f54","af3","a123", "b235", "abb567", "abc12345","b14", "c123", "dc234", "cb235"],输出：["a123", "abb567", "abc12345", "af3", "b14", "b235", "c123", "cb235", "dc234", "f54"]


package main

import (
	"fmt"
	"sort"
	"strconv"
)

// 分解字符串为字母部分和数字部分
func splitParts(s string) (letters string, num int) {
	var i int
	// 找到第一个数字的位置
	for i = 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			break
		}
	}
	letters = s[:i]
	numStr := s[i:]
	num, _ = strconv.Atoi(numStr) // 忽略错误处理（假设输入合法）
	return
}

func main() {
	// 输入数据
	input := []string{
		"f54","af3","a123", "b235", "abb567", "abc12345",
		"b14", "c123", "dc234", "cb235",
	}

	// 自定义排序规则
	sort.Slice(input, func(i, j int) bool {
		lettersI, numI := splitParts(input[i])
		lettersJ, numJ := splitParts(input[j])

		// 优先比较字母部分，再比较数字部分
		if lettersI < lettersJ {
			return true
		} else if lettersI == lettersJ {
			return numI < numJ
		}
		return false
	})

	// 输出结果
	fmt.Printf("排序结果: %#v\n", input)
}
