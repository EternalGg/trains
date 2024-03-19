package minMaxDifference

import (
	"fmt"
	"strconv"
)

func MinMaxDifference(num int) int {
	str := []byte(strconv.Itoa(num))
	mFirst := str[0]
	for i, b := range str {
		if b != 57 {
			mFirst = str[i]
			break
		}
	}
	first := str[0]
	max, min := make([]byte, len(str)), make([]byte, len(str))
	for i := 0; i < len(str); i++ {
		if str[i] == first {
			min[i] = byte(48)
		} else {
			min[i] = str[i]
		}
		if str[i] == mFirst {
			max[i] = byte(57)
		} else {
			max[i] = str[i]
		}
		fmt.Println(string(max), string(min))
	}
	maxToInt, _ := strconv.ParseInt(string(max), 10, 64)
	minToInt, _ := strconv.ParseInt(string(min), 10, 64)

	return int(maxToInt - minToInt)
}
