package gradingStudents

func GradingStudents(grades []int32) []int32 {
	// Write your code here

	for i := 0; i < len(grades); i++ {
		if grades[i] < 38 {
			continue
		}
		if grades[i]%5 == 3 {
			grades[i] += 2
		} else if grades[i]%5 == 4 {
			grades[i] += 1
		}
	}
	return grades
}
