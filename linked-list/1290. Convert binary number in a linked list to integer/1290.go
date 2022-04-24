package main

// https://leetcode.com/problems/convert-binary-number-in-a-linked-list-to-integer/

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type LinkedList struct {
	Head   *ListNode
	length int
}

func getDecimalValue(head *ListNode) (sum int) {
	for head != nil {
		sum = (sum << 1) + head.Val
		head = head.Next
	}
	return
}

func (l *LinkedList) prepend(n *ListNode) {
	second := l.Head
	l.Head = n
	l.Head.Next = second
	l.length++
}

func main() {
	head := []int{1, 0, 0, 0}
	myList := LinkedList{}
	node := &ListNode{}
	for i := len(head) - 1; i >= 0; i-- {
		node = &ListNode{Val: head[i]}
		myList.prepend(node)
	}
	fmt.Println(getDecimalValue(node))
}
