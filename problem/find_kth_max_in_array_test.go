package problem

import "testing"

func TestFindKthMaxInArray(t *testing.T) {
	//list := []int{2, 1, 6, 7, 3, 8, 9, 4, 5}
	list := []int{10, 50, 60, 40, 80, 90, 30, 20, 70}

	for i := 0; i < len(list); i++ {
		t.Log(FindKthMaxInArray(list, i+1))
	}
}
