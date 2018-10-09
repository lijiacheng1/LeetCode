package main

import "fmt"

func main() {
	l1next := &ListNode{9, nil}
	l1 := &ListNode{9, l1next}
	//	l2next := &ListNode{5, nil}
	l2 := &ListNode{1, nil}
	fmt.Println(addTwoNumbers(l1, l2))
	fmt.Println(addTwoNumbers(l1, l2).Next)
	fmt.Println(addTwoNumbers(l1, l2).Next.Next)
	fmt.Println(l1);
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	add := 0
	l3 := &ListNode{}
	pre := l3
	next := l3
	for {
		if l1 != nil || l2 != nil {

			if l1 == nil {
				l1 = &ListNode{}
				l1.Val = 0
			}
			if l2 == nil {
				l2 = &ListNode{}
				l2.Val = 0
			}
			he := l1.Val + l2.Val + add
			if he >= 10 {
				next.Val = he - 10
				add = 1
			} else {
				next.Val = he
				add = 0
			}
			pre = next
			next = &ListNode{}
			pre.Next = next
			l1 = l1.Next
			l2 = l2.Next
		} else {
			pre.Next = nil
			break
		}
	}
	if add != 0 {
		next.Val = add
		pre.Next = next
		add = 0
	}
	return l3
}
