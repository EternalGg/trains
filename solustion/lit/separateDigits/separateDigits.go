package separateDigits

import "strconv"

func SeparateDigits(nums []int) []int {
	result := make([]int, 0)

	for _, num := range nums {
		if num < 10 {
			result = append(result, num)
		} else {
			list := SplitNumber(num)
			result = append(result, list...)
		}
	}
	return result
}

func SplitNumber(num int) []int {
	result := make([]int, 0)
	str := strconv.Itoa(num)
	for _, i2 := range str {
		result = append(result, int(i2)-48)
	}
	return result
}
