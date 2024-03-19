package diagonalPrime

import "math"

func diagonalPrime(nums [][]int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		if isprime(nums[i][i]) {
			if nums[i][i] > result {
				result = nums[i][i]
			}
		}
		if isprime(nums[i][len(nums)-1-i]) {
			if nums[i][len(nums)-1-i] > result {
				result = nums[i][len(nums)-1-i]
			}
		}
	}
	if result == 1 {
		return 0
	}
	return result
}

func isprime(j int) bool {
	for i := 2; i <= int(math.Sqrt(float64(j))); i++ {
		if j%i == 0 {
			return false
		}
	}
	return true
}
