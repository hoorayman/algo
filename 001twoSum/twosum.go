package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		if subIndex, ok := m[target-v]; ok {
			return []int{i, subIndex}
		} else {
			m[v] = i
		}
	}
	return []int{-1, -1}
}

func main() {
	fmt.Println(twoSum([]int{1, 2, 3, 4, 5, 6}, 9))
}
