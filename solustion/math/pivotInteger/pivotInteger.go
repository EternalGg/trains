package pivotInteger

func PivotInteger(n int) int {
	list := make([]int, n)
	list[0] = 1
	for i := 1; i < n; i++ {
		list[i] += list[i-1] + i + 1
	}
	reverseList := make([]int, n)
	reverseList[n-1] = n
	for i := n - 2; i >= 0; i-- {
		reverseList[i] += reverseList[i+1] + i + 1
	}
	for i := 0; i < n; i++ {
		if reverseList[i] == list[i] {
			return i + 1
		}
	}
	return -1
}
