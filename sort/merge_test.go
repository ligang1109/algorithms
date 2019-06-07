package sort

import "testing"

func TestMerge(t *testing.T) {
	//list := []int{2, 1, 6, 7, 3, 8, 9, 4, 5}
	list := []int{10, 50, 60, 40, 80, 90, 30, 20, 70}

	Merge(list)

	t.Log(list)
}
