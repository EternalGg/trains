package divisorSubstrings

import (
	"fmt"
	"strconv"
)

func divisorSubstrings(num int, k int) int {
	result, str := 0, strconv.Itoa(num)
	for i := 0; i < len(str)-k; i++ {
		cut := str[i : i+k]
		fmt.Println(cut)
		now, _ := strconv.ParseInt(cut, 10, 64)

		if int64(num)%now == 0 {
			result++
		}
	}
	return result
}
