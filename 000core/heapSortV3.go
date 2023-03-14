package main

import (
	"fmt"
	"math/rand"
	"time"
)

type MaxHeap []int

func (mh *MaxHeap) Push(x int) {
	*mh = append(*mh, x)
	heapInsert(*mh, len(*mh)-1)
}

func heapInsert(heap []int, i int) { // go up
	for p := (i - 1) / 2; heap[i] > heap[p]; p = (i - 1) / 2 {
		heap[i], heap[p] = heap[p], heap[i]
		i = p
	}
}

func (mh *MaxHeap) Pop() (int, bool) {
	if len(*mh) == 0 {
		return 0, false
	}

	result := (*mh)[0]
	(*mh)[0], (*mh)[len(*mh)-1] = (*mh)[len(*mh)-1], (*mh)[0]
	*mh = (*mh)[:len(*mh)-1]
	heapify(*mh, 0) // index 0, no need to check go up, just go down

	return result, true
}

func heapify(heap []int, i int) { // go down
	for i*2+1 < len(heap) { // at least has a child
		maxChildIndex := i*2 + 1
		if maxChildIndex+1 < len(heap) && heap[maxChildIndex+1] > heap[maxChildIndex] {
			maxChildIndex = maxChildIndex + 1
		}
		if heap[maxChildIndex] > heap[i] {
			heap[i], heap[maxChildIndex] = heap[maxChildIndex], heap[i]
			i = maxChildIndex
		} else {
			break
		}
	}
}

func (mh *MaxHeap) Remove(i int) {
	if i >= len(*mh) {
		return
	}

	(*mh)[i], (*mh)[len(*mh)-1] = (*mh)[len(*mh)-1], (*mh)[i]
	*mh = (*mh)[:len(*mh)-1]
	// if go up, go down will not happen. because upper level always >= lower level
	heapInsert(*mh, i)
	heapify(*mh, i)
}

func (mh *MaxHeap) Reassign(i int, value int) {
	(*mh)[i] = value
	heapInsert(*mh, i)
	heapify(*mh, i)
}

func heapSort(array []int) {
	for i := len(array) - 1; i >= 0; i-- {
		heapify(array, i) // heapify an array
	}

	current := MaxHeap(array)
	for _, ok := current.Pop(); ok; _, ok = current.Pop() {
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	s := make([]int, 10)
	for i := range s {
		s[i] = rand.Intn(10)
	}
	fmt.Println("before:", s)
	t1 := time.Now()
	heapSort(s)
	fmt.Println(time.Now().Sub(t1))
	fmt.Println("after:", s)
}
