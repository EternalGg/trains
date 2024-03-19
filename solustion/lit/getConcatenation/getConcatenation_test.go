package getConcatenation

import "testing"

func getConcatenation(nums []int) []int {
	nums = append(nums, nums...)
	return nums
}
func TestGetConcatenation(t *testing.T) {

}
