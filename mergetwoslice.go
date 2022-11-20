package main

import (
	"fmt"
)

func MergeTwoSlice(a,b []int)  []int {
	if len(a) <=0 || len(b) <= 0 {
		return nil
	}
	
	lenA :=  len(a)
	lenB :=  len(a)
	ab := make([]int, lenA+lenB)
	index := 0
	i, j := 0, 0
	
	for i, j = 0, 0; i < lenA && j < lenB; {
		if a[i] < b[j] {
			ab[index] = a[i]
			index++
			i++
		} else if a[i] > b[j] {
			ab[index] = b[j]
			index++
			j++
		} else {
			ab[index] = a[i]
			index++
			i++
		}
	}
	//一定有一个数组是先遍历完，把另外一个数组的所有元素都添加进去
	for i < lenA {
		ab[index] = a[i]
		i++
		index++
	}
	for j < lenB {
		ab[index] = b[j]
		j++
		index++
	}
	return ab
}

func main() {
	a := []int{1,2,3,4,5,6,7,8}
	b := []int{1,3,4,7,8,10,13,24}
	
	ab := MergeTwoSlice(a,b)
	fmt.Printf("ab:%v\n",ab)
}
