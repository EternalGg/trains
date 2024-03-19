package isSameAfterReversals

import "strconv"

func IsSameAfterReversals(num int) bool {
	nToStr := strconv.Itoa(num)
	if nToStr[len(nToStr)-1] == 48 {
		return false
	}
	return true
}
