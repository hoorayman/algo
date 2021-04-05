package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// get nearest left number which greater or equal to a number in an ordered list
func nearestLeft(s []int, l, r, x int) int {
	if r > l {
		mid := l + ((r - l) >> 1) // '>>' is faster than '/'. Do not use (l+r)/2, overflow could happen if l or r is big enough
		if s[mid] < x {
			return nearestLeft(s, mid+1, r, x)
		}
		return nearestLeft(s, l, mid, x) // x <= s[mid]
	} else if r == l && x <= s[l] {
		return l
	}
	return -1
}

func main() {
	s := make([]int, 15)
	for i := range s {
		s[i] = rand.Intn(10)
	}
	sort.Ints(s)
	fmt.Println(s, nearestLeft(s, 0, len(s)-1, -100))
}
