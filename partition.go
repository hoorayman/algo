package main

import (
	"fmt"
	"math/rand"
	"time"
)

// make 2 parts: part 1 contains values <=n, part 2 contains values > n
func makePartition(a []int, n int) ([]int, []int) {
	i := 0
	for j := 0; j < len(a); j++ {
		if a[j] <= n {
			a[i], a[j] = a[j], a[i]
			i++
		}
	}
	return a[:i], a[i:]
}

// make 2 parts: part 1 contains values <n, leave the values ==n, part 2 contains values > n
func makePartitionV2(a []int, n int) ([]int, []int) {
	i, e := 0, 0
	for j := 0; j < len(a); j++ {
		if a[j] <= n {
			a[i], a[j] = a[j], a[i]
			if a[i] < n {
				a[e], a[i] = a[i], a[e]
				e++
			}
			i++
		}
	}
	return a[:e], a[i:]
}

func main() {
	a := make([]int, 10)
	rand.Seed(time.Now().UnixNano())
	for i := range a {
		a[i] = rand.Intn(10)
	}
	n := rand.Intn(20)
	fmt.Printf("array is %v, num is %d\n", a, n)
	t1 := time.Now()
	l, r := makePartitionV2(a, n)
	fmt.Println(time.Now().Sub(t1))
	fmt.Printf("after partitionV2 array is %v, left part is %v, right part is %v\n", a, l, r)
}
