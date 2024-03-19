package minTimeToType

import (
	"fmt"
)

func MinTimeToType(word string) int {
	result, last := len(word), int32(0)

	for _, value := range word {
		now := value - 97
		nav := 0
		pos := last - now
		fmt.Println(pos)
		if pos >= 0 {
			nav = 24 - int(pos)
		} else {
			pos = 0 - pos
			nav = 24 - int(pos)
		}
		if nav > int(pos) {
			result += int(pos)
		} else {
			result += nav
		}
		fmt.Println(result)
		last = now
	}
	return result
}
