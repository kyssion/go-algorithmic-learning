package main

import "go-algorithmic-learning/lib"

/**
给你两个非空 的链表，表示两个非负的整数。它们每位数字都是按照逆序的方式存储的，并且每个节点只能存储一位数字。

请你将两个数相加，并以相同形式返回一个表示和的链表。

你可以假设除了数字 0 之外，这两个数都不会以 0开头。

输入：l1 = [2,4,3], l2 = [5,6,4]
输出：[7,0,8]
解释：342 + 465 = 807.
示例 2：

输入：l1 = [0], l2 = [0]
输出：[0]
示例 3：

输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
输出：[8,9,9,9,0,0,0,1]


提示：

每个链表中的节点数在范围 [1, 100] 内
0 <= Node.val <= 9
题目数据保证列表表示的数字不含前导零


*/

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
