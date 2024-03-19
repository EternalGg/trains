package distMoney

func DistMoney(money int, children int) int {
	result := 0
	if money < children {
		return -1
	}
	money -= children
	for money-7 >= 0 {
		result++
		money -= 7
		children--
	}
	if children == 1 && money == 3 {
		result--
	}
	return result
}
