package oneEditAway

import "math"

func oneEditAway(first string, second string) bool {
	if math.Abs(float64(len(first)-len(second))) > 1 {
		return false
	}
	m1, m2 := map[uint8]int{}, map[uint8]int{}
	for i := 0; i < len(first); i++ {
		m1[first[i]]++
	}
	for i := 0; i < len(second); i++ {
		m2[second[i]]++
	}
	if len(first) == len(second) {
		//for i, i2 := range collection {

		//}
	}
	return false
}
