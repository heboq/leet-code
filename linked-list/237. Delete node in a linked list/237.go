package main

// https://leetcode.com/problems/delete-node-in-a-linked-list/

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type LinkedList struct {
	Head   *ListNode
	length int
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

func (node *ListNode) printStats() {
	fmt.Printf("\nDoing `*node = *node.Next`...\n")
	fmt.Printf(" node %11s: %v %5s: %T\n", "value", node, "type", node)
	fmt.Printf("*node %11s:  %v %5s:  %T\n", "value", *node, "type", *node)
	fmt.Printf(" node.Next %6s: %v %5s: %T\n", "value", node.Next, "type", node.Next)
	fmt.Printf("*node.Next %6s:  %v %5s:  %T\n", "value", *node.Next, "type", *node.Next)
	fmt.Println()
}

func deleteNode(node *ListNode) {
	// node.printStats()
	*node = *node.Next
}

func main() {
	var (
		head   = []int{4, 5, 1, 9, 8, 7, 6}
		myList = LinkedList{}
		// It is better to use map instead of slices below, but i was too tired.
		nodes  = make([]*ListNode, len(head))
		node   = new(ListNode)
		nodeID int
	)

	for i := len(head) - 1; i >= 0; i-- {
		node = &ListNode{Val: head[i]}
		nodes[i] = node
		myList.prepend(node)
	}

	myList.printData()
	printIDs(nodes)
	fmt.Print("Enter ID of the node you want to delete (exept tail node): ")
	if fmt.Scanln(&nodeID); nodeID < len(nodes)-1 {
		deleteNode(nodes[nodeID])
		myList.printData()
	} else {
		fmt.Println("\ninvalid input")
	}
}

func printIDs(n []*ListNode) {
	fmt.Print(0)
	for i := 1; i < len(n); i++ {
		fmt.Printf(" || %d", i)
	}
	fmt.Println(" - Indices")
}
