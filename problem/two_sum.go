package problem

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		tv := target - v
		ti, ok := m[tv]
		if ok {
			return []int{ti, i}
		}

		m[v] = i
	}

	return nil
}
