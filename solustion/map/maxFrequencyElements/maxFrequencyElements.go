package maxFrequencyElements

import "fmt"

func maxFrequencyElements(nums []int) int {
	max, countmap := 0, map[int]int{}
	for i := 0; i < len(nums); i++ {
		countmap[nums[i]]++
		if countmap[nums[i]] > max {
			max = countmap[nums[i]]
		}
	}
	result := 0
	for key, value := range countmap {
		if value == max {

			result += max
			fmt.Println(key, value, result)
		}
	}
	return max
}
