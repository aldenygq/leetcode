/*
题目：
给你一个偶数长度的字符串 s 。将其拆分成长度相同的两半，前一半为 a ，后一半为 b 。
两个字符串 相似 的前提是它们都含有相同数目的元音（'a'，'e'，'i'，'o'，'u'，'A'，'E'，'I'，'O'，'U'）。注意，s 可能同时含有大写和小写字母。
如果 a 和 b 相似，返回 true ；否则，返回 false 。

示例 1：
输入：s = "book"
输出：true
解释：a = "bo" 且 b = "ok" 。a 中有 1 个元音，b 也有 1 个元音。所以，a 和 b 相似。

示例 2：
输入：s = "textbook"
输出：false
解释：a = "text" 且 b = "book" 。a 中有 1 个元音，b 中有 2 个元音。因此，a 和 b 不相似。

注意，元音 o 在 b 中出现两次，记为 2 个。

提示：
2 <= s.length <= 1000
s.length 是偶数
s 由 大写和小写 字母组成
 */
package main

import (
	"fmt"
	"strings"
	"unicode"
)
const (
	VOWEL = "aeiouAEIOU"
)
/*
题解一：计数
思路与算法
题目给定一个偶数长度的字符串 ss，并给出字符串「相似」的定义：若两个字符串中含有相同数目的元音字母，则这两个字符串「相似」。现在我们将给定字符串 ss 拆分成长度相同的两半，前一半表示为字符串 aa，后一半为字符串 bb，我们需要判断字符串 aa 和 bb 是否「相似」，那么我们只需要按照「相似」的定义统计字符串 aa 和 bb 中的元音字母的个数是否相等即可。
 */
func halvesAreAlike(s string) bool {
	cnt := 0
	for _, c := range s[:len(s)/2] {
		if strings.ContainsRune(VOWEL, c) {
			cnt++
		}
	}
	for _, c := range s[len(s)/2:] {
		if strings.ContainsRune(VOWEL, c) {
			cnt--
		}
	}
	return cnt == 0
}
/*
复杂度分析
时间复杂度：O(n)，其中 n 为字符串 s的长度。
空间复杂度：O(n)，其中 n 为字符串 s 的长度，主要为表示字符串 a, b 的空间开销，也可以通过双指针在原字符串中遍历 a, b进行计数操作来实现O(1) 的空间开销。
 */

/*
题解二：位运算
解题思路
~只需要把字符串ss分成前后两部分分别统计即可。
~由于字符串只有大小写字母，统计元音字母个数，可以用int表示26个字母出现的位置。
~统计过程中需要把大小字母转换为统一的形式。
*/
func halvesAreAlike2(s string) bool {
	n := len(s)
	vowel := []byte{'a', 'e', 'i', 'o', 'u'}
	mask := 0
	for _, c := range vowel {
		mask |= 1 << (c -'a')
	}
	count1 := 0
	count2 := 0
	for i := 0; i < n / 2; i++ {
		if ((1 << (unicode.ToLower(rune(s[i])) - 'a')) & mask) != 0 {
			count1++
		}
		if ((1 << (unicode.ToLower(rune(s[n-1-i])) - 'a')) & mask) != 0 {
			count2++
		}
	}
	return count1 == count2
}
/*
复杂度分析
时间复杂度: O(n)，n为字符串的长度。
空间复杂度: O(1)。
 */

func main() {
	a := "leetcode"
	b := "good"
	c := "texkbook"
	fmt.Printf("r1:%v,r2:%v,r3:%v\n",halvesAreAlike(a),halvesAreAlike(b),halvesAreAlike(c))
}