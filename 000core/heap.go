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
	heapInsert(mh.heap, i)
	mh.count++
}

func heapInsert(array []int, beginIndex int) { // go up
	for i, p := beginIndex, (beginIndex-1)/2; array[i] > array[p]; p = (i - 1) / 2 {
		array[i], array[p] = array[p], array[i]
		i = p
	}
}

// Pop method, pop max of the heap
func (mh *MaxHeap) Pop() int {
	mh.heap[0], mh.heap[mh.count-1] = mh.heap[mh.count-1], mh.heap[0] // put max to last
	mh.count--                                                        // remove max
	heapify(mh.heap, 0, mh.count)
	return mh.heap[mh.count]
}

func heapify(array []int, beginIndex int, heapSize int) { // go down
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

// Remove method, remove the i-th element of the heap
func (mh *MaxHeap) Remove(i int) {
	mh.heap[i], mh.heap[mh.count-1] = mh.heap[mh.count-1], mh.heap[i] // put i-th to last
	mh.count--                                                        // remove i-th
	if i >= mh.count {
		return
	}
	heapInsert(mh.heap, i)        // check go up
	heapify(mh.heap, i, mh.count) // check go down
	// if go up, go down will not happen. because upper level always >= lower level
}

// Reassign method, reassign the i-th element with new value
func (mh *MaxHeap) Reassign(i int, value int) {
	mh.heap[i] = value
	heapInsert(mh.heap, i)        // check go up
	heapify(mh.heap, i, mh.count) // check go down
	// if go up, go down will not happen. because upper level always >= lower level
}

func main() {
	rand.Seed(time.Now().UnixNano())
	mh := NewMaxHeap()
	t1 := time.Now()
	for i := 0; i < 5; i++ {
		mh.Push(rand.Intn(20))
	}
	fmt.Println(mh)
	// for mh.count > 0 {
	// 	mh.Pop()
	// }
	mh.Reassign(3, 9)
	fmt.Println(time.Now().Sub(t1))
	fmt.Println(mh)
}
