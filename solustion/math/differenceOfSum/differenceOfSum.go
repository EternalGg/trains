package differenceOfSum

import (
	"fmt"
	"math"
	"strconv"
)

func DifferenceOfSum(nums []int) int {
	eleCount, singleCount := 0, 0

	for _, num := range nums {
		eleCount += num
	}
	for _, num := range nums {
		str := strconv.Itoa(num)
		for _, i2 := range str {
			fmt.Println(i2)
			singleCount += int(i2)
		}
	}
	fmt.Println(eleCount, singleCount)
	return int(math.Abs(float64(eleCount - singleCount)))
}
