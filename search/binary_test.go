package search

import "testing"

func TestBinary(t *testing.T) {
	list := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	for _, v := range []int{3, 7, 11, 4, 5, 0, 9, 10, 1} {
		t.Log(Binary(list, v))
	}
}
