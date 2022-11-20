/*
给你两个非空的链表，表示两个非负的整数。它们每位数字都是按照逆序的方式存储的，并且每个节点只能存储一位数字。
请将两个数相加，并以相同形式返回一个表示和的链表。
可以假设除了数字 0之外，这两个数都不会以0开头。
示例 1：
输入：l1 = [2,4,3], l2 = [5,6,4]
输出：[7,0,8]
解释：342 + 465 = 807.

示例 2：
输入：l1 = [0], l2 = [0]
输出：[0]
解释: 0 + 0 = 0

示例 3：
输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
输出：[8,9,9,9,0,0,0,1]
解释: 9999999 + 9999 = 10009998
 */
package main

import (
	"encoding/json"
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

/*
方法一：模拟
思路与算法
由于输入的两个链表都是逆序存储数字的位数的，因此两个链表中同一位置的数字可以直接相加。
我们同时遍历两个链表，逐位计算它们的和，并与当前位置的进位值相加。具体而言，如果当前两个链表处相应位置的数字为 n1,n2，进位值为carry，则它们的和为n1+n2+carry；其中，答案链表处相应位置的数字为(n1+n2+carry)/mod10，而新的进位值为
[n1+n2+carry]
—————————————
    10

如果两个链表的长度不同，则可以认为长度短的链表的后面有若干个 0。
此外，如果链表遍历结束后，有carry>0，还需要在答案链表的后面附加一个节点，节点的值为carry。
 */
func addTwoNumbers(l1, l2 *ListNode) (head *ListNode) {
	var tail *ListNode
	carry := 0
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10
		fmt.Printf("n1:%v\n",n1)
		fmt.Printf("n2:%v\n",n2)
		fmt.Printf("sum:%v\n",sum)
		fmt.Printf("carry:%v\n",carry)
		h,_ := json.Marshal(head)
		fmt.Printf("head:%v\n",string(h))
		if head == nil {
			head = &ListNode{Val: sum}
			tail = head
		} else {
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
		t,_ := json.Marshal(tail)
		fmt.Printf("tail:%v\n",string(t))
		fmt.Printf("*************************\n")
	}
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return
}
/*
复杂度分析
时间复杂度：O(\max(m,n))O(max(m,n))，其中 mm 和 nn 分别为两个链表的长度。我们要遍历两个链表的全部位置，而处理每个位置只需要 O(1)O(1) 的时间。
空间复杂度：O(1)O(1)。注意返回值不计入空间复杂度。
 */

func main() {
	
	var head = new(ListNode)
	head.Val = 3
	var node1 = new(ListNode)
	node1.Val = 1
	head.Next = node1
	var node2 = new(ListNode)
	node2.Val = 7
	node1.Next = node2
	
	
	
	var head1 = new(ListNode)
	head1.Val = 9
	var node11 = new(ListNode)
	node11.Val = 1
	head1.Next = node11
	var node21 = new(ListNode)
	node21.Val = 4
	node11.Next = node21
	
	i1,_ := json.Marshal(head)
	fmt.Printf("i1:%v\n",string(i1))
	
	i2,_ := json.Marshal(head1)
	fmt.Printf("i2:%v\n",string(i2))
	fmt.Printf("---------------------------\n")
	a := addTwoNumbers(head,head1)
	ap,_ := json.Marshal(a)
	fmt.Printf("a:%v\n",string(ap))
}
