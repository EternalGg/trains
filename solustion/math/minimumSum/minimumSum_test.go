package minimumSum

import (
	"fmt"
	"strconv"
	"testing"
)

func TemplateSort(num []int32) []int32 {
	qa, max := make([]int32, 4), num[0]
	for i := 0; i < 4; i++ {
		key := 0
		for j := 0; j < 4; j++ {
			if num[j] > max {
				max = num[j]
				key = j
			}
		}
		num[key] = 0
		qa[i] = max
		max = 0
	}
	return qa
}

func minimumSum(num int) int {
	qa, nToStr, cuts := 0, strconv.Itoa(num), make([]int32, 0)
	for _, value := range nToStr {
		now := value - 48
		cuts = append(cuts, now)
	}
	cuts = TemplateSort(cuts)
	qa += int(cuts[0]) + int(cuts[1])
	qa += int(cuts[2]*10) + int(cuts[3]*10)

	return qa
}

func TestMinimumSum(t *testing.T) {
	fmt.Println(minimumSum(1234))
}
