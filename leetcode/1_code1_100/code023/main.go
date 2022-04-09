package main

import (
	"container/heap"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

type ListsHeap []*ListNode

func (l ListsHeap) Len() int {
	return len(l)
}

func (l ListsHeap) Less(i, j int) bool {
	return l[i].Val < l[j].Val
}

func (l ListsHeap) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l *ListsHeap) Push(x any) {
	if l == nil {
		return
	}
	*l = append(*l, x.(*ListNode))
}

func (l *ListsHeap) Pop() any {
	if l == nil {
		return nil
	}
	old := *l
	n := len(old)
	x := old[n-1]
	*l = old[0 : n-1]
	return x
}

func mergeKLists(lists []*ListNode) *ListNode {
	h := ListsHeap(lists)
	heap.Init(&h)
	ansHeader := &ListNode{}
	itemH := ansHeader
	for len(h) > 0 {
		item := heap.Pop(&h).(*ListNode)
		itemH.Next = item
		itemH = itemH.Next
		if item.Next != nil {
			heap.Push(&h, item.Next)
		}
	}
	return ansHeader.Next
}
