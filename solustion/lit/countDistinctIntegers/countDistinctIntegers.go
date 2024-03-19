package countDistinctIntegers

import "strconv"

func CountDistinctIntegers(nums []int) int {
	result, uMap := 0, make(map[int64]bool)
	for i := 0; i < len(nums); i++ {
		now := nums[i]

		str := converseString(strconv.Itoa(now))
		parseInt, _ := strconv.ParseInt(str, 10, 64)
		if uMap[int64(now)] == false && parseInt != int64(now) {
			result += 2
			uMap[int64(now)] = true
			uMap[parseInt] = true
		}
	}
	return result
}

func converseString(str string) string {
	result := ""
	for _, letter := range str {
		result = string(letter) + result
	}
	return result
}
