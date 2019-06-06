package sort

func bubble(list []int) {
	cnt := len(list)
	for i := 0; i < cnt-1; i++ {
		for j := cnt - 1; j > i; j-- {
			if list[j] < list[j-1] {
				tmp := list[j]
				list[j] = list[j-1]
				list[j-1] = tmp
			}
		}
	}
}
