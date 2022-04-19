package main

// https://leetcode.com/problems/reverse-linked-list/

// I have spended 40 minutes

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type LinkedList struct {
	Head   *ListNode
	length int
}

// Initially prev = nil, because of its implicit declaration
// but because `prev` points to a struct, the latter can be assigned to it.
func reverseList(head *ListNode) (prev *ListNode) {
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	/* or
	for head != nil {
		head.Next, prev, head = prev, head, head.Next
	}
	*/
	return prev
}

func main() {

	head := []int{1, 2, 3, 4, 5, 6}
	myList := LinkedList{}
	node := new(ListNode)

	for i := len(head) - 1; i >= 0; i-- {
		node = &ListNode{Val: head[i]}
		myList.prepend(node)
	}

	myList.printData()
	myList.Head = reverseList(node)
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
