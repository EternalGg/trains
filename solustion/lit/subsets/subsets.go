package subsets

func Subsets(nums []int) [][]int {

	result := [][]int{}
	result = append(result, []int{})
	for _, num := range nums {
		for _, sub := range result {
			temp := make([]int, len(sub), len(sub)+1)
			copy(temp, sub)
			temp = append(temp, num)
			result = append(result, temp)
		}
	}
	return result
}
