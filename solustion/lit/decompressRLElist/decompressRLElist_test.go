package decompressRLElist

import "testing"

func decompressRLElist(nums []int) []int {
	qa := make([]int, 0)
	for i := 1; i < len(nums); i += 2 {
		now := make([]int, 0)
		for j := 0; j < nums[i-1]; j++ {
			now = append(now, nums[i])
		}
		qa = append(qa, now...)
	}
	return qa
}

func TestDecompressRLElist(t *testing.T) {

}
