package utopianTree

func utopianTree(n int32) int32 {
	// Write your code here
	result, B := int32(1), true
	for n != 0 {
		if B {
			result *= 2
			B = false
		} else {
			result++
			B = true
		}
		n--
	}
	return result
}
