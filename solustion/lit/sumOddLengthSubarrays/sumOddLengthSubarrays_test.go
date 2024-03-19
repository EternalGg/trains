package sumOddLengthSubarrays

import (
	"fmt"
	"testing"
)

func sumOddLengthSubarrays(arr []int) int {
	qa := 0
	for i := 1; i <= len(arr); i += 2 {
		for j := 0; j < len(arr)-i+1; j++ {
			now := 0
			for k := j; k < j+i; k++ {
				now += arr[k]
			}
			qa += now
		}
	}

	return qa
}

func TestSumOddLengthSubarrays(t *testing.T) {
	fmt.Println(
		sumOddLengthSubarrays([]int{1, 4, 2, 5, 3}))
}
