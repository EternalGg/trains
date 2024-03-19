package timeConversion

import (
	"fmt"
	"strconv"
)

func TimeConversion(s string) string {
	// Write your code here
	changed := s[:8]

	hour := changed[0:2]
	hourToInt, _ := strconv.Atoi(hour)

	switch s[8] {
	case 'A':
		if hourToInt < 12 {
			return changed
		} else {
			hourToInt -= 12

			last := changed[2:]
			head := strconv.Itoa(hourToInt)
			if hourToInt == 0 {
				head = "00"
			}
			return head + last

		}
	case 'P':
		if hourToInt < 12 {
			hourToInt += 12
			last := changed[2:]
			head := strconv.Itoa(hourToInt)
			return head + last
		} else {
			return changed
		}
	}
	fmt.Println(hour)
	fmt.Println(changed)

	return ""
}
