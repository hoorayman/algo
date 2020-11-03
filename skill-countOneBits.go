package main

import (
	"fmt"
)

func countOneBits(x int) int {
	count := 0
	for x != 0 {
		rightMostOne := x & ((^x) + 1) // x & ((^x) + 1) to get the right most one. For example, 00110010 right most one is 00000010
		x = x ^ rightMostOne           // set right most one to zero
		count++
	}
	return count
}

func main() {
	n := 0xEAB0
	fmt.Println(countOneBits(n))
}
