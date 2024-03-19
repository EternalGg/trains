package divisorGame

func DivisorGame(n int) bool {
	bob := false
	for n != 1 {
		guess := n
		for n%guess != 0 {
			guess--
		}
		n -= guess
		bob = !bob
	}
	if bob {
		return true
	} else {
		return false
	}
}
