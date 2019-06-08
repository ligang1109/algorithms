package search

func Binary(list []int, v int) int {
	end := len(list) - 1
	if end < 0 {
		return -1
	}

	start := 0

	for {
		middle := (end-start)/2 + start

		if list[middle] == v {
			return middle
		} else if list[middle] < v {
			start = middle + 1
		} else {
			end = middle - 1
		}

		if start == end {
			break
		}
	}

	if list[start] == v {
		return start
	}

	return -1
}
