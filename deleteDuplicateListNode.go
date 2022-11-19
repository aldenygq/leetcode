package main

import (
	"fmt"
	"encoding/json"
)
type ListNode struct {
	Val int
	Next *ListNode
}
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	
	cur := head
	for cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	
	return head
}

func delete2Duplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	
	dummy := &ListNode{0, head}
	
	cur := dummy
	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			x := cur.Next.Val
			for cur.Next != nil && cur.Next.Val == x {
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}
	}
	
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
	var node31 = new(ListNode)
	node31.Val = 5
	node21.Next = node31
	var node41 = new(ListNode)
	node41.Val = 6
	node31.Next = node41
	var node51 = new(ListNode)
	node51.Val = 6
	node41.Next = node51
	var node61 = new(ListNode)
	node61.Val = 7
	node51.Next = node61
	var node71 = new(ListNode)
	node71.Val = 7
	node61.Next = node71
	/*
	result := deleteDuplicates(head1)
	r ,_ := json.Marshal(result)
	fmt.Printf("r:%v\n",string(r))
	
	 */
	result2 := delete2Duplicates(head1)
	r2 ,_ := json.Marshal(result2)
	fmt.Printf("r2:%v\n",string(r2))
}
