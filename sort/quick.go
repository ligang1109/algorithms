package sort

func Quick(list []int) {
	if len(list) <= 1 {
		return
	}

	quick(list, 0, len(list)-1)
}

func quick(list []int, left, pivot int) {
	if left >= pivot {
		return
	}

	li := left
	ri := pivot - 1

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
				if ri == li {
					break
				}
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
		quick(list, left, li-1)
		quick(list, li+1, pivot)
	} else {
		quick(list, left, ri)
	}

}

func swap(list []int, a, b int) {
	tmp := list[a]
	list[a] = list[b]
	list[b] = tmp
}
