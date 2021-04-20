// 两个链表相加，不能先转换成int再相加，因为会有溢出的情况发生
// 所以直接模拟加法器的运作原理直接相加即可
package main

import (
	"container/list"
	"fmt"
)

func queueAdd(a, b *list.List) *list.List {
	var listA, listB []int
	for a.Len() > 0 {
		e := a.Front()
		listA = append(listA, e.Value.(int))
		a.Remove(e)
	}
	for b.Len() > 0 {
		e := b.Front()
		listB = append(listB, e.Value.(int))
		b.Remove(e)
	}

	result := list.New()
	tmp := listAdd(listA, listB)
	for _, v := range tmp {
		result.PushBack(v)
	}
	return result
}

func listAdd(a, b []int) []int {
	diff := len(a) - len(b)
	switch { // 使a和b等长
	case diff < 0:
		a = append(a, make([]int, -diff)...)
	case diff > 0:
		b = append(b, make([]int, diff)...)
	}

	result := make([]int, len(a))
	carry := 0
	for i := 0; i < len(a); i++ {
		sum := a[i] + b[i] + carry
		if sum >= 10 {
			result[i] = sum - 10
			carry = 1
		} else {
			result[i] = sum
			carry = 0
		}
	}
	if carry == 1 {
		result = append(result, 1)
	}
	return result
}

func main() {
	a, b := list.New(), list.New()
	a.PushBack(7)
	a.PushBack(8)
	a.PushBack(9)
	b.PushBack(7)
	b.PushBack(0)
	b.PushBack(2)
	b.PushBack(7)
	result := queueAdd(a, b)
	for result.Len() > 0 {
		e := result.Front()
		fmt.Println(e.Value.(int))
		result.Remove(e)
	}
}
