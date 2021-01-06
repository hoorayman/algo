package main

import "fmt"

// every nodeV1 stand for a letter, only root nodeV1 stand for just root but not a letter
type nodeV1 struct {
	children [26]*nodeV1 // 26 lowercase letter
	pass     int         // number of times passed this node
	end      int         // number of times this node is end
}

// TrieV1 define
type TrieV1 struct {
	root *nodeV1
}

// NewTrieV1 constructor
func NewTrieV1() TrieV1 {
	return TrieV1{&nodeV1{}}
}

// Insert a key word
func (trie TrieV1) Insert(key string) {
	if len(key) == 0 {
		return
	}
	np := trie.root // start from root node, np is node pointer
	for _, char := range key {
		i := char - 'a' // you can take index i as path to letter
		if np.children[i] == nil {
			np.children[i] = new(nodeV1)
		}
		np.pass++
		np = np.children[i]
	}
	np.pass++ // the last node pass++
	np.end++  // the last node end++
}

// Search a key word
func (trie TrieV1) Search(key string) int { // return key word end times
	if len(key) == 0 {
		return 0
	}
	np := trie.root // start from root node, np is node pointer
	for _, char := range key {
		i := char - 'a' // you can take index i as path to letter
		if np.children[i] == nil {
			return 0
		}
		np = np.children[i]
	}
	return np.end
}

func main() {
	trie := NewTrieV1()
	trie.Insert("hello")
	trie.Insert("hey")
	trie.Insert("hey")
	fmt.Println(trie.Search("hey"), trie.Search("hello"), trie.Search("ok"))
}
