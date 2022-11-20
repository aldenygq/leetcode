package main

import (
	"fmt"
)

func removeElement(nums []int, val int) int {
	left := 0
	for _, v := range nums { // v 即 nums[right]
		if v != val {
			nums[left] = v
			left++
		}
	}
	return left
}

func main() {
	var a []int = []int{1,3,4,4,5,4,6,7,4,5,6,2}
	fmt.Printf("result:%v\n",removeElement(a,3))
}