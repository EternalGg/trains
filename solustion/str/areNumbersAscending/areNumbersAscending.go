package areNumbersAscending

import (
	"strconv"
)

func AreNumbersAscending(s string) bool {
	list, judgeList, lastValue := make([]int, 0), make([]int64, 0), ""
	for key, letter := range s {
		if letter >= 48 && letter <= 57 {
			list = append(list, key)
		}
	}

	for key, value := range list {
		lastValue += string(s[value])
		if key == len(list)-1 {
			goto here
		}
		if list[key+1] == value+1 {
			continue
		}
	here:
		now, _ := strconv.ParseInt(lastValue, 10, 64)
		if len(judgeList) >= 1 {
			if judgeList[len(judgeList)-1] >= now {
				return false
			}
		}
		lastValue = ""
		judgeList = append(judgeList, now)
	}
	return true
}
