package minmaxSum

import "fmt"

func MiniMaxSum(arr []int32) {
	// Write your code here
	min, max := 0, 0
	for i, value := range arr {
		if value > arr[max] {
			max = i
		}
		if value < arr[min] {
			min = i
		}
	}
	minsum, maxsum := int64(0), int64(0)
	for i := 0; i < len(arr); i++ {
		if i == min {
			continue
		} else {
			maxsum += int64(arr[i])
		}
	}

	for i := 0; i < len(arr); i++ {
		if i == max {
			continue
		} else {
			minsum += int64(arr[i])
		}
	}

	fmt.Println(minsum, maxsum)
}
