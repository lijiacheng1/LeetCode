package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2, -1, -1, -1}, 0))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	sliceList := []*ListNode{}
	lenList := 0
	for nowList := head; nowList != nil; nowList = nowList.Next {
		sliceList = append(sliceList,nowList)
		lenList++
	}
	if lenList==1{
		return nil
	}else if n==1 {
		sliceList[lenList-2].Next = nil
	}else if n == lenList {
		head = head.Next
	} else {
		sliceList[lenList-n-1].Next = sliceList[lenList-n+1]
	}
	return head
}
