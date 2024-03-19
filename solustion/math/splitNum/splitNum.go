package splitNum

import (
	"strconv"
)

func SplitNum(num int) int {
	its := strconv.Itoa(num)
	result := 0
	if len(its)%2 == 0 {
		first := its[:len(its)/2]
		later := its[len(its)/2:]
		f, _ := strconv.ParseInt(first, 10, 64)
		l, _ := strconv.ParseInt(later, 10, 64)
		result = int(f + l)
	} else {
		first := its[:len(its)/2]
		later := its[len(its)/2:]
		f, _ := strconv.ParseInt(first, 10, 64)
		l, _ := strconv.ParseInt(later, 10, 64)
		a := int(f + l)
		result = a
		first = its[:(len(its)/2)+1]
		later = its[(len(its)/2)+1:]
		bf, _ := strconv.ParseInt(first, 10, 64)
		bl, _ := strconv.ParseInt(later, 10, 64)
		b := int(bf + bl)
		if b > a {
			result = b
		}
	}
	return result
}
