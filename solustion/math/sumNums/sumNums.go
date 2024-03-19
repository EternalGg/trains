package sumNums

func sumNums(n int) int {

	return n + sumNums(n-1)
}
