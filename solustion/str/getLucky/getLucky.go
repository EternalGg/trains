package getLucky

import (
	"fmt"
	"strconv"
)

func GetLucky(s string, k int) int {
	cut := ""
	cut = StrToInt(s)
	for i := 0; i < k; i++ {
		cut = SinglePlus(cut)
	}
	result, _ := strconv.ParseInt(cut, 10, 64)
	return int(result)
}
func StrToInt(str string) string {
	sToByte := []byte(str)
	byteToStr := ""
	for _, byte := range sToByte {
		now := int(byte - 96)
		byteToStr += strconv.Itoa(now)
	}
	fmt.Println(byteToStr)
	return byteToStr
}
func SinglePlus(num string) string {
	result, plus := "", 0

	for _, word := range num {
		plus += int(word - 48)
	}
	result += strconv.Itoa(plus)
	return result
}
