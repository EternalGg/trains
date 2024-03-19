package timeRequiredToBuy

func TimeRequiredToBuy(tickets []int, k int) int {
	result, pointer := 0, 0
	for tickets[k] != 0 {
		if pointer == len(tickets) {
			pointer = 0
		}
		if tickets[pointer] > 0 {
			tickets[pointer] -= 1
			result++
		}
		pointer++
	}
	return result
}
