package minNumberOfHours

func MinNumberOfHours(initialEnergy int, initialExperience int, energy []int, experience []int) int {
	result := 0
	eCount := 0
	for _, i2 := range energy {
		eCount += i2
	}
	if eCount > initialEnergy {
		result += eCount - initialEnergy + 1
	}

	for i := 0; i < len(experience); i++ {
		if initialExperience <= experience[i] {
			result += experience[i] - initialExperience + 1
			initialExperience += experience[i] + 1
		} else {
			initialExperience += experience[i]
		}

	}
	return result
}
