package main

import "go-algorithmic-learning/lib"

func main() {
	lib.ShowListNode(reverseKGroup(lib.GetListNode([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}), 3))
}

type ListNode struct {
	Val  int
	Next *lib.ListNode
}

func reverseKGroup(head *lib.ListNode, k int) *lib.ListNode {
	if k == 1 {
		return head
	}
	ans := &lib.ListNode{}
	index := ans
	index.Next = head
	leftNode := index.Next
	rightNode := leftNode
	for leftNode != nil {
		for i := 1; i < k && rightNode != nil; i++ {
			rightNode = rightNode.Next
		}
		endNode := rightNode
		rightNode = leftNode.Next
		nextIndex := leftNode
		for leftNode != endNode && rightNode != nil {
			nextNode := rightNode.Next
			rightNode.Next = leftNode
			leftNode = rightNode
			rightNode = nextNode
		}
		index.Next.Next = rightNode
		index.Next = leftNode
		index = nextIndex
		leftNode = rightNode
	}
	return ans.Next
}
