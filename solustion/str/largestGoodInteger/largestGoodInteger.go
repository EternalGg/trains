package largestGoodInteger

import "strconv"

func LargestGoodInteger(num string) string {
	strGroup, lastNum, last := make([]string, 0), int32(0), ""
	for _, value := range num {
		if value == lastNum {
			last += string(lastNum)
			if len(last) == 3 {
				strGroup = append(strGroup, last)
			}
		} else {
			lastNum = value
			last = string(value)
		}
	}
	numList := make([]int, 0)
	for _, str := range strGroup {
		num, _ := strconv.ParseInt(str, 10, 32)
		numList = append(numList, int(num))
	}
	if len(numList) == 0 {
		return ""
	}
	max := 0
	for _, i2 := range numList {
		if i2 > max {
			max = i2
		}
	}
	if max == 0 {
		return "000"
	}
	return strconv.Itoa(max)
}
