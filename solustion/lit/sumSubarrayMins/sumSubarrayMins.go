package sumSubarrayMins

import "fmt"

func SumSubarrayMins(arr []int) int {
	length, left, right, count, level := len(arr), 0, 0, 0, 2
	for _, value := range arr {
		count += value
	}
	fmt.Println(left, right)
	for level <= length {

	}
	return count
}
