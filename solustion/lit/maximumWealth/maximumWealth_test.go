package maximumWealth

import "testing"

func maximumWealth(accounts [][]int) int {
	max := 0
	for _, account := range accounts {
		now := 0
		for _, value := range account {
			now += value
		}
		if now > max {
			max = now
		}
	}
	return max
}

func TestMaximumWealth(t *testing.T) {

}
