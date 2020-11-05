// 链表的特点是：node一旦创建node的地址当然不变，变的只是node中的指向别的node的指针值
package main

import (
	"errors"
	"fmt"
	"time"
)

type sllNode struct {
	value int
	next  *sllNode
}

// SLL single linked list definition
type SLL struct {
	root  *sllNode
	count int
}

// NewSLL constructor
func NewSLL() SLL {
	return SLL{}
}

// Size method
func (sll *SLL) Size() int {
	return sll.count
}

// IsEmpty method
func (sll *SLL) IsEmpty() bool {
	return sll.count == 0
}

// Append method
func (sll *SLL) Append(v int) {
	if sll.root == nil {
		sll.root = &sllNode{value: v}
	} else {
		var np *sllNode
		np = sll.root
		for np.next != nil {
			np = np.next
		}
		np.next = &sllNode{value: v}
	}
	sll.count++
}

// Pop remove last node and return the value of the removed node
func (sll *SLL) Pop() (int, error) {
	if sll.root == nil {
		return 0, errors.New("no node to pop")
	} else if sll.count == 1 {
		v := sll.root.value
		sll.root = nil
		sll.count = 0
		return v, nil
	} else {
		var np *sllNode
		np = sll.root
		for np.next.next != nil {
			np = np.next
		}
		v := np.next.value
		np.next = nil
		sll.count--
		return v, nil
	}
}

// Reverse all the node
func (sll *SLL) Reverse() {
	if sll.count > 1 { // current node is the node which is to be moved after the root
		var np *sllNode    // np point to the current node
		lnp := sll.root    // lnp is a pointer to the node on the left of current node
		np = sll.root.next // the beginning current node pointer is sll.root.next

		for np != nil {
			p1 := np.next // np.next is to be replaced, so back it up
			lnp.next = p1
			p2 := sll.root // sll.root is to be replaced, so back it up
			sll.root = np  // root point to the current node
			sll.root.next = p2
			np = p1
		}
	}
}

// Remove all the node whose value is equal to v
func (sll *SLL) Remove(v int) {
	var lnp *sllNode
	for np := sll.root; np != nil; np = np.next {
		if np.value == v {
			if lnp == nil {
				sll.root = np.next
			} else {
				lnp.next = np.next
			}
			sll.count--
		} else {
			lnp = np
		}
	}
}

// Print method
func (sll *SLL) Print() {
	var np *sllNode
	np = sll.root
	for np != nil {
		fmt.Printf("%d ", np.value)
		np = np.next
	}
	fmt.Printf("\n")
}

func main() {
	t1 := time.Now()
	sll := NewSLL()
	for i := 0; i < 10; i++ {
		sll.Append(i)
		sll.Append(i + 1)
	}
	sll.Print()
	sll.Reverse()
	sll.Print()
	sll.Remove(9)
	sll.Print()
	fmt.Println(sll, time.Now().Sub(t1))
}
