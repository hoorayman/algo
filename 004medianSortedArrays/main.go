package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := merge(nums1, nums2)
	if (len(nums1)+len(nums2))%2 != 0 {
		return float64(m[(len(nums1)+len(nums2)+1)/2-1])
	}
	return float64((m[(len(nums1)+len(nums2))/2-1] + m[(len(nums1)+len(nums2))/2])) / 2
}

func merge(a, b []int) []int {
	result := []int{}
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			result = append(result, a[i])
			i++
		} else {
			result = append(result, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		result = append(result, a[i])
	}
	for ; j < len(b); j++ {
		result = append(result, b[j])
	}
	return result
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{}, []int{1, 2}))
}
