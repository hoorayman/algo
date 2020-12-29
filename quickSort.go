package main

import (
	"fmt"
	"math/rand"
	"time"
)

func quickSort(a []int) {
	rand.Seed(time.Now().UnixNano())
	quick(a)
}

func quick(a []int) {
	if len(a) > 1 {
		l, r := partition(a)
		quick(l)
		quick(r)
	}
}

func partition(a []int) ([]int, []int) {
	r := rand.Intn(len(a))
	a[0], a[r] = a[r], a[0] // fix the random pivot to a[0], for easy compare

	i, j := 1, len(a)-1
	for {
		for ; i < len(a) && a[i] < a[0]; i++ { // ignore head unnecessary value, cannot use a[i] <= pivot because will cause unbalanced partition, for example [5 4 5 4 5] divide to [5 4 5 4 5] and []
		}
		for ; j > 0 && a[j] > a[0]; j-- { // ignore tail unnecessary value, cannot use a[j] >= pivot because will cause unbalanced partition
		}
		if i < j {
			a[i], a[j] = a[j], a[i] // two-way exchange, in case for most values are same
			i++                     // even if most values are same, we can divide balance, because i++,j-- and i<j, i and j go in opposite direction
			j--
		} else {
			break
		}
	}
	if i == len(a) { // if this happen, the pivot is the greatest value in left. if use renturn a[:i], a[i:] there a[i:] will be empty [], cause unnecessary or endless partition loop
		a[0], a[i-1] = a[i-1], a[0] // put pivot to rightmost because it is the greatest value
		i--
	}
	return a[:i], a[i:]
}

func main() {
	a := make([]int, 100000)
	for i := range a {
		a[i] = rand.Int()
	}
	// fmt.Println(a)
	t1 := time.Now()
	quickSort(a)
	fmt.Println(time.Now().Sub(t1))
	// fmt.Println(a)
}
