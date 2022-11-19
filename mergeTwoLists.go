package main

import (
	"encoding/json"
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func mergeTwoLists(head1 *ListNode, head2 *ListNode) *ListNode {
	if head1 == nil {
		return head2
	}
	if head2 == nil {
		return head1
	}
	preHead := &ListNode{Val: -1} // 虚拟头结点
	pre := preHead
	for head1 != nil && head2 != nil {
		if head1.Val < head2.Val {
			pre.Next = head1
			head1 = head1.Next
			pre = pre.Next
		} else {
			pre.Next = head2
			head2 = head2.Next
			pre = pre.Next
		}
	}
	if head1 == nil {
		pre.Next = head2
	}
	if head2 == nil {
		pre.Next = head1
	}
	return preHead.Next
}
func main() {
	var head = new(ListNode)
	head.Val = 1
	var node1 = new(ListNode)
	node1.Val = 2
	head.Next = node1
	var node2 = new(ListNode)
	node2.Val = 7
	node1.Next = node2
	
	
	
	var head1 = new(ListNode)
	head1.Val = 3
	var node11 = new(ListNode)
	node11.Val = 4
	head1.Next = node11
	var node21 = new(ListNode)
	node21.Val = 5
	node11.Next = node21
	
	result := mergeTwoLists(head,head1)
	r,_ := json.Marshal(result)
	fmt.Printf("result:%v\n",string(r))
}