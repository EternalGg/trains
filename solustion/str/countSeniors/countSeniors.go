package countSeniors

func countSeniors(details []string) int {
	result := 0
	for _, detail := range details {
		if (detail[11] == '6' && detail[12] != '0') || detail[11] == '7' || detail[11] == '8' || detail[11] == '9' {
			result++
		}
	}
	return result
}
