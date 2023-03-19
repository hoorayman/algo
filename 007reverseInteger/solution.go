package main

func reverse(x int) int {
	var u32 uint32
	minInt := int32(^(^u32 >> 1))
	positive := false
	n := int32(x)
	if n > 0 { // 因为负数能hold住所有正数，所以都转换成负数，仅判断负数是否溢出即可
		positive = true
		n = -n
	}

	var result int32
	for n != 0 {
		if result < minInt/10 || result == minInt/10 && result*10+n%10 >= 0 { // overflow
			return 0
		}
		result = result*10 + n%10
		n /= 10
	}
	if positive {
		result = -result
	}

	return int(result)
}
