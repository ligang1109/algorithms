package sort

import "testing"

func TestBubble(t *testing.T) {
	list := []int{2, 1, 6, 7, 3, 8, 9, 4, 5}
	Bubble(list)

	t.Log(list)
}
