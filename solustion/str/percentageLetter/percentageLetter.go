package percentageLetter

import "fmt"

func PercentageLetter(s string, letter byte) int {
	count, result, length := 0, 0, len(s)

	for _, value := range s {
		if byte(value) == letter {
			count++
		}
	}
	if count == 0 {
		return 0
	}
	divde := (float64(count) / float64(length)) * 100
	fmt.Println(divde)
	result = int(divde)
	return result
}
