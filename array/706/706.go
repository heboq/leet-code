package main

// https://leetcode.com/problems/design-hashmap/

import "fmt"

type MyHashMap struct {
	ht []struct {
		key, value int
	}
}

func Constructor() (new MyHashMap) {
	return
}

func (hm *MyHashMap) Put(key int, value int) {
	for id, entry := range hm.ht {
		if entry.key == key {
			hm.ht[id].value = value
			return
		}
	}
	hm.ht = append(hm.ht, struct{ key, value int }{key, value})
}

func (hm *MyHashMap) Get(key int) int {
	for _, entry := range hm.ht {
		if entry.key == key {
			return entry.value
		}
	}
	return -1
}

func (hm *MyHashMap) Remove(key int) {
	for id, entry := range hm.ht {
		if entry.key == key {
			hm.ht = append(hm.ht[:id], hm.ht[id+1:]...)
		}
	}
}

func main() {
	obj := Constructor()
	fmt.Println(obj)
	obj.Put(1, 1)
	fmt.Println(obj)
	obj.Put(2, 2)
	fmt.Println(obj)
	fmt.Println(obj.Get(1))
	fmt.Println(obj.Get(2))
	fmt.Println(obj.Get(3))
	obj.Put(2, 1)
	fmt.Println(obj)
	fmt.Println(obj.Get(2))
	obj.Remove(2)
	fmt.Println(obj)
	fmt.Println(obj.Get(2))
}
