package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func get2OddOccur(s []int) (int, int, error) {
	if len(s) == 0 {
		return 0, 0, errors.New("empty list found")
	}
	xor := 0
	for _, v := range s {
		xor = xor ^ v
	} // if the two odd number is x and y, now xor == x^y

	rightMostOne := xor & ((^xor) + 1) // so, the right most one of x or y must be equal to variable rightMostOne
	xor2 := 0
	for _, v := range s {
		if (v & ((^v) + 1)) == rightMostOne {
			xor2 = xor2 ^ v
		}
	}
	return xor2, xor ^ xor2, nil
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
	// add 2 odd occur number
	for i := 0; i < 2; i++ {
		f := [3]int{1, 3, 5}[rand.Intn(3)]
		n := rand.Intn(30)
		var t []int
		for j := 0; j < f; j++ {
			t = append(t, n)
		}
		s = append(s, t...)
	}
	// shuffle
	for i := len(s) - 1; i > 0; i-- { // Fisher–Yates shuffle
		j := rand.Intn(i + 1)
		s[i], s[j] = s[j], s[i]
	}
	if v1, v2, e := get2OddOccur(s); e == nil {
		fmt.Println(s, v1, v2)
	}
}
