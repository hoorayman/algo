package main

import "fmt"

func reverse(x int) int {
	var u32 uint32
	minInt32 := int32(^(^u32 >> 1))
	positive := int32(x)&minInt32 == 0
	if positive { // change positive to negtive, so we have to only check negtive overflow
		x = -x // negtive can hold all positive
	}

	var result int32
	y := int32(x)
	for y != 0 {
		if result < minInt32/10 || result == minInt32/10 && (result*10+y%10) > 0 { // overflow
			return 0
		}
		result = result*10 + y%10
		y /= 10
	}
	if positive {
		result = -result
	}

	return int(result)
}

func main() {
	fmt.Println(reverse(1463847412))
}
