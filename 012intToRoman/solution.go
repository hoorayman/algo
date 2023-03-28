package main

func intToRoman(num int) string {
	dict := map[int][]string{
		0: {"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}, // 个位：0-9
		1: {"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}, // 十位：0-9
		2: {"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}, // 百位：0-9
		3: {"", "M", "MM", "MMM"},                                       // 千位：0-3
	}

	result := ""
	for i := 0; num > 0; i++ {
		digit := num % 10
		result = dict[i][digit] + result
		num /= 10
	}

	return result
}
