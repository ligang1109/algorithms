package problem

import "fmt"

type LinkNode struct {
	Value int
	Next  *LinkNode
}

func LinkSort(head *LinkNode) {
	if head == nil || head.Next == nil {
		return
	}

	end := head.Next
	for {
		if end.Next != nil {
			end = end.Next
		} else {
			break
		}
	}

	linkSort(head, end)
}

func linkSort(head, end *LinkNode) {
	if head == nil || head == end {
		return
	}

	great := head
	less := head.Next

	for {
		for {
			if great.Next.Value < head.Value {
				great = great.Next
				if great.Next == end.Next {
					break
				}
			} else {
				break
			}
		}

		less = great.Next
		for {
			if less == end.Next {
				break
			}
			if less.Value > head.Value {
				less = less.Next
			} else {
				break
			}
		}

		if less != end.Next {
			swapNode(great.Next, less)
			great = great.Next
		} else {
			break
		}
	}

	if great != head {
		swapNode(head, great)
	}

	linkSort(head, great)
	linkSort(great.Next, end)
}

func swapNode(a, b *LinkNode) {
	tmp := a.Value
	a.Value = b.Value
	b.Value = tmp
}

func PrintLinkNode(head *LinkNode) {
	for p := head; p != nil; p = p.Next {
		fmt.Println(p.Value)
	}
}
