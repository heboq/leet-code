package main

import "fmt"

// https://leetcode.com/problems/linked-list-cycle/

type ListNode struct {
	Val  int
	Next *ListNode
}

type LinkedList struct {
	head   *ListNode
	length int
}

func hasCycle(head *ListNode) bool {
	nodes := make(map[*ListNode]struct{})
	for head != nil {
		if _, found := nodes[head]; found {
			return true
		}
		nodes[head] = struct{}{}
		head = head.Next
	}
	return false
}

// setPointerTo sets pointer from a tail node to target one.
func (tail *ListNode) setPointerTo(target *ListNode) {
	tail.Next = target
}

func main() {
	list := []int{3, 2, 0, -4}
	myList := LinkedList{}
	nodes := []*ListNode{}
	pos := 1

	for i := len(list) - 1; i >= 0; i-- {
		node := &ListNode{Val: list[i]}
		myList.prepend(node)
	}

	myList.printData()

	if pos != -1 {
		currentNode := myList.head
		for currentNode != nil {
			nodes = append(nodes, currentNode)
			if currentNode.Next == nil {
				currentNode.setPointerTo(nodes[pos])
				break
			}
			currentNode = currentNode.Next
		}
	}

	fmt.Println(hasCycle(myList.head))
}

func (l *LinkedList) prepend(node *ListNode) {
	node.Next = l.head
	l.head = node
	l.length++
}

func (l LinkedList) printData() {
	toPrint := l.head
	if l.head == nil {
		fmt.Println()
		return
	}

	for toPrint != nil && l.length != 1 {
		fmt.Printf("%d -> ", toPrint.Val)
		toPrint = toPrint.Next
		l.length--
	}

	fmt.Printf("%d\n", toPrint.Val)
}
