package sort

import (
	"errors"
	"fmt"
)

type IntMaxHeap struct {
	len  int
	cap  int
	heap []int
}

func NewIntMaxHeap(cap int) *IntMaxHeap {
	return &IntMaxHeap{
		len:  0,
		cap:  cap,
		heap: make([]int, cap),
	}
}

func (imh *IntMaxHeap) Push(v int) error {
	if imh.len == imh.cap {
		return errors.New("heap full")
	}

	imh.heap[imh.len] = v
	imh.swim(imh.len)
	imh.len++

	return nil
}

func (imh *IntMaxHeap) Pop() (int, error) {
	if imh.len == 0 {
		return 0, errors.New("heap empty")
	}

	result := imh.heap[0]

	imh.len--
	imh.heap[0] = imh.heap[imh.len]
	imh.sink(0)

	return result, nil
}

func (imh *IntMaxHeap) Len() int {
	return imh.len
}

func (imh *IntMaxHeap) Cap() int {
	return imh.cap
}

func (imh *IntMaxHeap) Print() {
	fmt.Println(imh.heap)
}

func (imh *IntMaxHeap) parent(i int) int {
	return (i - 1) / 2
}

func (imh *IntMaxHeap) lchild(i int) int {
	return i*2 + 1
}

func (imh *IntMaxHeap) rchild(i int) int {
	return i*2 + 2
}

func (imh *IntMaxHeap) swim(i int) {
	if i == 0 {
		return
	}

	parent := imh.parent(i)
	if imh.heap[parent] < imh.heap[i] {
		imh.swap(parent, i)
		imh.swim(parent)
	}
}

func (imh *IntMaxHeap) sink(i int) {
	lchild := imh.lchild(i)
	if lchild >= imh.len {
		return
	}

	max := lchild
	rchild := imh.rchild(i)
	if rchild < imh.len {
		if imh.heap[rchild] > imh.heap[lchild] {
			max = rchild
		}
	}

	if imh.heap[i] < imh.heap[max] {
		imh.swap(i, max)
		imh.sink(max)
	}
}

func (imh *IntMaxHeap) swap(a, b int) {
	tmp := imh.heap[a]
	imh.heap[a] = imh.heap[b]
	imh.heap[b] = tmp
}
