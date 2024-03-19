package maximumValue

import "strconv"

func MaximumValue(strs []string) int {
	max := 0
	for _, str := range strs {
		localMax := int32(0)
		num, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			localMax = int32(len(str))
		} else {
			localMax = int32(num)
		}
		if int(localMax) > max {
			max = int(localMax)
		}
	}
	return max
}
