package mostFrequent

func MostFrequent(nums []int, key int) int {
	list, iMap := make([]int, 0), make(map[int]int)
	for i := 0; i <= len(nums)-2; i++ {
		if nums[i] == key {
			if iMap[nums[i+1]] == 0 {
				list = append(list, nums[i+1])
			}
			iMap[nums[i+1]]++
		}
	}
	max := list[0]
	for i := 0; i < len(list); i++ {
		if iMap[list[i]] > iMap[max] {
			max = list[i]
		}
	}
	return max
}
