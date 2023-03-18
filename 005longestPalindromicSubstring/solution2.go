package main

func longestPalindrome(s string) string {
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		l1, r1 := expandAroundCenter(s, i, i)   // 回文串长度为奇数
		l2, r2 := expandAroundCenter(s, i, i+1) // 回文串长度为偶数
		len1, len2 := r1-l1-1, r2-l2-1
		maxLen := max(len1, len2)
		if maxLen > end-start {
			left, right := 0, 0
			if maxLen == len1 {
				left, right = l1, r1
			} else {
				left, right = l2, r2
			}
			if left < 0 { // 如果不满足中心扩展算法的条件1
				left = 0
				right--
			} else if right > len(s)-1 { // 如果不满足中心扩展算法的条件2
				left++
				right = len(s) - 1
			} else if s[left] != s[right] { // 如果不满足中心扩展算法的条件3
				left++
				right--
			}
			start, end = left, right+1
		}
	}
	return s[start:end]
}

// 中心扩展算法，算法复杂度O(N^2)
func expandAroundCenter(s string, left int, right int) (int, int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return left, right
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
