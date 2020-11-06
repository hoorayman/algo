package main

import (
	"errors"
	"fmt"
	"time"
)

type dllNode struct {
	value int
	prev  *dllNode
	next  *dllNode
}

// DLL double linked list definition
type DLL struct {
	root  *dllNode
	count int
}

// NewDLL constructor
func NewDLL() DLL {
	return DLL{}
}

// Size method
func (dll *DLL) Size() int {
	return dll.count
}

// IsEmpty method
func (dll *DLL) IsEmpty() bool {
	return dll.count == 0
}

// Append method
func (dll *DLL) Append(v int) {
	if dll.root == nil {
		dll.root = &dllNode{value: v}
	} else {
		var np *dllNode
		np = dll.root
		for np.next != nil {
			np = np.next
		}
		np.next = &dllNode{value: v, prev: np}
	}
	dll.count++
}

// Pop remove last node and return the value of the removed node
func (dll *DLL) Pop() (int, error) {
	if dll.count > 1 {
		var np *dllNode
		np = dll.root
		for np.next.next != nil {
			np = np.next
		}
		v := np.next.value
		np.next = nil
		dll.count--
		return v, nil
	} else if dll.count == 1 {
		v := dll.root.value
		dll.root = nil
		dll.count = 0
		return v, nil
	} else { // dll.count < 1
		return 0, errors.New("no node to pop")
	}
}

// Reverse all the node
func (dll *DLL) Reverse() {
	if dll.count > 1 { // current node is the node which is to be moved after the root
		var np *dllNode    // np point to the current node
		np = dll.root.next // the beginning current node pointer is dll.root.next

		for np != nil {
			next := np.next        // back the next node up
			np.prev.next = np.next // this line and the next if line take the current node off the double linked list
			if np.next != nil {
				np.next.prev = np.prev
			}
			fp := dll.root // back the first node up
			dll.root = np  // this line and the next 3 lines insert the current node after the root
			np.prev = nil
			np.next = fp
			fp.prev = np
			np = next
		}
	}
}

// Remove all the node whose value is equal to v
func (dll *DLL) Remove(v int) {
	for np := dll.root; np != nil; np = np.next {
		if np.value == v {
			if np.prev == nil { // first node
				dll.root = np.next
				if np.next != nil {
					np.next.prev = nil
				}
			} else {
				np.prev.next = np.next
				if np.next != nil {
					np.next.prev = np.prev
				}
			}
			dll.count--
		}
	}
}

// Print method
func (dll *DLL) Print() {
	var np *dllNode
	np = dll.root
	for np != nil {
		fmt.Printf("%d ", np.value)
		np = np.next
	}
	fmt.Printf("\n")
}

func main() {
	t1 := time.Now()
	dll := NewDLL()
	for i := 0; i < 10; i++ {
		dll.Append(i)
		dll.Append(i + 1)
	}
	dll.Remove(1)
	dll.Print()
	dll.Reverse()
	dll.Print()
	fmt.Println(dll, time.Now().Sub(t1))
}
