package main

import (
	"fmt"
	"math/rand"
	"time"
)

// if index begin at 0, left child is 2*i+1, right child is 2*i+2, parent is (i-1)/2
// if index begin at 1, left child is 2*i(i<<1), right child is 2*i+1(i<<1|1), parent is i/2(i>>1)

// 1.take array as a max heap
// 2.fake push element to heap
// 3.loop pop to get the sorted array
func heapSort(a []int) {
	for i := range a { // fake push element to heap
		for j, p := i, (i-1)/2; a[j] > a[p]; p = (j - 1) / 2 {
			a[j], a[p] = a[p], a[j]
			j = p
		}
	}
	count := len(a)
	for count > 0 {
		a[0], a[count-1] = a[count-1], a[0]
		count--
		for i := 0; 2*i+1 < count; { // at least has a child
			maxChildIndex := 2*i + 1
			if maxChildIndex+1 < count && a[maxChildIndex+1] > a[maxChildIndex] { // if has right child and right child is bigger than left child
				maxChildIndex = maxChildIndex + 1
			}
			if a[i] >= a[maxChildIndex] {
				break
			}
			a[i], a[maxChildIndex] = a[maxChildIndex], a[i]
			i = maxChildIndex
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	s := make([]int, 100000)
	for i := range s {
		s[i] = rand.Intn(10000)
	}
	t1 := time.Now()
	heapSort(s)
	fmt.Println(time.Now().Sub(t1))
	// fmt.Println(s)
}
