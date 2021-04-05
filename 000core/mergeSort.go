package main

import (
	"fmt"
	"math/rand"
	"time"
)

func mergeSort(a []int) []int {
	if len(a) > 1 {
		merge(mergeSort(a[:len(a)/2]), mergeSort(a[len(a)/2:]))
	}
	return a
}

func merge(l []int, r []int) {
	tmp := make([]int, len(l)+len(r))
	ti, li, ri := 0, 0, 0
	for ; li < len(l) && ri < len(r); ti++ {
		if l[li] < r[ri] {
			tmp[ti] = l[li]
			li++
		} else {
			tmp[ti] = r[ri]
			ri++
		}
	}
	for ; li < len(l); li++ { // if upper for end caused by r run out
		tmp[ti] = l[li]
		ti++
	}
	for ; ri < len(r); ri++ { // if upper for end caused by l run out
		tmp[ti] = r[ri]
		ti++
	}
	copy(l, tmp[:len(l)])
	copy(r, tmp[len(l):])
}

func main() {
	rand.Seed(time.Now().UnixNano())
	a := make([]int, 100000)
	for i := range a {
		a[i] = rand.Intn(20)
	}
	// fmt.Println(a)
	t1 := time.Now()
	mergeSort(a)
	fmt.Println(time.Now().Sub(t1))
	// fmt.Println(a)
}
