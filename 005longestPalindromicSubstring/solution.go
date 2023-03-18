package main

func longestPalindrome(s string) string {
	array := []byte(s)

	max := 0
	left, right := 0, 0
	for i := range array {
		for j := 0; j < len(array); j++ {
			if isPalindrome(array[j : i+1]) {
				if i+1-j > max {
					max = i + 1 - j
					left, right = j, i+1
				}
				break
			}
		}
	}

	return string(array[left:right])
}

func isPalindrome(s []byte) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 { // 第3个for循环，算法复杂度O(N^3)
		if s[i] != s[j] {
			return false
		}
	}

	return true
}
