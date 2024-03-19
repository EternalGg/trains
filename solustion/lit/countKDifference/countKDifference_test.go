package countKDifference

import (
	"math"
	"testing"
)

func countKDifference(nums []int, k int) int {
	qa := 0
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			if math.Abs(float64(nums[i]-nums[j])) == float64(k) {
				qa++
			}
		}
	}
	return qa
}

func TestCountKDifference(t *testing.T) {

}
