package main

import (
	"fmt"
)

func main() {
	l1 := mergeKLists([]*ListNode{&ListNode{1, nil}, &ListNode{2, nil}})
	fmt.Println(l1, l1.Next)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	result := new(ListNode)
	result = merge23(result, lists)
	return result
}
func merge23(result *ListNode, lists []*ListNode) *ListNode {
	index := -1
	min := 9999999999
	for i, listnode := range lists {
		if listnode != nil {
			if listnode.Val < min {
				min = listnode.Val
				index = i
			}
		}
	}
	if index == -1 {
		return nil
	}
	result.Next = new(ListNode)
	result.Val = min
	lists[index] = lists[index].Next
	result.Next = merge23(result.Next, lists)
	return result
}
