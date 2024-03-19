package largestNumber

import (
	"sort"
	"strconv"
)

func LargestNumber(nums []int) string {
	strList := make([]string, 0)
	for i := 0; i < len(nums); i++ {
		strList = append(strList, strconv.Itoa(nums[i]))
	}
	sort.Slice(strList, func(i, j int) bool {
		if strList[i][0] == strList[j][0] {
			return strList[i]+strList[j] > strList[j]+strList[i]
		}
		return strList[i] > strList[j]
	})
	result := ""
	for i := 0; i < len(strList); i++ {
		result += strList[i]
	}
	return result
}
