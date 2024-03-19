package hurdleRace

func hurdleRace(k int32, height []int32) int32 {
	// Write your code here
	max := int32(0)
	for i := 0; i < len(height); i++ {
		if height[i]-k > max {
			max = height[i] - k
		}
	}
	return max
}
