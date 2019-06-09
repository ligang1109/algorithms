package problem

func FindKthMaxInArray(list []int, k int) int {
	if (len(list) < k) {
		return -1
	}

	return quickFind(list, 0, len(list)-1, len(list)-k)
}

func quickFind(list []int, left, pivot, k int) int {
	if left >= pivot {
		return list[left]
	}

	right := pivot - 1
	li := left
	ri := right

	for {
		for {
			if list[li] < list[pivot] {
				li++
				if li == pivot {
					break
				}
			} else {
				break
			}
		}

		for {
			if ri <= li {
				break
			}

			if list[ri] >= list[pivot] {
				ri--
			} else {
				break
			}
		}

		if li < ri {
			swap(list, li, ri)
		} else {
			break
		}
	}

	if li == ri {
		swap(list, li, pivot)
		if li == k {
			return list[li]
		} else if li < k {
			return quickFind(list, li+1, right, k)
		}
		return quickFind(list, left, li-1, k)
	}

	if pivot == k {
		return list[pivot]
	}

	return quickFind(list, left, ri, k)
}

func swap(list []int, a, b int) {
	tmp := list[a]
	list[a] = list[b]
	list[b] = tmp
}
