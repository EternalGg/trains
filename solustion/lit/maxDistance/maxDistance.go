package maxDistance

func maxDistance(colors []int) int {
	max := 0
	for i := 0; i < len(colors); i++ {
		for j := i; j < len(colors); j++ {
			if colors[i] != colors[j] {
				if j-i > max {
					max = j - i
				}
			}
		}
	}
	return max
}
