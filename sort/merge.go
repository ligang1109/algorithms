package sort

func Merge(list []int) {
	if len(list) <= 1 {
		return
	}

	msort(list, 0, len(list)-1)
}

func msort(list []int, start, end int) {
	if start == end {
		return
	}

	min := (end-start)/2 + start

	msort(list, start, min)
	msort(list, min+1, end)
	merge(list, start, min, end)
}

func merge(list []int, start, min, end int) {
	var tmpList []int
	li := start
	ri := min + 1

	for {
		if list[li] < list[ri] {
			tmpList = append(tmpList, list[li])
			li++
			if li > min {
				break
			}
		} else {
			tmpList = append(tmpList, list[ri])
			ri++
			if ri > end {
				break
			}
		}
	}

	for i := li; i <= min; i++ {
		tmpList = append(tmpList, list[i])
	}
	for i := ri; i <= end; i++ {
		tmpList = append(tmpList, list[i])
	}

	copy(list[start:], tmpList)
}
