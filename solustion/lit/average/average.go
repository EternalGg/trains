package average

import "fmt"

func average(salary []int) float64 {
	if len(salary) <= 2 {
		return 0
	}
	min, max, count := salary[0], salary[0], 0
	for _, value := range salary {
		count += value
		if value > max {
			max = value
		}
		if value < min {
			min = value
		}
	}
	fmt.Println(min, max)
	count -= min + max
	fCount := float64(count)
	divde := float64(len(salary) - 2)
	return fCount / divde
}
