package main

import (
	"container/heap"
	"go-algorithmic-learning/lib"
)

func main() {
	lib.ShowListNode(mergeKLists([]*lib.ListNode{
		lib.GetListNode([]int{-9, -7, -7}),
		lib.GetListNode([]int{-6, -4, -1, 1}),
		lib.GetListNode([]int{-6, -5, -2, 0, 0, 1, 2}),
		lib.GetListNode([]int{-9, -8, -6, -5, -4, 1, 2, 4}),
		lib.GetListNode([]int{-10}),
		lib.GetListNode([]int{-5, 2, 3}),
	}))
}

type ListsHeap []*lib.ListNode

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
	*l = append(*l, x.(*lib.ListNode))
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

// 小顶堆直接使用仓库
func mergeKListsV1(lists []*lib.ListNode) *lib.ListNode {
	h := ListsHeap(lists)
	heap.Init(&h)
	ansHeader := &lib.ListNode{}
	itemH := ansHeader
	for len(h) > 0 {
		item := heap.Pop(&h).(*lib.ListNode)
		itemH.Next = item
		itemH = itemH.Next
		if item.Next != nil {
			heap.Push(&h, item.Next)
		}
	}
	return ansHeader.Next
}

// 小顶堆手动实现

func mergeKLists(lists []*lib.ListNode) *lib.ListNode {
	createHeap(&lists)
	ansHeader := &lib.ListNode{}
	itemH := ansHeader
	for len(lists) > 0 {
		listItem := popHeap(&lists)
		itemH.Next = listItem
		itemH = itemH.Next
		if listItem.Next != nil {
			addHeap(&lists, listItem.Next)
		}
	}
	return ansHeader.Next
}

func popHeap(lists *[]*lib.ListNode) *lib.ListNode {
	if lists == nil || len(*lists) == 0 {
		return nil
	}
	ans := (*lists)[0]
	(*lists)[0] = (*lists)[len(*lists)-1]
	*lists = (*lists)[0 : len(*lists)-1]
	rebuildHeal(lists)
	return ans
}

func createHeap(lists *[]*lib.ListNode) {
	if lists == nil {
		return
	}
	newList := make([]*lib.ListNode, 0, len(*lists))
	for _, item := range *lists {
		if item == nil {
			continue
		}
		addHeap(&newList, item)
	}
	*lists = newList
}

func rebuildHeal(lists *[]*lib.ListNode) {
	index := 0
	length := len(*lists)
	for index < length {
		leftIndex := index*2 + 1
		rightIndex := index*2 + 2
		if leftIndex > length-1 || ((*lists)[index].Val <= (*lists)[leftIndex].Val && (rightIndex >= length || (*lists)[index].Val <= (*lists)[rightIndex].Val)) {
			break
		}
		needIndex := -1
		if rightIndex > length-1 || (*lists)[leftIndex].Val < (*lists)[rightIndex].Val {
			needIndex = leftIndex
		} else {
			needIndex = rightIndex
		}
		(*lists)[index], (*lists)[needIndex] = (*lists)[needIndex], (*lists)[index]
		index = needIndex
	}
}

func addHeap(lists *[]*lib.ListNode, item *lib.ListNode) {
	if lists == nil {
		return
	}
	*lists = append(*lists, item)
	index := len(*lists) - 1
	for index > 0 {
		beforeIndex := 0
		if index%2 == 0 {
			beforeIndex = index/2 - 1
		} else {
			beforeIndex = index / 2
		}
		if (*lists)[beforeIndex].Val > (*lists)[index].Val {
			(*lists)[beforeIndex], (*lists)[index] = (*lists)[index], (*lists)[beforeIndex]
			index = beforeIndex
			continue
		}
		break
	}
}
