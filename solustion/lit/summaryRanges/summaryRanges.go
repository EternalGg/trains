package summaryRanges

import (
	"fmt"
	"strconv"
)

func SummaryRanges(nums []int) []string {
	result := make([]string, 0)
	if len(nums) == 0 {
		return result
	}
	if len(nums) == 1 {
		now := strconv.Itoa(nums[0])
		result = append(result, now)
		return result
	}
	strat, last, now := nums[0], nums[0], ""
	for i := 1; i < len(nums); i++ {
		if nums[i] == last+1 {
			last = nums[i]
		} else {
			if last-strat >= 1 {
				now = fmt.Sprint(strat, "->", last)
			} else {
				now += strconv.Itoa(last)
			}
			result = append(result, now)
			strat = nums[i]
			last = nums[i]
			now = ""
		}
	}

	if last-strat >= 1 {
		now = fmt.Sprint(strat, "->", last)
	} else {
		now += strconv.Itoa(last)
	}
	result = append(result, now)

	return result
}
