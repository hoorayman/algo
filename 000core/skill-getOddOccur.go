package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func getOddOccur(s []int) (int, error) {
	if len(s) == 0 {
		return 0, errors.New("empty list found")
	}
	xor := 0
	for _, v := range s {
		xor = xor ^ v
	}
	return xor, nil
}

func main() {
	var s []int
	rand.Seed(time.Now().UnixNano())
	// add some even occur numbers
	for i := 0; i < 4; i++ {
		f := [3]int{2, 4, 6}[rand.Intn(3)]
		n := rand.Intn(30)
		var t []int
		for j := 0; j < f; j++ {
			t = append(t, n)
		}
		s = append(s, t...)
	}
	// add one odd occur number
	f := [3]int{1, 3, 5}[rand.Intn(3)]
	n := rand.Intn(30)
	var t []int
	for i := 0; i < f; i++ {
		t = append(t, n)
	}
	s = append(s, t...)
	// shuffle
	for i := len(s) - 1; i > 0; i-- { // Fisher–Yates shuffle
		j := rand.Intn(i + 1)
		s[i], s[j] = s[j], s[i]
	}
	if v, e := getOddOccur(s); e == nil {
		fmt.Println(s, v)
	}
}
