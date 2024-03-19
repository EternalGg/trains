package replaceElements

func ReplaceElements(arr []int) []int {
	qa, max := make([]int, len(arr)), -1

	for i := len(arr) - 1; i >= 0; i-- {
		global := arr[i]
		qa[i] = max
		if global > max {
			max = global
		}
	}
	return qa
}
