/*
  当前字符的不重复最大长度取决于：
  1. 以前一个字符结尾的不重复最大长度+1（+1是因为有当前字符）
  2. 当前字符与上一次出现当前字符的长度
*/
package main

func lengthOfLongestSubstring(s string) int {
	position := make(map[byte]int)

	max, maxDistanceEndWithPreIndex := 0, 0
	for i, c := range []byte(s) {
		preIndex := -1
		if x, ok := position[c]; ok {
			preIndex = x
		}

		distance := i - preIndex
		if distance < maxDistanceEndWithPreIndex+1 {
			maxDistanceEndWithPreIndex = distance
		} else {
			maxDistanceEndWithPreIndex = maxDistanceEndWithPreIndex + 1
		}
		if maxDistanceEndWithPreIndex > max {
			max = maxDistanceEndWithPreIndex
		}
		position[c] = i
	}

	return max
}
