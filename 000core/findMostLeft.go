package main

import "fmt"

func main() {
	source := []int{3, 5, 5, 5, 6, 6, 7, 7, 7, 9, 9, 12}

	fmt.Println(findMostLeft(source, 0, len(source)-1, 7))
}

func findMostLeft(s []int, l, r, x int) int {
	if r >= l {
		mid := l + ((r - l) >> 1)
		if x <= s[mid] {
			return findMostLeft(s, l, mid-1, x)
		} else { // x > s[mid]
			return findMostLeft(s, mid+1, r, x)
		}
	}

	if s[l] == x {
		return l
	}

	return -1
}
