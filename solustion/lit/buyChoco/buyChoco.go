package buyChoco

func BuyChoco(prices []int, money int) int {
	min1st, min2ed := min(prices[0], prices[1])
	for i := 2; i < len(prices); i++ {
		if prices[i] < min2ed {
			if prices[i] < min1st {
				min2ed = min1st
				min1st = prices[i]
			} else {
				min2ed = prices[i]
			}
		}
	}
	println(min1st, min2ed)
	if min2ed+min1st > money {
		return money
	} else {
		return money - min1st - min2ed
	}
}

func min(a, b int) (int, int) {
	if a > b {
		return b, a
	} else {
		return a, b
	}
}
