package distanceBetweenBusStops

import "testing"

func distanceBetweenBusStops(distance []int, start int, destination int) int {
	negative, postive := 0, 0
	if destination < start {
		temple := start
		start = destination
		destination = temple
	}

	for i, i2 := range distance {
		if i >= start && i <= destination {
			postive += i2
		} else {
			negative += i2
		}
	}
	if negative > postive {
		return postive
	} else {
		return negative
	}
}

func TestDistanceBetweenBusStops(t *testing.T) {
	//distanceBetweenBusStops()
}
