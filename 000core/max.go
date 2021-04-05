package main

import (
	"fmt"
	"math/rand"
	"time"
)

func max(s []int) int {
	if len(s) > 1 {
		mid := len(s) >> 1
		l := max(s[:mid])
		r := max(s[mid:])
		if l < r {
			return r
		}
		return l
	} else if len(s) == 1 {
		return s[0]
	}
	panic("can not get max from empty list!")
}

func main() {
	t1 := time.Now()
	s := make([]int, 10)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		s[i] = rand.Intn(50)
	}
	fmt.Println(s, max(s), time.Now().Sub(t1))
}
