package canBeIncreasing

import "fmt"

func canBeIncreasing(nums []int) bool {
	diff := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			fmt.Println(nums[i])
			diff++
		}
	}
	if diff > 1 {
		return false
	} else {
		return true
	}
}
