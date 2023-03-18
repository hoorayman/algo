package main

import "fmt"

func longestPalindrome(s string) string {
	transToBytes := []byte(s)
	charPosition := map[byte][]int{}

	for i := 0; i < len(transToBytes); i++ {
		charPosition[transToBytes[i]] = append(charPosition[transToBytes[i]], i)
	}

	maxSubstrLen := 0
	beginIndex, endIndex := 0, 0
	for _, indexs := range charPosition {
		for span := len(indexs); span > 0; span-- {
			for i := 0; i+span <= len(indexs); i++ {
				if indexs[i+span-1]-indexs[i]+1 > maxSubstrLen &&
					isPalindrome(transToBytes[indexs[i]:indexs[i+span-1]+1]) {
					maxSubstrLen = indexs[i+span-1] - indexs[i] + 1
					beginIndex, endIndex = indexs[i], indexs[i+span-1]
				}
			}
		}
	}
	return string(transToBytes[beginIndex : endIndex+1])
}

func isPalindrome(s []byte) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(longestPalindrome("aaabaa"))
}
