package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// Graph definition
type Graph struct {
	vertex   int // vertex number
	edge     int // edge number
	directed bool
	neighbor [][]int // neighbor table
}

// New constructor
func New(n int, d bool) Graph {
	return Graph{vertex: n, directed: d, neighbor: make([][]int, n)}
}

// V method, return vertex number
func (g *Graph) V() int {
	return g.vertex
}

// E method, return edge number
func (g *Graph) E() int {
	return g.edge
}

func (g *Graph) hasEdge(a, b int) bool {
	sort.Ints(g.neighbor[a])
	i := sort.SearchInts(g.neighbor[a], b)
	if i < len(g.neighbor[a]) && g.neighbor[a][i] == b {
		return true
	}
	return false
}

// AddEdge method, add an edge from a to b (and vise versa if the graph is not directed), disallow self-loop edge and parallel edge
func (g *Graph) AddEdge(a, b int) {
	if a < 0 || a >= g.vertex || b < 0 || b >= g.vertex || a == b {
		return
	}
	if g.hasEdge(a, b) {
		return
	}
	g.neighbor[a] = append(g.neighbor[a], b)
	if !g.directed {
		g.neighbor[b] = append(g.neighbor[b], a)
	}
	g.edge++
}

// PrintNeighbor method
func (g *Graph) PrintNeighbor(n int) {
	if n < 0 || n >= g.vertex {
		return
	}
	fmt.Printf("%d: %v\n", n, g.neighbor[n])
}

func (g *Graph) dfs(i int, v, r []int) []int {
	v[i] = 1
	r = append(r, i)
	for _, e := range g.neighbor[i] {
		if v[e] == 0 {
			r = g.dfs(e, v, r)
		}
	}
	return r
}

// DFS method, depth first search traversal
func (g *Graph) DFS(i int) []int {
	if i < 0 || i >= g.vertex {
		return nil
	}
	visited := make([]int, g.vertex)
	return g.dfs(i, visited, nil)
}

func (g *Graph) component(i int, v []int) {
	v[i] = 1
	for _, e := range g.neighbor[i] {
		if v[e] == 0 {
			g.component(e, v)
		}
	}
}

// Component method, return how many components does the graph has
func (g *Graph) Component() int {
	visited := make([]int, g.vertex)
	count := 0
	for i := 0; i < g.vertex; i++ {
		if visited[i] == 0 {
			g.component(i, visited)
			count++
		}
	}
	return count
}

func (g *Graph) dfsFindPath(i, d int, v, r []int) []int {
	v[i] = 1
	r = append(r, i)
	if i == d {
		return r // found it
	}
	for _, e := range g.neighbor[i] {
		if v[e] == 0 {
			r = g.dfsFindPath(e, d, v, r)
		}
	}
	if r[len(r)-1] != d {
		r = r[:len(r)-1] // delete last level if not found
	}
	return r
}

// DFSFindPath method, find a path using depth first search from a to b
func (g *Graph) DFSFindPath(a, b int) []int {
	if a < 0 || a >= g.vertex || b < 0 || b >= g.vertex {
		return nil
	}
	visited := make([]int, g.vertex)
	return g.dfsFindPath(a, b, visited, nil)
}

func (g *Graph) visited(i, d int, v [][]int) bool {
	for _, e := range v[i] {
		if e == d {
			return true
		}
	}
	return false
}

func (g *Graph) dfstrace(i int, v [][]int, r []int) []int {
	r = append(r, i)
	for _, e := range g.neighbor[i] {
		if g.visited(i, e, v) == false {
			v[i] = append(v[i], e)
			r = g.dfstrace(e, v, r)
		}
	}
	return r
}

// DFSTrace method, get trace using depth first search traversal
func (g *Graph) DFSTrace(i int) []int {
	if i < 0 || i >= g.vertex {
		return nil
	}
	visited := make([][]int, g.vertex)
	return g.dfstrace(i, visited, nil)
}

func main() {
	n := 100000
	g := New(n, false)
	t1 := time.Now()
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < n; i++ {
		g.AddEdge(rand.Intn(n), rand.Intn(n))
	}
	// fmt.Println(g.DFS(8))
	// fmt.Println(g.DFSTrace(0))
	// fmt.Println(g.Component())
	// fmt.Println(g.DFSFindPath(0, 8))
	fmt.Println(time.Now().Sub(t1))
}
