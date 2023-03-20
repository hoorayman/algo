package main

import "fmt"

func myAtoi(s string) int {
	array := []byte(s)
	start, end := 0, 0
	for i := 0; i < len(array) && array[i] == ' '; i++ { // 略过空格
		start++
	}
	positive := true
	for i := start; i < len(array) && (array[i] == '+' || array[i] == '-'); i++ { // 略过符号
		if array[i] == '-' {
			positive = false
			start++
			break
		} else if array[i] == '+' {
			start++
			break
		}
	}
	end = start
	for i := start; i < len(array) && array[i] >= '0' && array[i] <= '9'; i++ { // 略过数字
		end++
	}
	result := process(positive, array[start:end])

	return int(result)
}

func process(positive bool, s []byte) int32 {
	var u32 uint32
	minInt := int32(^(^u32 >> 1))
	checker := minInt / 10
	var result int32
	overflow := false
	for i := 0; i < len(s); i++ {
		if result < checker || result == checker && result*10-int32(s[i]-'0') >= 0 {
			overflow = true
			break
		}
		result = result*10 - int32(s[i]-'0') // 相当于result = result*10 + -int32(s[i]-'0')， 为的是转换成负数， 仅需判断负数的overflow即可
	}
	if overflow {
		if positive {
			return minInt - 1
		} else {
			return minInt
		}
	} else {
		if positive && result == minInt {
			return minInt - 1
		}
	}
	if positive {
		result = -result
	}

	return result
}

func main() {
	fmt.Println(myAtoi("2147483648"))
}
