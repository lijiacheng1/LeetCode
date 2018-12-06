package main

import (
	"fmt"
)

func main() {
	fmt.Println(mergeTwoLists(nil, nil))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	rem := new(ListNode)
	result := rem
	for ; l1 != nil && l2 != nil; {
		if l1.Val < l2.Val {
			rem.Next = l1
			l1 = l1.Next
		} else {
			rem.Next = l2
			l2 = l2.Next
		}
		rem = rem.Next
	}
	if l1 == nil {
		rem.Next = l2
	} else {
		rem.Next = l1
	}
	return result.Next
}
