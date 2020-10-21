package main

import (
	"fmt"
	"math/rand"
	"time"
)

func partition(s []int) ([]int, []int) {
	length := len(s)
	if length < 2 {
		return nil, s
	}
	r := rand.Intn(length)
	s[0], s[r] = s[r], s[0]

	i, j := 1, length-1
	for {
		for ; i < length && s[i] < s[0]; i++ {
		}
		for ; i < j && s[j] > s[0]; j-- {
		}
		if i < length && i < j {
			s[i], s[j] = s[j], s[i]
			i++
			j--
		} else {
			break
		}
	}
	s[0], s[i-1] = s[i-1], s[0]
	if i < length {
		return s[:i], s[i:]
	}
	return s[:i-1], s[i-1:]
}

func quickSort(s []int) {
	left, right := partition(s)
	if len(left) > 1 {
		quickSort(left)
	}
	if len(right) > 1 {
		quickSort(right)
	}
}

func main() {
	s := make([]int, 100000)
	for i := range s {
		s[i] = rand.Int()
	}
	t1 := time.Now()
	quickSort(s)
	fmt.Println(time.Now().Sub(t1))
}
