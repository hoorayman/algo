package main

import (
	"container/list"
	"fmt"
)

type node struct {
	val   int
	left  *node
	right *node
}

func level(np *node) { // level traversal
	queue := list.New()
	if np != nil {
		queue.PushBack(np) // init queue
	}
	for queue.Len() > 0 {
		e := queue.Front()
		np := e.Value.(*node)
		if np.left != nil {
			queue.PushBack(np.left)
		}
		if np.right != nil {
			queue.PushBack(np.right)
		}
		fmt.Println(np.val)
		queue.Remove(e)
	}
}

func maxWidth(np *node) int {
	queue := list.New()
	nodeLevel := make(map[*node]int) // the key is that use a map to record node level
	if np != nil {
		queue.PushBack(np) // init queue
		nodeLevel[np] = 1  // init node level
	}
	levelWidth := make(map[int]int)
	for queue.Len() > 0 {
		e := queue.Front()
		np := e.Value.(*node)
		level := nodeLevel[np]
		if np.left != nil {
			queue.PushBack(np.left)
			nodeLevel[np.left] = level + 1
		}
		if np.right != nil {
			queue.PushBack(np.right)
			nodeLevel[np.right] = level + 1
		}
		levelWidth[level]++
		queue.Remove(e)
	}
	maxWidth := 0
	for _, v := range levelWidth {
		if v > maxWidth {
			maxWidth = v
		}
	}
	return maxWidth
}

func main() {
	nl := []node{node{1, nil, nil}, node{2, nil, nil}, node{3, nil, nil}, node{4, nil, nil}}
	nl[0].left = &nl[1]
	nl[0].right = &nl[2]
	nl[2].right = &nl[3]
	fmt.Println("Level print bst:")
	level(&nl[0])
	fmt.Println("Max width of bst is:", maxWidth(&nl[0]))
}
