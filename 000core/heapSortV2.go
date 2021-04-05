package main

import (
	"fmt"
	"math/rand"
	"time"
)

// if index begin at 0, left child is 2*i+1, right child is 2*i+2, parent is (i-1)/2
// if index begin at 1, left child is 2*i(i<<1), right child is 2*i+1(i<<1|1), parent is i/2(i>>1)

// 1.change array to a max heap by heapify from tail to head. O(N)
// 2.loop pop to get the sorted array
func heapSortV2(a []int) {
	for i := len(a) - 1; i >= 0; i-- { // this is the way that heapify an array
		heapify(a, i, len(a))
	}
	count := len(a)
	for count > 0 {
		a[0], a[count-1] = a[count-1], a[0]
		count--
		heapify(a, 0, count)
	}
}

func heapify(array []int, beginIndex int, heapSize int) {
	for i := beginIndex; 2*i+1 < heapSize; { // at least has a child
		maxChildIndex := 2*i + 1
		if maxChildIndex+1 < heapSize && array[maxChildIndex+1] > array[maxChildIndex] { // if has right child and right child is bigger than left child
			maxChildIndex = maxChildIndex + 1
		}
		if array[i] >= array[maxChildIndex] {
			break
		}
		array[i], array[maxChildIndex] = array[maxChildIndex], array[i]
		i = maxChildIndex
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	s := make([]int, 100000)
	for i := range s {
		s[i] = rand.Int()
	}
	t1 := time.Now()
	heapSortV2(s)
	fmt.Println(time.Now().Sub(t1))
	// fmt.Println(s)
}
