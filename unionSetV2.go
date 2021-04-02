package main

import (
	"container/list"
	"fmt"
)

type element struct {
	value interface{}
}

type unionSet struct {
	nodes   map[interface{}]*element // nodes make a poelementer which stands for a node
	parents map[*element]*element    // stand for node's parent
	size    map[*element]int         // a node can be in size only if it is a head node
}

func newUnionSet(nodes ...interface{}) unionSet {
	result := unionSet{nodes: make(map[interface{}]*element), parents: make(map[*element]*element), size: make(map[*element]int)}
	for _, v := range nodes {
		e := element{v}
		result.nodes[v] = &e
		result.parents[&e] = &e
		result.size[&e] = 1
	}
	return result
}

func (u unionSet) findFather(n *element) *element {
	queue := list.New()
	for n != u.parents[n] {
		queue.PushBack(n)
		n = u.parents[n]
	}
	for queue.Len() > 0 { // optimize height of union set
		e := queue.Front()
		u.parents[e.Value.(*element)] = n
		queue.Remove(e)
	}
	return n
}

func (u unionSet) isSameSet(a, b interface{}) bool {
	na, naOK := u.nodes[a]
	nb, nbOK := u.nodes[b]
	if naOK == false || nbOK == false {
		return false
	}
	return u.findFather(na) == u.findFather(nb)
}

func (u unionSet) union(a, b interface{}) {
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
	u := newUnionSet(1, 2, 3, 4, 5, 6, 7, 8)
	fmt.Println(u.isSameSet(2, 3))
	u.union(2, 3)
	u.union(4, 5)
	u.union(2, 5)
	fmt.Println(u.isSameSet(3, 5))
}
