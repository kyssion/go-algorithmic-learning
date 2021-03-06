package lib

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func GetListNode(intArr []int) *ListNode {
	baseNode := &ListNode{}
	fatherNode := baseNode

	for _, item := range intArr {
		nowNode := &ListNode{
			Val: item,
		}
		baseNode.Next = nowNode
		baseNode = nowNode
	}
	return fatherNode.Next
}

func ShowListNode(listNode *ListNode) {
	for {
		if listNode == nil {
			break
		}
		fmt.Printf("%d ", listNode.Val)
		listNode = listNode.Next
	}
	println()
}

func ShowIntArr(arrs [][]int) {
	for _, arrItem := range arrs {
		for _, i := range arrItem {
			fmt.Printf("%d", i)
		}
		println()
	}
}

func ShowStrArr(arrs [][]string) {
	for _, arrItem := range arrs {
		for _, i := range arrItem {
			fmt.Printf("%s ", i)
		}
		println()
	}
}
