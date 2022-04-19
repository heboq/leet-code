package main

// https://leetcode.com/problems/remove-linked-list-elements/

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type LinkedList struct {
	Head   *ListNode
	length int
}

func removeElements(head *ListNode, val int) *ListNode {
	if head.Next == nil {
		return nil
	}
	dummyHead := &ListNode{-1, head}
	currentNode := dummyHead

	for currentNode.Next != nil {
		if currentNode.Next.Val == val {
			currentNode.Next = currentNode.Next.Next
		} else {
			currentNode = currentNode.Next
		}
	}
	return dummyHead.Next
}

func main() {
	// head := []int{6, 6, 6, 6, 6, 6, 6}
	// head := []int{1, 2, 6, 3, 4, 5, 6}
	head := []int{}
	myList := LinkedList{}
	node := new(ListNode)

	for i := len(head) - 1; i >= 0; i-- {
		node = &ListNode{Val: head[i]}
		myList.prepend(node)
	}

	myList.printData()
	fmt.Println(removeElements(node, 6))
	myList.Head = removeElements(node, 6)
	myList.printData()
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
