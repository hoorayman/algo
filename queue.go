package main

import (
	"errors"
	"fmt"
	"time"
)

// Queue (FIFO) implemented by array
type Queue struct {
	data          []int
	capacity      int
	lastPopIndex  int // the index of last element which has been poped
	lastPushIndex int // the index of last element which has been pushed
	// current capacity of queue is lastPushIndex-lastPopIndex
}

// NewQueue constructor
func NewQueue(size int) Queue {
	if size < 1 { // because we will mod q.capacity, so size must be >0
		panic("invalid size")
	}
	data := make([]int, size)
	return Queue{data: data, capacity: size, lastPopIndex: -1, lastPushIndex: -1}
}

// Push append value to the end of queue
func (q *Queue) Push(v int) error {
	if q.lastPushIndex-q.lastPopIndex < q.capacity {
		q.lastPushIndex++
		q.data[q.lastPushIndex%q.capacity] = v
		return nil
	}
	return errors.New("queue is full")
}

// Pop pop first value of queue
/*
if do not take care of the lastPopIndex and lastPushIndex overflow,the pop method can be:
func (q *Queue) Pop() (int, error) {
	if q.lastPopIndex < q.lastPushIndex {
		q.lastPopIndex++
		return q.data[q.lastPopIndex%q.capacity], nil
	}
	return 0, errors.New("queue is empty")
}
we must take care of lastPopIndex and lastPushIndex overflow,
so the pop method will be as the following
*/
func (q *Queue) Pop() (int, error) {
	distance := q.lastPushIndex - q.lastPopIndex // remember the distance
	if distance > 0 {
		q.lastPopIndex++
		q.lastPopIndex = q.lastPopIndex % q.capacity
		q.lastPushIndex = q.lastPopIndex + distance - 1 // -1 because of q.lastPopIndex++
		return q.data[q.lastPopIndex], nil
	}
	return 0, errors.New("queue is empty")
}

// Print show the underlying data of queue
func (q *Queue) Print() {
	fmt.Println(q.data)
}

func main() {
	t1 := time.Now()
	q := NewQueue(5)
	for i := 0; i < 10; i++ {
		q.Push(i)
	}
	q.Pop()
	q.Pop()
	q.Pop()
	q.Pop()
	for i := 10; i < 100; i++ {
		q.Push(i)
	}
	q.Pop()
	q.Push(14)
	q.Push(15)
	q.Pop()
	q.Push(14)
	for i := 10; i < 100; i++ {
		q.Pop()
		q.Push(i)
	}
	fmt.Println(q, time.Now().Sub(t1))
}
