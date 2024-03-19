package constructRectangle

func ConstructRectangle(area int) []int {
	result := make([]int, 2)
	time := 1
	result[0] = area
	result[1] = 1
	for !(time >= area/2) {
		time++
		if (area % (time)) == 0 {
			if time > area/time {
				return result
			}
			result[0] = area / time
			result[1] = time

		}
	}
	return result
}
