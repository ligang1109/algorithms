package sort

import (
	"fmt"
	"testing"
)

func TestIntMaxHeap(t *testing.T) {
	list := []int{10, 50, 60, 40, 80, 90, 30, 20, 70}
	heap := NewIntMaxHeap(len(list))
	for _, v := range list {
		_ = heap.Push(v)
	}

	heap.Print()

	for {
		v, e := heap.Pop()
		if e != nil {
			break
		}

		fmt.Println(v)
	}
}
