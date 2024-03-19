package numberOfEmployeesWhoMetTarget

func numberOfEmployeesWhoMetTarget(hours []int, target int) int {
	t := 0
	for i := 0; i < len(hours); i++ {
		if hours[i] >= target {
			t++
		}
	}
	return t
}
