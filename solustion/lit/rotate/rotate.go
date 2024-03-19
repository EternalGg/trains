package rotate

func Rotate(nums []int, k int) {
	if len(nums) <= k {
		cNums := make([]int, 0)
		cNums = append(cNums, nums...)
		for i := 0; i < len(nums); i++ {
			now := i + k
			if now < len(nums) {
				nums[i] = cNums[now]
			} else {
				nums[now%len(nums)] = cNums[i]
			}
		}
	} else {
		head, last := nums[len(nums)-k:], nums[:len(nums)-k]
		result := append(head, last...)
		for i := 0; i < len(nums); i++ {
			nums[i] = result[i]
		}
	}
}

func rotate(matrix [][]int) {
	list, length := make([]int, 0), len(matrix)
	for i := 0; i < length; i++ {
		for j := length - 1; j >= 0; j-- {
			list = append(list, matrix[j][i])
		}
	}
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			matrix[i][j] = list[(i*length)+j]
		}
	}
}
