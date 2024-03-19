package dividePlayers

func DividePlayers(skill []int) int64 {
	if len(skill) == 2 {
		return int64(skill[0] * skill[1])
	}
	m := map[int]int{}
	count := 0
	max := 0
	for i := 0; i < len(skill); i++ {
		count += skill[i]
		if skill[i] > max {
			max = skill[i]
		}
		m[skill[i]]++
	}
	avg := (count / (len(skill) / 2))
	if avg < max {
		return -1
	}

	result := int64(0)
	for key, value := range m {
		if m[avg-key] != value {
			return -1
		} else {
			if key*2 == avg {
				result += int64((key * (avg - key)) * (value / 2))
				m[key] = 0
			} else {
				m[key] = 0
				m[avg-key] = 0
				result += int64((key * (avg - key)) * value)
			}
		}
	}
	return result
}
