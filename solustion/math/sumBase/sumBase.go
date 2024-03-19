package sumBase

import (
	"strconv"
)

func SumBase(n int, k int) int {
	qa := int32(0)
	z := strconv.FormatInt(int64(n), k)
	for _, i2 := range z {
		qa += i2 - 48
	}
	return int(qa)
}
