package findShortestSubArray

import "fmt"

type Count struct {
	time  int
	first int
	last  int
}

func findShortestSubArray(nums []int) int {
	cMap, list := make(map[int]*Count), make([]int, 0)
	for i, num := range nums {
		if cMap[num] == nil {
			newCount := new(Count)
			newCount.time = 1
			newCount.first = i
			newCount.last = i
			cMap[num] = newCount
			list = append(list, num)
		} else {
			cMap[num].time++
			cMap[num].last = i
		}
	}

	result, max := 0, 0
	for _, i2 := range list {
		if cMap[i2].time > max {
			result = cMap[i2].last - cMap[i2].first
			max = cMap[i2].time
		} else if cMap[i2].time == max {
			fmt.Println(cMap[i2].last - cMap[i2].first)
			if cMap[i2].last-cMap[i2].first < result {
				result = cMap[i2].last - cMap[i2].first
			}
		}
	}
	return result
}
