package plusMinus

import (
	"fmt"
	"strconv"
)

func PlusMinus(arr []int32) {
	// Write your code here
	pos, na, zero, qa := float64(0), float64(0), float64(0), make([]float64, 3)

	for _, number := range arr {
		if number == 0 {
			zero++
			continue
		}
		if number > 0 {
			pos++
		} else {
			na++
		}
	}

	qa[0] = (pos / float64(len(arr)))
	qa[1] = (na / float64(len(arr)))
	qa[2] = (zero / float64(len(arr)))
	ans1 := strconv.FormatFloat(qa[0], 'f', 6, 32)
	ans2 := strconv.FormatFloat(qa[1], 'f', 6, 32)
	ans3 := strconv.FormatFloat(qa[2], 'f', 6, 32)
	fmt.Println(ans1)
	fmt.Println(ans2)
	fmt.Println(ans3)

}
