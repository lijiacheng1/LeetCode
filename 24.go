package main

func main() {
	swapPairs()
}
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	rem := head
	if head != nil && head.Next != nil{
		head = head.Next
	}
	for ; rem != nil && rem.Next != nil;  {
		reverse24(rem)
		rem = rem.Next
	}
	return head
}

func reverse24(head *ListNode) {
	rem := head.Next
	if rem.Next.Next!= nil{
		head.Next = rem.Next.Next
	}else {
		head.Next = rem.Next
	}
	rem.Next = head

}