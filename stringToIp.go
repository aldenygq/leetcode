package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "25525522135"
	slc := restoreIpAddresses(str)
	fmt.Printf("ip:%v\n",slc)
}
func restoreIpAddresses(s string) []string {
	n := len(s)
	if n < 4 {
		return nil
	}
	set :=[]string{}
	ans := []string{}
	var back func(int, int)
	back = func(start int, next int) {
		if start == n && next == 4 {
			ans = append(ans, strings.Join(set, "." ))
			return
		}
		if start >= n || next >= 4 {
			return
		}
		
		if n - start < 4 - next -1 {
			return
		}
		num := 0
		for i:=0;i<3 && start+i<n;i++ {
			if i == 0 && s[start+i] == '0' {
				set = append(set,"0")
				back(start+1,next+1)
				set = set[:next]
				return
			}
			num = num * 10 +  int(s[start+i] - '0')
			
			if num > 255 {
				return
			}
			if start+i+1 >= n {
				set =  append(set,s[start:])
			}else{
				set =  append(set,s[start:start+i+1])
			}
			back(start+i+1,next+1)
			set = set[:next]
		}
	}
	back(0,0)
	return ans
}