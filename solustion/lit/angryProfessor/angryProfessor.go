package angryProfessor

func angryProfessor(k int32, a []int32) string {
	// Write your code here
	c := int32(0)
	for i := 0; i < len(a); i++ {
		if a[i] <= 0 {
			c++
		}
	}
	if c >= k {
		return "NO"
	} else {
		return "YES"
	}
}
