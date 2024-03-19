package digitSum

import (
	"fmt"
	"strconv"
)

func digitSum(s string, k int) string {
	for len(s) <= k {
		strGroup := make([]string, 0)
		for i := 0; i < len(s); i += k {
			now := ""
			if i+k > len(s)-1 {
				now = s[i:]
			} else {
				now = s[i : i+k]
			}
			strGroup = append(strGroup, now)
		}
		nowS := ""
		for _, str := range strGroup {
			now := 0
			for _, value := range str {
				now += int(value) - 48
			}
			nowS += strconv.Itoa(now)
		}
		s = nowS
		fmt.Println(s)
	}
	return s
}
