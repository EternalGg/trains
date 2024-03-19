package minMaxGame

func MinMaxGame(nums []int) int {
	list := nums
	for len(list) != 1 {
		isMin := true
		newList := make([]int, 0)
		for i := 0; i < len(list); i += 2 {
			if isMin {
				result := min(list[i], list[i+1])
				newList = append(newList, result)
			} else {
				result := max(list[i], list[i+1])
				newList = append(newList, result)
			}
			isMin = !isMin
		}
		list = newList
	}
	return list[0]
}
func min(n1, n2 int) int {
	if n1 > n2 {
		return n2
	} else {
		return n1
	}
}
func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	} else {
		return n2
	}
}
