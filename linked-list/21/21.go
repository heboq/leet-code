package main

// https://leetcode.com/problems/merge-two-sorted-lists/

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type LinkedList struct {
	Head   *ListNode
	length int
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) (head *ListNode) {
	if list1 == nil {
		head = list2
	} else if list2 == nil {
		head = list1
	} else if list1.Val > list2.Val {
		head = list2
		head.Next = mergeTwoLists(list1, list2.Next)
	} else {
		head = list1
		head.Next = mergeTwoLists(list1.Next, list2)
	}
	return
}

// func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
// 	dummy := new(ListNode)
// 	for node := dummy; l1 != nil || l2 != nil; node = node.Next {
// 		if l1 == nil {
// 			node.Next = l2
// 			break
// 		} else if l2 == nil {
// 			node.Next = l1
// 			break
// 		} else if l1.Val < l2.Val {
// 			node.Next = l1
// 			l1 = l1.Next
// 		} else {
// 			node.Next = l2
// 			l2 = l2.Next
// 		}
// 	}
// 	return dummy.Next
// }

func main() {

	list1 := []int{}
	list2 := []int{}
	myList1, myList2 := LinkedList{}, LinkedList{}
	mergedList := LinkedList{}

	for i := len(list1) - 1; i >= 0; i-- {
		node := &ListNode{Val: list1[i]}
		myList1.prepend(node)
	}

	for i := len(list2) - 1; i >= 0; i-- {
		node := &ListNode{Val: list2[i]}
		myList2.prepend(node)
	}

	myList1.printData()
	myList2.printData()
	mergedList.Head = mergeTwoLists(myList1.Head, myList2.Head)
	mergedList.printData()
	myList1.printData()
	myList2.printData()
}

func (l *LinkedList) prepend(n *ListNode) {
	second := l.Head
	l.Head = n
	l.Head.Next = second
	l.length++
}

func (l LinkedList) printData() {
	toPrint := l.Head
	if l.Head == nil {
		fmt.Println()
		return
	}
	for l.length != 1 && toPrint.Next != nil {
		fmt.Printf("%d -> ", toPrint.Val)
		toPrint = toPrint.Next
		l.length--
	}
	fmt.Printf("%d\n", toPrint.Val)
}
