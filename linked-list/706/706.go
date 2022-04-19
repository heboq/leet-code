package main

// https://leetcode.com/problems/design-hashmap/

import (
	"fmt"
)

// ArraySize is the size of the HashTable (hash table / hash map).
const ArraySize = 5

// HashTable holds an array.
type HashTable struct {
	array [ArraySize]*bucket
}

// bucket is a linked list in each slot of the HashTable array.
type bucket struct {
	head *bucketNode
}

// bucketNode is a linked list node that holds a key.
type bucketNode struct {
	key, value int
	next       *bucketNode
}

// Put inserts a key and a value into the HashTable array.
func (ht *HashTable) Put(key int, value int) {
	index := hash(key)
	ht.array[index].put(key, value)
}

/* Get takes in a key and returns a value to which
 * the specified key of the HashTable is mapped. */
func (ht *HashTable) Get(key int) int {
	index := hash(key)
	return ht.array[index].get(key)
}

/* Remove takes in a key and removes it and it's
 * corresponding value from the HashTable. */
func (ht *HashTable) Remove(key int) {
	index := hash(key)
	ht.array[index].remove(key)
}

/* put takes in a key and a value, and inserts them into
 * newly created node in the bucket. */
func (b *bucket) put(k int, v int) {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k {
			currentNode.value = v
			return
		}
		currentNode = currentNode.next
	}
	newNode := &bucketNode{key: k, value: v}
	newNode.next = b.head
	b.head = newNode
}

/* get takes in a key and returns it's corresponding value
 * or -1 if this bucket contains no mapping for the key. */
func (b *bucket) get(k int) int {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k {
			return currentNode.value
		}
		currentNode = currentNode.next
	}
	return -1
}

/* remove takes in a key and removes from the bucket a node,
 * which contains this key and it's corresponding value. */
func (b *bucket) remove(k int) {
	if b.head != nil {
		// key is head of the bucket
		if b.head.key == k {
			b.head = b.head.next
			return
		}

		currentNode := b.head
		for currentNode.next != nil {
			if currentNode.next.key == k {
				// delete
				currentNode.next = currentNode.next.next
				return
			}
			currentNode = currentNode.next
		}
	}
}

// hash is hash function.
func hash(key int) int {
	return key % ArraySize
}

// Constructor creates a bucket in each slot of the hash table.
func Constructor() (result HashTable) {
	for i := range result.array {
		result.array[i] = new(bucket)
	}
	return
}

func main() {
	obj := Constructor()
	obj.Remove(2)
	// hashMap.printData()
	obj.Put(1, 1)
	// hashMap.printData()
	obj.Put(6, 1)
	obj.Put(2, 2)
	// hashMap.printData()
	fmt.Println(obj.Get(1))
	fmt.Println(obj.Get(2))
	fmt.Println(obj.Get(3))
	obj.Put(2, 1)
	// hashMap.printData()
	fmt.Println(obj.Get(2))
	obj.Remove(2)
	// hashMap.printData()
	fmt.Println(obj.Get(2))
}

// func (ht HashTable) printData() {
// 	toPrint := ht.array
// 	if ht.array == nil {
// 		fmt.Println()
// 		return
// 	}
// 	fmt.Print("[")
// 	for l.length != 1 && toPrint.next != nil {
// 		fmt.Printf("[%d, %d], ", toPrint.key, toPrint.value)
// 		toPrint = toPrint.next
// 		l.length--
// 	}
// 	fmt.Printf("[%d, %d]]\n", toPrint.key, toPrint.value)
// }
