/*
无向图实际上就是双向的有向图，所以实现了有向图就能实现无向图
*/
package main

import (
	"container/list"
	"fmt"
)

// 图结构的描述。就是点集+边集
type graph struct {
	nodes map[int]*node      // 即0号对应的点，1号对应的点，以此类推
	edges map[*edge]struct{} // golang实现集合的trick
}

// 点结构的描述，邻接表法
type node struct {
	value int
	in    int     // 入度，即有几条边直接指向本node。比如同一个node有多条边都指向本node，也算多个入度
	out   int     // 出度，即有几条边直接从本node指向别的node
	nexts []*node // 即从本node直接指向的node列表
	edges []*edge // 即从本node直接指向外部的边
}

// 边结构的描述
type edge struct {
	weight int
	from   *node
	to     *node
}

// 接收一个二维数组，[[边的权重int, from点的编号int, to点的编号int], ...]，返回一个graph
func newGraph(s [][3]int) graph {
	g := graph{nodes: make(map[int]*node), edges: make(map[*edge]struct{})}
	for _, v := range s {
		weight := v[0]
		from := v[1]
		to := v[2]

		var fNode, tNode *node
		var fExist, tExist bool
		if fNode, fExist = g.nodes[from]; !fExist { // 判断from节点是否已经存在
			fNode = &node{value: from}
			g.nodes[from] = fNode
		}
		if tNode, tExist = g.nodes[to]; !tExist { // 判断to节点是否已经存在
			tNode = &node{value: to}
			g.nodes[to] = tNode
		}
		// 下面语句不判断from节点或to节点存不存在，以使得即使相同from和to也可以加上多条边
		e := &edge{weight: weight, from: fNode, to: tNode}
		g.edges[e] = struct{}{}
		fNode.out++                              // 不管from节点、to节点存不存在，from节点的出度都++
		fNode.nexts = append(fNode.nexts, tNode) // 不管from节点、to节点存不存在，from节点的nexts都更新
		fNode.edges = append(fNode.edges, e)     // 不管from节点、to节点存不存在，from节点的edges都更新
		tNode.in++                               // 不管from节点、to节点存不存在，to节点的入度都++
	}
	return g
}

// 宽度优先遍历
func (g graph) level(i int) {
	n, ok := g.nodes[i]
	if !ok {
		return
	}
	queue := list.New()
	set := make(map[*node]struct{}) // set是为了避免环的问题。二叉树不用这个是因为二叉树没有环
	queue.PushBack(n)               // 初始化queue
	set[n] = struct{}{}
	for queue.Len() > 0 {
		e := queue.Front()
		n := e.Value.(*node)
		fmt.Println(n.value)
		for _, n := range n.nexts {
			if _, ok := set[n]; !ok {
				set[n] = struct{}{}
				queue.PushBack(n)
			}
		}
		queue.Remove(e)
	}
}

// 深度优先遍历
func (g graph) deep(i int) {
	n, ok := g.nodes[i]
	if !ok {
		return
	}
	set := make(map[*node]struct{}) // set是为了避免环的问题
	set[n] = struct{}{}
	goDeep(n, set)
}

func goDeep(n *node, set map[*node]struct{}) {
	fmt.Println(n.value)
	for _, n := range n.nexts {
		if _, ok := set[n]; !ok {
			set[n] = struct{}{}
			goDeep(n, set)
		}
	}
}

// 拓扑排序。必须是有向无环图，因为如果有环的话就有可能根本就没有入度为0的节点。
// 拓扑排序就是根据有向无环图的先后依赖排出先后顺序，常用于编译器
// 拓扑排序的方法是选出入度是0的节点，然后消除它指向节点的入度，然后再选出入度是0的节点，以此类推直到没有入度是0的节点为止
func (g graph) topologySort() []int {
	inMap := make(map[*node]int) // 存放node对应的入度，将来修改入度就改这个，避免直接改graph
	zeroInQueue := list.New()    // 入度为0的节点队列
	for _, n := range g.nodes {
		inMap[n] = n.in
		if n.in == 0 {
			zeroInQueue.PushBack(n)
		}
	}

	var result []int
	for zeroInQueue.Len() > 0 {
		e := zeroInQueue.Front()
		n := e.Value.(*node)
		result = append(result, n.value)
		for _, v := range n.nexts {
			inMap[v]--
			if inMap[v] == 0 {
				zeroInQueue.PushBack(v)
			}
		}
		zeroInQueue.Remove(e)
	}
	return result
}

func main() {
	g := newGraph([][3]int{{1, 2, 3}, {4, 3, 6}, {4, 6, 5}, {4, 3, 7}})
	fmt.Println("宽度优先遍历：")
	g.level(2)
	fmt.Println("深度优先遍历：")
	g.deep(2)
	fmt.Println("拓扑排序：", g.topologySort())
}
