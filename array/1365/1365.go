package main

// https://leetcode.com/problems/how-many-numbers-are-smaller-than-the-current-number/

import (
	"fmt"
	"sort"
)

func smallerNumbersThanCurrent(nums []int) []int {
	sortedNums := make([]int, len(nums))
	copy(sortedNums, nums)
	sort.Ints(sortedNums)
	countMap := map[int]int{}
	for i, num := range sortedNums {
		if _, ok := countMap[num]; !ok {
			countMap[num] = i
		}
	}
	counts := make([]int, len(nums))
	for i, num := range nums {
		counts[i] = countMap[num]
	}
	return counts
}

func main() {
	nums := []int{8, 1, 2, 2, 3}
	fmt.Print(smallerNumbersThanCurrent(nums))
}
