package main

import "fmt"

/*
可以不断地将原数的末尾数字弹出，然后添加到反转后的数字的开头，
直到原数小于等于反转后的数字，这样可以反转一半的数字。
然后将原数与反转后的数字进行比较，如果它们相等，则原数是回文数。

在将原数反转的过程中，如果原数有奇数位，
那么在反转过程中，reversed 的位数会多一位，
这个多出来的一位就是原数的中间数字。
例如，对于数字 12321，反转后得到 1232，此时 x=12，reversed=123，
我们需要将 reversed 除以 10 去掉多余的一位，
然后判断 x 是否等于 reversed 或者等于 reversed 除以 10（即去掉中间数字的 reversed）。
例如：
初始：x = 12321, reversed = 0
1. 取出 x 的末位数字 1，reversed 变成 1
    x = 1232, reversed = 1
2. 取出 x 的末位数字 2，reversed 变成 12
    x = 123, reversed = 12
3. 取出 x 的末位数字 3，reversed 变成 123
    x = 12, reversed = 123
此时 x <= reversed，退出循环
最后判断 x 是否等于 reversed 或者等于 reversed 除以 10
*/
func isPalindrome(x int) bool {
	if x < 0 || x != 0 && x%10 == 0 { // x<0时，或者，x>0且是10的倍数时
		return false
	}

	result := 0
	for x > result {
		result = result*10 + x%10
		x /= 10
	}

	return x == result || x == result/10
}

func main() {
	fmt.Println(isPalindrome(12321))
}
