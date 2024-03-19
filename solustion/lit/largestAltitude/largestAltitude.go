package largestAltitude

func LargestAltitude(gain []int) int {
	Altitude, max := make([]int, len(gain)+1), 0
	Altitude[0] = 0

	for i, Alititudes := range gain {
		Altitude[i+1] = Altitude[0] + Alititudes
		if Altitude[i+1] > max {
			max = Altitude[i+1]
		}
	}
	return max
}
