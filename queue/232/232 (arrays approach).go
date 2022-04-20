package main

import "fmt"

// https://leetcode.com/problems/implement-queue-using-stacks/

type MyQueue struct {
	in  stack
	out stack
}

type stack struct {
	items []int
}

func Constructor() MyQueue {
	return MyQueue{
		stack{[]int{}},
		stack{[]int{}},
	}
}

func (q *MyQueue) Push(x int) {
	q.in.push(x)
}

func (s *stack) push(x int) {
	s.items = append(s.items, x)
}

func (q *MyQueue) Pop() int {
	for !q.in.empty() {
		q.out.push(q.in.pop())
	}
	front := q.out.pop()
	for !q.out.empty() {
		q.in.push(q.out.pop())
	}
	return front
}

func (s *stack) pop() int {
	l := len(s.items) - 1
	toRemove := s.items[l]
	s.items = s.items[:l]
	return toRemove
}

func (q *MyQueue) Peek() int {
	for !q.in.empty() {
		q.out.push(q.in.pop())
	}
	front := q.out.peek()
	for !q.out.empty() {
		q.in.push(q.out.pop())
	}
	return front
}

func (s *stack) peek() int {
	return s.items[len(s.items)-1]
}

func (q *MyQueue) Empty() bool {
	return q.in.empty()
}

func (s *stack) empty() bool {
	return len(s.items) == 0
}

func main() {
	obj := Constructor()
	obj.Push(3)
	obj.Push(4)
	obj.Push(5)
	obj.Push(6)
	fmt.Println(obj.Pop())
	obj.Push(7)
	fmt.Println(obj.Peek())
	obj.Push(8)
	fmt.Println(obj.Pop())
	fmt.Println(obj.Pop())
	fmt.Println(obj.Peek())
	fmt.Println(obj.Empty())
}
