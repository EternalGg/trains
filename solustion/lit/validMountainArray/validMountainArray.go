package validMountainArray

import "fmt"

func validMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}
	max := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] >= arr[max] {
			fmt.Println(arr[i], max)
			if arr[i] == arr[max] {
				return false
			} else {
				max = i
			}
		}
	}
	if max == 0 {
		return false
	}
	for i := max + 1; i < len(arr); i++ {
		// 后一个比前面大
		if arr[i] <= arr[i-1] {
			return false
		}
	}
	for i := 0; i < max; i++ {
		if arr[i] <= arr[i-1] {

			return false
		}
	}
	return true
}
