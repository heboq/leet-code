package main

// https://leetcode.com/problems/middle-of-the-linked-list/

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type LinkedList struct {
	Head   *ListNode
	length int
}

func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func main() {
	head := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 34, 23, 52, 6, 3}
	myList := LinkedList{}
	node := &ListNode{}

	for i := len(head) - 1; i >= 0; i-- {
		node = &ListNode{Val: head[i]}
		myList.prepend(node)
	}

	myList.printData()
	fmt.Println("Length of Linked List is ", myList.length)

	myList.Head = middleNode(node)
	myList.printData()
	myList.length = (myList.length >> 1) + 1
	fmt.Println("New length of Linked List is ", myList.length)
}

func (l *LinkedList) prepend(n *ListNode) {
	second := l.Head
	l.Head = n
	l.Head.Next = second
	l.length++
}

func (l LinkedList) printData() {
	toPrint := l.Head
	for l.length != 1 && toPrint.Next != nil {
		fmt.Printf("%d -> ", toPrint.Val)
		toPrint = toPrint.Next
		l.length--
	}
	fmt.Printf("%d\n", toPrint.Val)
}
