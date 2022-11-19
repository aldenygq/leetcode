package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int{
	n := len(nums)
	if n == 0 {
		return 0
	}
	slow := 1
	for fast := 1; fast < n; fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
func main() {
	a := []int{1,1,2,3,3,4,4,5,6,7,7,8,8,9,9}
	fmt.Printf("len:%v\n",removeDuplicates(a))
}
