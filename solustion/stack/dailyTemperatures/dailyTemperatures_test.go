package dailyTemperatures

import (
	"fmt"
	"testing"
)

func TestDailyTemperatures(t *testing.T) {
	list := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println(DailyTemperatures(list))
}
