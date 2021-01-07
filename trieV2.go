package main

import "fmt"

// every nodeV2 stand for a letter, only root nodeV2 stand for just root but not a letter
type nodeV2 struct {
	children map[int]*nodeV2 // any letter
	pass     int             // number of times passed this node
	end      int             // number of times this node is the end
}

// TrieV2 define
type TrieV2 struct {
	root *nodeV2
}

// NewTrieV2 constructor
func NewTrieV2() TrieV2 {
	return TrieV2{&nodeV2{make(map[int]*nodeV2), 0, 0}}
}

// Insert a key word to the trie/prefix tree
func (trie TrieV2) Insert(key string) {
	if len(key) == 0 {
		return
	}
	np := trie.root // start from root node, np is node pointer
	for _, char := range key {
		i := int(char) // you can take index i as path to letter
		if np.children[i] == nil {
			np.children[i] = &nodeV2{make(map[int]*nodeV2), 0, 0}
		}
		np.pass++
		np = np.children[i]
	}
	np.pass++ // the last node pass++
	np.end++  // the last node end++
}

// Search a key word, return how many times the key word is inserted
func (trie TrieV2) Search(key string) int { // return key word end times
	if len(key) == 0 {
		return 0
	}
	np := trie.root // start from root node, np is node pointer
	for _, char := range key {
		i := int(char) // you can take index i as path to letter
		if np.children[i] == nil {
			return 0
		}
		np = np.children[i]
	}
	return np.end
}

// PrefixNumber return how many words begin with the prefix argument
func (trie TrieV2) PrefixNumber(prefix string) int {
	if len(prefix) == 0 {
		return 0
	}
	np := trie.root // start from root node, np is node pointer
	for _, char := range prefix {
		i := int(char) // you can take index i as path to letter
		if np.children[i] == nil {
			return 0
		}
		np = np.children[i]
	}
	return np.pass
}

// Delete a key word from the trie/prefix tree just once
func (trie TrieV2) Delete(key string) {
	if trie.Search(key) > 0 {
		np := trie.root // start from root node, np is node pointer
		for _, char := range key {
			np.pass--
			i := int(char)
			if np.children[i].pass-1 == 0 {
				np.children[i] = nil // if np.children[i].pass will be 0, drop it
				return
			}
			np = np.children[i]
		}
		np.pass-- // the last node pass--
		np.end--  // the last node end--
	}
}

func main() {
	trie := NewTrieV2()
	trie.Insert("hello")
	trie.Insert("hey")
	trie.Insert("hey")
	trie.Delete("hey")
	trie.Delete("hey")
	fmt.Println(trie.Search("hey"), trie.Search("hello"), trie.Search("ok"),
		trie.PrefixNumber("he"), trie.PrefixNumber("y"))
}
