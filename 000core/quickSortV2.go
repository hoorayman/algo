package main

import (
	"fmt"
	"math/rand"
	"time"
)

func quickSortV2(a []int) {
	rand.Seed(time.Now().UnixNano())
	quickV2(a)
}

func quickV2(a []int) {
	if len(a) > 1 {
		l, r := partitionV2(a)
		quickV2(l)
		quickV2(r)
	}
}

// make 2 parts: part 1 contains values <pivot, leave the values ==pivot, part 2 contains values > pivot
func partitionV2(a []int) ([]int, []int) {
	r := rand.Intn(len(a))
	pivot := a[r]

	i, e := 0, 0
	for j := 0; j < len(a); j++ {
		if a[j] <= pivot {
			a[i], a[j] = a[j], a[i]
			if a[i] < pivot {
				a[e], a[i] = a[i], a[e]
				e++
			}
			i++
		}
	}
	return a[:e], a[i:]
}

func main() {
	a := make([]int, 100000)
	rand.Seed(time.Now().UnixNano())
	for i := range a {
		a[i] = rand.Int()
	}
	// fmt.Println(a)
	t1 := time.Now()
	quickSortV2(a)
	fmt.Println(time.Now().Sub(t1))
	// fmt.Println(a)
}
