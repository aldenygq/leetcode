package main

import (
	"encoding/json"
	"fmt"
)
type ListNode struct {
	Val int
	Next *ListNode
}

func getLength(head *ListNode) (length int) {
	for ; head != nil; head = head.Next {
		length++
	}
	return
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	length := getLength(head)
	fmt.Printf("length:%v\n",length)
	dummy := &ListNode{0, head}
	d,_ := json.Marshal(dummy)
	fmt.Printf("d:%v\n",string(d))
	cur := dummy
	c,_ := json.Marshal(cur)
	fmt.Printf("cur:%v\n",string(c))
	for i := 0; i < length-n; i++ {
		cur = cur.Next
		cu,_ := json.Marshal(cur)
		fmt.Printf("new cur:%v\n",string(cu))
	}
	current,_ := json.Marshal(cur)
	fmt.Printf("current:%v\n",string(current))
	cur.Next = cur.Next.Next
	curr,_ := json.Marshal(cur)
	fmt.Printf("end cur:%v\n",string(curr))
	return dummy.Next
}
func main() {
	var head1 = new(ListNode)
	head1.Val = 3
	var node11 = new(ListNode)
	node11.Val = 4
	head1.Next = node11
	var node21 = new(ListNode)
	node21.Val = 5
	node11.Next = node21
	
	result := removeNthFromEnd(head1,4)
	r ,_ := json.Marshal(result)
	fmt.Printf("r:%v",string(r))
}