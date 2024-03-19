package migratoryBirds

import "sort"

func migratoryBirds(arr []int32) int32 {
	// Write your code here
	iMap, max := map[int32]int{}, 0
	for i := 0; i < len(arr); i++ {
		iMap[arr[i]]++
		if iMap[arr[i]] > max {
			max = iMap[arr[i]]
		}
	}
	list := []int32{}
	for i, i2 := range iMap {
		if i2 == max {
			list = append(list, i)
		}
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i] < list[j]
	})
	return list[0]
}
