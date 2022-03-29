package main

import "go-algorithmic-learning/lib"

func main() {
	listNode := addTwoNumbers(lib.GetListNode([]int{
		9, 9, 9, 9,
	}), lib.GetListNode([]int{
		1,
	}))

	lib.ShowListNode(listNode)
}

func addTwoNumbers(l1 *lib.ListNode, l2 *lib.ListNode) *lib.ListNode {
	ans := &lib.ListNode{}
	fatherNode := ans
	step := 0
	for {
		if l1 == nil && l2 == nil {
			break
		}
		item1 := 0
		if l1 != nil {
			item1 = l1.Val
			l1 = l1.Next
		}
		item2 := 0
		if l2 != nil {
			item2 = l2.Val
			l2 = l2.Next
		}
		item1 = item1 + item2 + step
		ansItem := &lib.ListNode{}
		if item1 >= 10 {
			step = 1
			ansItem.Val = item1 - 10
		} else {
			ansItem.Val = item1
			step = 0
		}
		ans.Next = ansItem
		ans = ansItem
	}
	if step != 0 {
		ans.Next = &lib.ListNode{
			Val:  step,
			Next: nil,
		}
	}
	return fatherNode.Next
}
