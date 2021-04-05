package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func binarySearch(s []int, l, r, x int) int {
	if r >= l {
		mid := l + ((r - l) >> 1) // '>>' is faster than '/'. Do not use (l+r)/2, overflow could happen if r is big enough
		if s[mid] == x {
			return mid
		} else if x < s[mid] {
			return binarySearch(s, l, mid-1, x)
		} else { // s[mid] < x
			return binarySearch(s, mid+1, r, x)
		}
	}
	return -1
}

func main() {
	s := make([]int, 15)
	for i := range s {
		s[i] = rand.Intn(100)
	}
	sort.Ints(s)
	fmt.Println(s, binarySearch(s, 0, len(s)-1, 62))
}
