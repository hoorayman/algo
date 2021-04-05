package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 已知数组相邻元素两两不等，找出局部最小
func localMin(s []int) int {
	l, r := 0, len(s)-1
	if r < 1 { // length of s <= 1
		return -1
	}
	if s[1] > s[0] { // index 0 is local min
		return 0
	}
	if s[r-1] > s[r] { // index len(s)-1 is local min
		return r
	}
	// 数组的左端是向下趋势，右端是向上趋势，且数组相邻元素两两不等，所以数组中间必有局部最小
	for l < r {
		mid := l + ((r - l) >> 1) // '>>' is faster than '/'. Do not use (l+r)/2, overflow could happen if r is big enough
		if s[mid-1] > s[mid] && s[mid] < s[mid+1] {
			return mid
		} else if s[mid-1] < s[mid] {
			r = mid
		} else if s[mid] > s[mid+1] {
			l = mid
		}
	}
	return -1
}

func main() {
	var s []int
	rand.Seed(time.Now().UnixNano())
	n := 2 + rand.Intn(16)
	for i := 0; i < n; i++ {
		s = append(s, i)
	}
	// shuffle
	for i := len(s) - 1; i > 0; i-- { // Fisher–Yates shuffle
		j := rand.Intn(i + 1)
		s[i], s[j] = s[j], s[i]
	}
	fmt.Println(s, localMin(s))
}
