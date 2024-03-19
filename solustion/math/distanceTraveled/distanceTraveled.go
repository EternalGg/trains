package distanceTraveled

func distanceTraveled(mainTank int, additionalTank int) int {
	result := 0
	for mainTank != 0 {
		if mainTank >= 5 {
			if additionalTank > 0 {
				mainTank -= 4
				additionalTank -= 1
			} else {
				mainTank -= 5
			}
			result += 5
		} else {
			result += mainTank
			mainTank = 0
		}
	}
	return result * 10
}
