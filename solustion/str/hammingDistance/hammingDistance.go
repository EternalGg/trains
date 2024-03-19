package hammingDistance

import (
	"strconv"
)

func HammingDistance(x int, y int) int {
	xToDouble, yToDouble, qa := strconv.FormatInt(int64(x), 2), strconv.FormatInt(int64(y), 2), 0
	larger, smaller := "", ""
	if len(xToDouble) > len(yToDouble) {
		larger = xToDouble
		smaller = yToDouble
	} else {
		larger = yToDouble
		smaller = xToDouble
	}
	xLen, yLen := len(xToDouble), len(yToDouble)
	for i := 0; i < len(larger); i++ {
		if i >= len(smaller) {
			if larger[len(larger)-1-i] == 49 {
				qa++
			}
		} else {
			if xToDouble[xLen-1-i] != yToDouble[yLen-1-i] {
				qa++
			}
		}
	}

	return qa
}
