package main

import (
	"fmt"
	"math/rand"
	"time"
)

// MaxHeap definition
// if index begin at 0, left child is 2*i+1, right child is 2*i+2, parent is (i-1)/2
// if index begin at 1, left child is 2*i(i<<1), right child is 2*i+1(i<<1+1), parent is i/2(i>>1)
type MaxHeap struct {
	heap  []int
	count int
}

// NewMaxHeap constructor
func NewMaxHeap() MaxHeap {
	return MaxHeap{}
}

// Size method
func (mh *MaxHeap) Size() int {
	return mh.count
}

// IsEmpty method
func (mh *MaxHeap) IsEmpty() bool {
	return mh.count == 0
}

// Push method
func (mh *MaxHeap) Push(element int) {
	mh.heap = append(mh.heap, element)
	i := mh.count
	for p := (i - 1) / 2; element > mh.heap[p]; p = (i - 1) / 2 {
		mh.heap[i], mh.heap[p] = mh.heap[p], mh.heap[i]
		i = p
	}
	mh.count++
}

// Pop method, pop max of the heap
func (mh *MaxHeap) Pop() int {
	mh.heap[0], mh.heap[mh.count-1] = mh.heap[mh.count-1], mh.heap[0] // put max to last
	mh.count--                                                        // remove max
	for i := 0; 2*i+1 < mh.count; {                                   // if has at least one child
		maxChildIndex := 2*i + 1
		// if has right child
		// right child is 2*i+2, so is maxChildIndex+1
		if maxChildIndex+1 < mh.count && mh.heap[maxChildIndex+1] > mh.heap[maxChildIndex] {
			maxChildIndex = maxChildIndex + 1
		}
		if mh.heap[i] >= mh.heap[maxChildIndex] {
			break
		}
		mh.heap[i], mh.heap[maxChildIndex] = mh.heap[maxChildIndex], mh.heap[i]
		i = maxChildIndex
	}
	return mh.heap[mh.count]
}

func main() {
	rand.Seed(time.Now().UnixNano())
	mh := NewMaxHeap()
	t1 := time.Now()
	for i := 0; i < 10; i++ {
		mh.Push(rand.Intn(20))
	}
	for mh.count > 0 {
		mh.Pop()
	}
	fmt.Println(time.Now().Sub(t1))
	fmt.Println(mh)
}
