package main

import "go-algorithmic-learning/lib"

func main() {
	nodelist := lib.GetListNode([]int{1, 2})
	lib.ShowListNode(removeNthFromEnd(nodelist, 1))
}

// todo 这里有一个地方 需要注意一下 ， 就是边界问题， i = 0 遍历的时候 = 是n+1 < 是n ， node节点也是这样的
func removeNthFromEnd(head *lib.ListNode, n int) *lib.ListNode {
	if n == 0 {
		return head
	}
	startHeader := &lib.ListNode{
		Next: head,
	}
	skipNode := startHeader
	for i, baseNode := 0, head; baseNode != nil; i++ {
		baseNode = baseNode.Next
		if i >= n {
			skipNode = skipNode.Next
		}
	}
	skipNode.Next = skipNode.Next.Next
	return startHeader.Next
}
