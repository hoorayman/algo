package main

import (
	"fmt"
)

func swapWithoutTemp(x, y int) (int, int) {
	// XOR:
	// 0 ^ N == N, N ^ N == 0
	// 满足交换律和结合律
	x = x ^ y
	y = x ^ y // now y == x^y^y == x
	x = x ^ y // now x == x^y^x == x^x^y == y
	return x, y
}

func main() {
	x, y := -3, 7
	fmt.Println("before: ", x, y)
	x, y = swapWithoutTemp(x, y)
	fmt.Println("after: ", x, y)
}
