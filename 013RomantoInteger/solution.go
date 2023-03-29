package main

func romanToInt(s string) int {
	dict := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	result := 0
	bytes := []byte(s)
	for i := 0; i < len(bytes); i++ {
		if i+1 < len(bytes) && dict[bytes[i]] < dict[bytes[i+1]] { // 后边字符比前面的小就减，否则加
			result -= dict[bytes[i]]
		} else {
			result += dict[bytes[i]]
		}
	}

	return result
}
