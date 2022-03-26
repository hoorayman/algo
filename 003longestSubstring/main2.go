/*
  sub问题的解法用以xxx结尾的思路，如果用以xxx开头的思路必然是O(n^2)的复杂度了
*/

package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	position := make(map[byte]int)
	array := []byte(s)

	result, maxLenOfSubstringEndWithPreChar := 0, 0
	for index, char := range array {
		lastIndexOfChar := -1
		if i, ok := position[char]; ok {
			lastIndexOfChar = i
		}

		maxLenOfSubstringEndWithCurrentChar := min(index-lastIndexOfChar, maxLenOfSubstringEndWithPreChar+1)
		if maxLenOfSubstringEndWithCurrentChar > result {
			result = maxLenOfSubstringEndWithCurrentChar
		}
		maxLenOfSubstringEndWithPreChar = maxLenOfSubstringEndWithCurrentChar
		position[char] = index
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}
