package main

import (
	"fmt"
	"sort"
)

func main() {
	slc := []int{21,9,6,4,3,7,9,21,5,21}
	if len(slc) <= 0 {
		fmt.Printf("slc is null")
	}
	max := slc[0]
	index := 0
	var indexs []int = make([]int,0)
	for i:=1;i < len(slc);i++ {
		if slc[i] < max {
			continue
		}else if slc[i] > max {
			max = slc[i]
			index = i
			indexs = append(indexs,index)
		}else if slc[i] == max {
			//max = slc[i]
			index = i
			indexs = append(indexs,index)
		}
	}
	if slc[index] == max {
		indexs = append(indexs,0)
		sort.Ints(indexs)
	}
	
	fmt.Printf("indexs:%v\n",indexs)
}
