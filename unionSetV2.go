package main

import (
	"container/list"
	"fmt"
)

type unionSet struct {
	nodes   map[int]*int  // nodes make a pointer which stands for a node
	parents map[*int]*int // stand for node's parent
	size    map[*int]int  // a node can be in size only if it is a head node
}

func NewUnionSet(nodes ...int) unionSet {
	result := unionSet{nodes: make(map[int]*int), parents: make(map[*int]*int), size: make(map[*int]int)}
	for _, v := range nodes {
		n := v
		result.nodes[n] = &n
		result.parents[&n] = &n
		result.size[&n] = 1
	}
	return result
}

func (u unionSet) findFather(n *int) *int {
	queue := list.New()
	for n != u.parents[n] {
		queue.PushBack(n)
		n = u.parents[n]
	}
	for queue.Len() > 0 { // optimize height of union set
		e := queue.Front()
		u.parents[e.Value.(*int)] = n
		queue.Remove(e)
	}
	return n
}

func (u unionSet) isSameSet(a, b int) bool {
	na, naOK := u.nodes[a]
	nb, nbOK := u.nodes[b]
	if naOK == false || nbOK == false {
		return false
	}
	return u.findFather(na) == u.findFather(nb)
}

func (u unionSet) union(a, b int) {
	na, naOK := u.nodes[a]
	nb, nbOK := u.nodes[b]
	if naOK == false || nbOK == false {
		return
	}
	aHead := u.findFather(na)
	bHead := u.findFather(nb)
	if aHead != bHead {
		sizeA := u.size[aHead]
		sizeB := u.size[bHead]
		if sizeA > sizeB { // merge small size to big size union set to minimize the amount of node to be optimized in findFather method
			u.parents[bHead] = aHead
			u.size[aHead] = sizeA + sizeB
			delete(u.size, bHead)
		} else {
			u.parents[aHead] = bHead
			u.size[bHead] = sizeA + sizeB
			delete(u.size, aHead)
		}
	}
}

func main() {
	u := NewUnionSet(1, 2, 3, 4, 5, 6, 7, 8)
	fmt.Println(u.isSameSet(2, 3))
	u.union(2, 3)
	u.union(4, 5)
	u.union(2, 5)
	fmt.Println(u.isSameSet(2, 3))
}
