package kidsWithCandies

import "testing"

func kidsWithCandies(candies []int, extraCandies int) []bool {
	max, qa := 0, make([]bool, len(candies))

	for _, candy := range candies {
		if candy > max {
			max = candy
		}
	}
	for i, candy := range candies {
		if candy+extraCandies >= max {
			qa[i] = true
		} else {
			qa[i] = false
		}
	}
	return qa
}

func TestKidsWithCandies(t *testing.T) {

}
