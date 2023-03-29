package main

func longestCommonPrefix(strs []string) string {
	first := []byte(strs[0]) // 拿第一个做比较即可

	max := len(first)
	for _, s := range strs[1:] {
		bytes := []byte(s)
		j := 0
		for ; j < len(bytes) && j < len(first) && bytes[j] == first[j]; j++ {
		}
		if j < max {
			max = j
		}
	}

	return string(first[:max])
}
