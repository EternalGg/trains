package maximumUnits

import "sort"

func maximumUnits(boxTypes [][]int, truckSize int) int {
	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})
	count := 0
	for i := 0; i < len(boxTypes); i++ {
		if truckSize >= boxTypes[i][0] {
			truckSize -= boxTypes[i][0]
			count += boxTypes[i][0] * boxTypes[i][1]
		} else {
			count += truckSize * boxTypes[i][1]
			break
		}
	}
	return count
}
