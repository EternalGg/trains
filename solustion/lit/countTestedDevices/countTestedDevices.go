package countTestedDevices

func countTestedDevices(batteryPercentages []int) int {
	result := 0

	for i := 0; i < len(batteryPercentages); i++ {
		if batteryPercentages[i] > 0 {
			result++
			for i := i + 1; i < len(batteryPercentages); i++ {
				batteryPercentages[i]--
			}
		}
	}
	return result
}
