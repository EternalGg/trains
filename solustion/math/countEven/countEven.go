package countEven

func CountEven(num int) int {
	result := 0
	for i := 1; i <= num; i++ {
		now := plus(i)
		if now%2 == 0 {
			result++
		}
	}
	return result
}

func plus(n int) int {
	if n < 10 {
		return n
	}
	result, now := 0, 10
	for n >= now {
		now *= 10
	}

	for n != 0 {
		divide := n / now
		result += divide
		n -= divide * now
		now /= 10
	}
	return result
}
