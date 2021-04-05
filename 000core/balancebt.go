package main

import (
	"fmt"
	"math"
)

type node struct {
	val   int
	left  *node
	right *node
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func process(np *node) (isBalance bool, height int) { // get isBalance and tree height from each sub tree
	if np == nil {
		return true, 0
	}
	lBalance, lHeight := process(np.left)
	rBalance, rHeight := process(np.right)
	if lBalance && rBalance && math.Abs(float64(lHeight-rHeight)) < 2.0 {
		return true, max(lHeight, rHeight) + 1
	}
	return false, max(lHeight, rHeight) + 1
}

func isBalance(np *node) (balance bool) {
	balance, _ = process(np)
	return
}

func main() {
	nl := []node{node{1, nil, nil}, node{2, nil, nil}, node{3, nil, nil}, node{4, nil, nil}}
	// nl[0].left = &nl[1]
	nl[0].right = &nl[2]
	nl[2].right = &nl[3]
	fmt.Println(isBalance(&nl[0]))
}
