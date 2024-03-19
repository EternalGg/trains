package canCompleteCircuit

func canCompleteCircuit(gas []int, cost []int) int {
	zeroPoint, count := make([]int, 0), 0
	for i := 0; i < len(gas); i++ {
		gas[i] -= cost[i]
		count += gas[i]
		if count == 0 {
			zeroPoint = append(zeroPoint, i)
		}
	}
	for i := 0; i < len(zeroPoint); i++ {
		if zeroPoint[i] == 0 {
			zeroPoint[i] = len(zeroPoint) - 1
		} else {
			zeroPoint[i]--
		}
	}
	if len(zeroPoint) == 0 {
		return -1
	} else {
		return zeroPoint[0]
	}
}
