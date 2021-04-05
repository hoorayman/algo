package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type node struct {
	key   int
	value int
	left  *node
	right *node
}

// BST binary search tree definition
type BST struct {
	root  *node
	count int
}

// New constructor
func New() BST {
	bst := BST{}
	return bst
}

// Size method
func (bst *BST) Size() int {
	return bst.count
}

// IsEmpty method
func (bst *BST) IsEmpty() bool {
	return bst.count == 0
}

func (bst *BST) setkv(p *node, key, value int) *node {
	if p == nil {
		bst.count++
		return &node{key: key, value: value}
	}
	if key == p.key {
		p.value = value
	} else if key > p.key {
		p.right = bst.setkv(p.right, key, value)
	} else {
		p.left = bst.setkv(p.left, key, value)
	}
	return p
}

// Set method
func (bst *BST) Set(key, value int) {
	bst.root = bst.setkv(bst.root, key, value)
}

func (bst *BST) findk(p *node, key int) *node {
	if p == nil || key == p.key {
		return p
	} else if key > p.key {
		return bst.findk(p.right, key)
	} else {
		return bst.findk(p.left, key)
	}
}

// Search method
func (bst *BST) Search(key int) (int, error) {
	p := bst.findk(bst.root, key)
	if p == nil {
		return 0, errors.New("not exist")
	}
	return p.value, nil
}

// Contain method
func (bst *BST) Contain(key int) bool {
	p := bst.findk(bst.root, key)
	if p == nil {
		return false
	}
	return true
}

func (bst *BST) min(p *node) *node {
	for p != nil && p.left != nil {
		p = p.left
	}
	return p
}

// Mink method
func (bst *BST) Mink() (int, error) {
	p := bst.min(bst.root)
	if p == nil {
		return 0, errors.New("not exist")
	}
	return p.key, nil
}

func (bst *BST) max(p *node) *node {
	for p != nil && p.right != nil {
		p = p.right
	}
	return p
}

// Maxk method
func (bst *BST) Maxk() (int, error) {
	p := bst.max(bst.root)
	if p == nil {
		return 0, errors.New("not exist")
	}
	return p.key, nil
}

// findp method: return current node pointer and parent node pointer
func (bst *BST) findp(p *node, key int, pp *node) (*node, *node) {
	if p == nil || key == p.key {
		return p, pp
	} else if key > p.key {
		return bst.findp(p.right, key, p)
	} else {
		return bst.findp(p.left, key, p)
	}
}

func (bst *BST) removeMax(p *node) *node {
	if p == nil {
		return p
	}
	if p.right == nil {
		bst.count--
		return p.left
	}
	p.right = bst.removeMax(p.right)
	return p
}

// Del method
func (bst *BST) Del(key int) {
	p, pp := bst.findp(bst.root, key, nil)
	if p == nil {
		return
	}
	var np *node
	if p.right == nil {
		np = p.left
	} else if p.left == nil {
		np = p.right
	} else {
		np = bst.max(p.left)
		p.left = bst.removeMax(p.left)
		np.left = p.left
		np.right = p.right
		bst.count++
	}
	if pp == nil {
		bst.root = np
	} else if pp.key < p.key {
		pp.right = np
	} else {
		pp.left = np
	}
	bst.count--
}

// Printk method: print every key in BST begin at a node pointer
func (bst *BST) Printk(p *node) {
	if p == nil {
		return
	}
	fmt.Printf("%d ", p.key)
	bst.Printk(p.left)
	bst.Printk(p.right)
}

func main() {
	bst := New()
	t1 := time.Now()
	for i := 0; i < 100000; i++ {
		bst.Set(rand.Int(), i)
	}
	for bst.root != nil {
		bst.Del(bst.root.key)
	}
	fmt.Println(time.Now().Sub(t1))
	bst.Printk(bst.root)
}
