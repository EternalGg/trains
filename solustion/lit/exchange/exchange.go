package exchange

func exchange(nums []int) []int {
	odd, even := make([]int, 0), make([]int, 0)

	for _, num := range nums {
		if num%2 == 0 {
			even = append(even, num)
		} else {
			odd = append(odd, num)
		}
	}
	result := append(odd, even...)
	return result
}
