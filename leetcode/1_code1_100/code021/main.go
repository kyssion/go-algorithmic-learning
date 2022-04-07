package main

import "go-algorithmic-learning/lib"

func main() {
	lib.ShowListNode(mergeTwoLists(lib.GetListNode([]int{1, 3, 5}), lib.GetListNode([]int{
		1, 2, 4,
	})))
}

func mergeTwoLists(list1 *lib.ListNode, list2 *lib.ListNode) *lib.ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	header := &lib.ListNode{}
	extOne := header
	for {
		if list1.Val < list2.Val {
			extOne.Next = list1
			list1 = list1.Next
		} else {
			extOne.Next = list2
			list2 = list2.Next
		}
		extOne = extOne.Next
		if list1 == nil {
			extOne.Next = list2
			break
		}
		if list2 == nil {
			extOne.Next = list1
			break
		}
	}
	return header.Next
}
