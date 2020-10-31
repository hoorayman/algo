package main

import (
	"fmt"
	"math/rand"
	"time"
)

func bubbleSort(s []int) []int {
	for i := len(s); i > 1; i-- {
		for j := 0; j+1 < i; j++ {
			if s[j+1] < s[j] {
				s[j+1], s[j] = s[j], s[j+1]
			}
		}
	}
	return s
}

func main() {
	s := make([]int, 100000)
	for i := range s {
		s[i] = rand.Int()
	}
	t1 := time.Now()
	bubbleSort(s)
	fmt.Println(time.Now().Sub(t1))
}
