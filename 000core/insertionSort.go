package main

import (
	"fmt"
	"math/rand"
	"time"
)

func insertionSort(s []int) []int {
	for i := 0; i < len(s); i++ {
		for j := i - 1; j > -1; j-- {
			if s[j+1] < s[j] {
				s[j+1], s[j] = s[j], s[j+1]
			} else {
				break
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
	insertionSort(s)
	fmt.Println(time.Now().Sub(t1))
}
