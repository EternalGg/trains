package frequencySort

import "testing"

func TestFrequencySort(t *testing.T) {
	list := []int{-1, 1, -6, 4, 5, -6, 1, 4, 1}
	FrequencySort(list)
}
