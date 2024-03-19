package temperatureTrend

import "fmt"

func TemperatureTrend(temperatureA []int, temperatureB []int) int {
	trendA, trendB := make([]int, len(temperatureA)-1), make([]int, len(temperatureA)-1)
	for i := 1; i < len(temperatureA); i++ {
		switch temperatureA[i] - temperatureA[i-1] {
		case 0:
			trendA[i] = 1
		default:
			if temperatureA[i]-temperatureA[i-1] < 0 {
				trendA[i] = 0
			} else {
				trendA[i] = 2
			}
		}
		switch temperatureB[i] - temperatureB[i-1] {
		case 0:
			trendA[i] = 1
		default:
			if temperatureB[i]-temperatureB[i-1] < 0 {
				trendA[i] = 0
			} else {
				trendA[i] = 2
			}
		}
	}
	Max, count := 0, 0
	fmt.Println(trendB, trendA)
	for i := 0; i < len(trendA); i++ {
		if trendA[i] == trendB[i] {
			count++
			if count > Max {
				Max = count
			}
		} else {
			count = 0
		}
	}
	return Max
}
