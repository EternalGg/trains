package isHappy

import (
	"strconv"
	"time"
)

func isHappy(n int) bool {
	if n == 1 {
		return true
	}
	if n < 10 && n != 1 {
		return false
	}
	time.Sleep(5)
	defer func() bool {
		return false
	}()
	for n != 1 {
		this := 0
		str := strconv.Itoa(n)
		for _, value := range str {
			now := int(value) - 48
			now = now * now
			this += now
		}
		if this == 1 {
			return true
		}
		n = this
	}
	return false
}
