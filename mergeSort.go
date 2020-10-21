package main

import (
	"fmt"
	"math/rand"
	"time"
)

func merge(s1, s2 []int) {
	if len(s1) == 0 || len(s2) == 0 {
		return
	}
	if s1[len(s1)-1] <= s2[0] {
		return
	}
	t := make([]int, len(s1)+len(s2))
	i, j, x := 0, 0, 0
	for ; i < len(s1) && j < len(s2); x++ {
		if s1[i] > s2[j] {
			t[x] = s2[j]
			j++
		} else {
			t[x] = s1[i]
			i++
		}
	}
	for ; i < len(s1); x++ {
		t[x] = s1[i]
		i++
	}
	for ; j < len(s2); x++ {
		t[x] = s2[j]
		j++
	}
	copy(s1, t[:len(s1)])
	copy(s2, t[len(s1):])
}

func mergeSort(s []int) []int {
	d := len(s) / 2
	if d > 0 {
		merge(mergeSort(s[:d]), mergeSort(s[d:]))
	}
	return s
}

func main() {
	s := make([]int, 100000)
	for i := range s {
		s[i] = rand.Int()
	}
	t1 := time.Now()
	mergeSort(s)
	fmt.Println(time.Now().Sub(t1))
}
