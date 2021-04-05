package main

import (
	"fmt"
	"math/rand"
	"time"
)

func selectionSort(s []int) []int {
	for i := 0; i < len(s); i++ {
		minIndex := i
		for j := i + 1; j < len(s); j++ {
			if s[j] < s[minIndex] {
				minIndex = j
			}
		}
		s[i], s[minIndex] = s[minIndex], s[i]
	}
	return s
}

func main() {
	s := make([]int, 100000)
	for i := range s {
		s[i] = rand.Int()
	}
	t1 := time.Now()
	selectionSort(s)
	fmt.Println(time.Now().Sub(t1))
}
