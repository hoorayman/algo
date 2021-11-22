package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	left, right := nums1, nums2
	if len(nums1) > len(nums2) {
		left, right = nums2, nums1
	}
	i := findIndexOfLeft(left, right, 0, len(left)-1)
	sumLen := len(left) + len(right)
	if sumLen%2 == 0 {
		if i == -1 {
			return float64(right[sumLen/2-1]+right[sumLen/2]) / 2
		}
		x, y := 0, 0
		j := sumLen/2 + 1 - (i + 1) - 1
		if left[i] > right[j] {
			y = left[i]
			if i-1 >= 0 {
				if left[i-1] > right[j] {
					x = left[i-1]
				} else {
					x = right[j]
				}
			} else {
				x = right[j]
			}
		} else {
			y = right[j]
			if j-1 >= 0 {
				if right[j-1] > left[i] {
					x = right[j-1]
				} else {
					x = left[i]
				}
			} else {
				x = left[i]
			}
		}
		return float64(x+y) / 2
	} else {
		if i == -1 {
			return float64(right[sumLen/2])
		}
		j := sumLen/2 + 1 - (i + 1) - 1
		if left[i] > right[j] {
			return float64(left[i])
		}
		return float64(right[j])
	}
}

func findIndexOfLeft(left, right []int, i, j int) int {
	if i <= j {
		leftMid := i + (j-i)/2
		rightMid := (len(left)+len(right))/2 + 1 - (i + (j-i)/2 + 1) - 1

		if leftMid+1 >= len(left) {
			if rightMid+1 >= len(right) || left[leftMid] <= right[rightMid+1] {
				return leftMid
			}
			j = leftMid - 1
			return findIndexOfLeft(left, right, i, j)
		}
		if rightMid+1 >= len(right) {
			if right[rightMid] <= left[leftMid+1] {
				return leftMid
			}
			i = leftMid + 1
			return findIndexOfLeft(left, right, i, j)
		}
		if left[leftMid] <= right[rightMid+1] && right[rightMid] <= left[leftMid+1] {
			return leftMid
		}
		if left[leftMid] > right[rightMid+1] {
			j = leftMid - 1
			return findIndexOfLeft(left, right, i, j)
		}
		if right[rightMid] > left[leftMid+1] {
			i = leftMid + 1
			return findIndexOfLeft(left, right, i, j)
		}
	}
	return -1
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{4, 6, 7}, []int{1, 2, 3}))
}
