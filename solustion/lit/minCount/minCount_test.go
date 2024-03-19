package minCount

import "testing"

func minCount(coins []int) int {
	qa := 0
	for _, coin := range coins {
		if coin%2 == 0 {
			qa += coin / 2
		} else {
			qa += (coin + 1) / 2
		}
	}
	return qa
}

func TestMinCount(t *testing.T) {

}
