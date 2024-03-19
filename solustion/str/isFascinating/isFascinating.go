package isFascinating

import (
	"strconv"
	"strings"
)

func isFascinating(n int) bool {
	if strings.Contains(strconv.Itoa(n), "0") {
		return false
	} else {
		double, thried := n*2, n*3
		comble := strconv.Itoa(double) + strconv.Itoa(n) + strconv.Itoa(thried)
		if strings.Contains(strconv.Itoa(n), "0") || len(comble) != 9 {
			return false
		} else {
			for i := 1; i <= 9; i++ {
				if !strings.Contains(comble, strconv.Itoa(i)) {
					return false
				}
			}
			return true
		}
	}
}
