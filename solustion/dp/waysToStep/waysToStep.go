package waysToStep

func waysToStep(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	if n == 3 {
		return 4
	}
	list := make([]int, n)
	list[0] = 1
	list[1] = 2
	list[2] = 4

	for i := 3; i < n; i++ {
		list[i] = (list[i-1] + list[i-2] + list[i-3]) % 1000000007
	}
	return list[n-1]
}
