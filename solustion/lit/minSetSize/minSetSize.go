package minSetSize

import "sort"

func minSetSize(arr []int) int {
	count, length, result, add := map[int]int{}, len(arr), 0, 0
	for i := 0; i < length; i++ {
		count[arr[i]]++
	}
	list := []int{}
	for _, value := range count {
		list = append(list, value)
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i] > list[j]
	})
	for _, value := range list {
		result++
		if add+value >= length/2 {
			break
		}
	}
	return result
}
