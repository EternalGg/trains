package moveZeroes

func MoveZeroes(nums []int) {
	list, point := make([]int, 0), 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			list = append(list, -1)
			point++
		} else {
			list = append(list, point)
		}
	}
	for i := 0; i < len(list); i++ {
		if list[i] != -1 {
			nums[i-list[i]] = nums[i]
		}
	}
	for j := len(nums) - 1; j >= len(nums)-point; j-- {
		nums[j] = 0
	}
}
