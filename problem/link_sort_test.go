package problem

import "testing"

func TestLinkSort(t *testing.T) {
	list := []int{2, 1, 6, 7, 3, 8, 9, 4, 5}
	//list := []int{10, 50, 60, 40, 80, 90, 30, 20, 70}
	//list := []int{3, 1}

	head := new(LinkNode)
	p := head
	for i := 0; i < len(list)-1; i++ {
		p.Value = list[i]
		p.Next = new(LinkNode)
		p = p.Next
	}
	p.Value = list[len(list)-1]

	//PrintLinkNode(head)
	LinkSort(head)
	PrintLinkNode(head)
}
