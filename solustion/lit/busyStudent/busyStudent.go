package busyStudent

func BusyStudent(startTime []int, endTime []int, queryTime int) int {
	qa := 0
	for i := 0; i < len(startTime); i++ {
		if queryTime >= startTime[i] && queryTime <= endTime[i] {
			qa++
		}
	}
	return qa
}
