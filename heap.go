package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// MinHeap definition
type MinHeap struct {
	heap  []int
	count int
}

// New constructor
func New() MinHeap {
	mh := MinHeap{}
	return mh
}

// Size method
func (mh *MinHeap) Size() int {
	return mh.count
}

// IsEmpty method
func (mh *MinHeap) IsEmpty() bool {
	return mh.count == 0
}

// Push method
func (mh *MinHeap) Push(element int) {
	if len(mh.heap) <= mh.count {
		mh.heap = append(mh.heap, element)
	}
	mh.count++
	i := mh.count
	for i > 1 && mh.heap[i/2-1] > element { // when index begin at 1, parent is i/2
		mh.heap[i-1], mh.heap[i/2-1] = mh.heap[i/2-1], mh.heap[i-1]
		i /= 2
	}
}

// PopMin method
func (mh *MinHeap) PopMin() (int, error) {
	if mh.count <= 0 {
		return 0, errors.New("pop nil heap")
	}
	element := mh.heap[mh.count-1]
	mh.heap[0], mh.heap[mh.count-1] = mh.heap[mh.count-1], mh.heap[0]
	mh.count--
	for i := 1; 2*i <= mh.count; { // when index begin at 1, left child is 2i, right child is 2i+1
		minChildIndex := 2*i - 1
		if 2*i+1 <= mh.count && mh.heap[2*i] < mh.heap[2*i-1] {
			minChildIndex = 2 * i
		}
		if element > mh.heap[minChildIndex] {
			mh.heap[i-1], mh.heap[minChildIndex] = mh.heap[minChildIndex], mh.heap[i-1]
			i = minChildIndex + 1
		} else {
			break
		}
	}
	return mh.heap[mh.count], nil
}

func main() {
	mh := New()
	t1 := time.Now()
	for i := 0; i < 100000; i++ {
		mh.Push(rand.Int())
	}
	for _, e := mh.PopMin(); e == nil; {
		_, e = mh.PopMin()
	}
	fmt.Println(time.Now().Sub(t1))
}
