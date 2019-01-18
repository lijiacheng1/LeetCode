package main

func main() {
	swapPairs(nil)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	result := head
	if head == nil || head.Next == nil {
		return result
	} else {
		result = head.Next
	}
	remN, remP, cir := head, head, head
	couter := 0
	for ; cir != nil; {
		couter++
		if couter%2 != 0 {
			cir = cir.Next
		} else {
			//remN记录3的地址，不管3存不存在
			remP = remN
			remN = cir.Next
			cir.Next = remP
			if remN != nil && remN.Next != nil {
				remP.Next = remN.Next
			} else if remN != nil {
				remP.Next = remN
			} else {
				remP.Next = nil
			}
			cir = remN
		}
	}
	return result
}
