package getTotalX

func getTotalX(a []int32, b []int32) int32 {
	// Write your code here
	result := int32(0)
	for i := a[len(a)-1]; i <= b[0]; i++ {
		for _, ai := range a {
			if i%ai != 0 {
				goto here
			}
		}
		for _, bi := range b {
			if bi%i != 0 {
				goto here
			}
		}
		result++
	here:
	}
	return result
}
