package main

import "go-algorithmic-learning/lib"

func main() {

}

func swapPairs(head *lib.ListNode) *lib.ListNode {
	ans := &lib.ListNode{}
	indexItem := ans
	indexItem.Next = head
	for head != nil && head.Next != nil {
		indexItem.Next = head.Next
		head.Next = head.Next.Next
		indexItem.Next.Next = head
		indexItem = indexItem.Next.Next
		head = indexItem.Next
	}
	return ans.Next
}
