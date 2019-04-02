package main

func main() {
	reverseKGroup(nil, 2)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	searchk := head
	for i := 0;i<k;i++{
		searchk = searchk.Next
	}
	result := searchk
	rem := head
	list := []*ListNode{}
	remTail := head
	for ; rem != nil; {
		for i := 0; i < k; i++ {
			list[i] = rem
			rem = rem.Next
			if i == k-1 {
				for u := 1; u < i; u++ {
					list[u].Next = list[u-1]
				}
				remTail.Next = list[k-1]
				remTail = list[0]
			} else if rem == nil {
				remTail.Next = list[0]
				break
			}
		}
	}
	return result
}
