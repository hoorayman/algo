package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Union definition
type UnionFind struct {
	id    []int
	rank  []int
	count int
}

// New constructor
func New(n int) UnionFind {
	id := make([]int, n)
	rank := make([]int, n)
	// each element is pointed to itself by default
	for i := range id {
		id[i] = i
	}
	return UnionFind{id, rank, n}
}

// Find method
func (u *UnionFind) Find(i int) int {
	for i != u.id[i] {
		i = u.id[i]
	}
	return i
}

// IsConnected method
func (u *UnionFind) IsConnected(a, b int) bool {
	return u.Find(a) == u.Find(b)
}

// Union method
func (u *UnionFind) Union(a, b int) {
	aRoot, bRoot := u.Find(a), u.Find(b)
	if aRoot == bRoot {
		return
	}
	if u.rank[aRoot] < u.rank[bRoot] {
		u.id[aRoot] = bRoot
	} else if u.rank[aRoot] > u.rank[bRoot] {
		u.id[bRoot] = aRoot
	} else { // u.rank[aRoot] == u.rank[bRoot]
		u.id[aRoot] = bRoot
		u.rank[bRoot]++
	}
}

func main() {
	u := New(100000)
	t1 := time.Now()
	for i := 0; i < 100000; i++ {
		u.Union(rand.Int()%100000, rand.Int()%100000)
	}
	for i := 0; i < 100000; i++ {
		u.IsConnected(rand.Int()%100000, rand.Int()%100000)
	}
	fmt.Println(time.Now().Sub(t1))
}
