package findKthPositive

func FindKthPositive(arr []int, k int) int {
	result, pointer := 0, 0
	for k != 0 {
		if result+1 == arr[pointer] {
			if pointer != len(arr)-1 {
				pointer++
			}
		} else {
			k--
		}
		result++
	}
	return result
}
