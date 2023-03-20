package main

/*
双指针法

area的大小取决于2个方面：
1、较小的height
2、较长的底边

可以用两层for暴力解法，但是有很多不需要判断的条件，所以需要简化
例如：[2, 10, 8]或[2, 1, 8]
暴力解为：2:10， 2:8， 10:8和2:1， 2:8， 1:8
可见两端的2:8，最大area为2*2=4，中间比2大的10或者比2小的1不起决定作用，所以根本不用判断
另外，因为起决定作用的是较小的数，也就是2，所以左端++去寻找下一个较小的值（起决定作用的值）；
同理，如果右边较小，则右端--去寻找下一个较小的值（起决定作用的值）
所以复杂度减少到 O(n)
*/
func maxArea(height []int) int {
	max := 0
	i, j := 0, len(height)-1
	for i < j {
		area := 0
		if height[i] < height[j] {
			area = height[i] * (j - i)
			i++
		} else {
			area = height[j] * (j - i)
			j--
		}
		if area > max {
			max = area
		}
	}

	return max
}
