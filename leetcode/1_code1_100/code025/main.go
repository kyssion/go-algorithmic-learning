package main

import "go-algorithmic-learning/lib"

func main() {
	lib.ShowListNode(reverseKGroupV2(lib.GetListNode([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}), 3))
}

type ListNode struct {
	Val  int
	Next *lib.ListNode
}

// 方法1 外部变量 优化代码 超出的部分不进行转化
func reverseKGroupV1(head *lib.ListNode, k int) *lib.ListNode {
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
		//if rightNode == nil {
		//	break
		//}
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

// 方法2 循环里面有变量
func reverseKGroupV2(head *lib.ListNode, k int) *lib.ListNode {
	if k <= 1 || head == nil {
		return head
	}
	ans := &lib.ListNode{}
	index := ans
	index.Next = head
	for {
		endNode := index.Next
		for i := 1; i < k && endNode != nil; i++ {
			endNode = endNode.Next
		}
		if endNode == nil {
			break
		}
		nextStart := index.Next
		startNode := index.Next
		nextNode := startNode.Next
		nextToNode := nextNode.Next
		for startNode != endNode {
			nextToNode = nextNode.Next
			nextNode.Next = startNode
			startNode = nextNode
			nextNode = nextToNode
		}
		index.Next.Next = nextToNode
		index.Next = startNode
		index = nextStart
	}
	return ans.Next
}
