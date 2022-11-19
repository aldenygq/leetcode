package main

import (
	"fmt"
)

var symbolValues = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

func romanToInt(s string) (ans int) {
	n := len(s)
	fmt.Printf("s:%v\n",s)
	for i := range s {
		value := symbolValues[s[i]]
		fmt.Printf("value:%v\n",value)
		fmt.Printf("index:%v\n",i)
		fmt.Printf("n:%v\n",n)
		
		if i < n-1 && value < symbolValues[s[i+1]] {
			fmt.Printf("symbolValues[s[i+1]]:%v\n",symbolValues[s[i+1]])
			ans -= value
			fmt.Printf("ans:%v\n",ans)
		} else {
			ans += value
			fmt.Printf("ans:%v\n",ans)
		}
		fmt.Printf("=============\n")
	}
	return
}

func main() {
	s := "MCMXCIV"
	fmt.Printf("num:%v\n",romanToInt(s))
}
