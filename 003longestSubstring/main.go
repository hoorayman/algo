package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	b := []byte(s)
	result := 0
	for i := 0; i < len(b); i++ {
		m := make(map[byte]int)
		subLen := 0
		for j := i; j < len(b); j++ {
			if index, ok := m[b[j]]; ok {
				i = index
				break
			} else {
				m[b[j]] = j
				subLen++
			}
		}
		if subLen > result {
			result = subLen
		}
		if subLen >= len(b)-i {
			break
		}
	}
	return result
}

func main() {
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}
