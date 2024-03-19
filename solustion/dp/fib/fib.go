package fib

func Fib(n int) int {
	if n == 1 {
		return 1
	}
	if n == 0 {
		return 0
	}
	fMap := make(map[int]int)
	fMap[0] = 0
	fMap[1] = 1
	for i := 2; i <= 30; i++ {
		fMap[i] = fMap[i-1] + fMap[i-2]
	}

	return fMap[n-1] + fMap[n-2]
}
