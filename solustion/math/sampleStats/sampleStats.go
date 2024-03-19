package sampleStats

import "sort"

func sampleStats(count []int) []float64 {
	result := make([]float64, 5)
	sort.Slice(count, func(i, j int) bool {
		return count[i] < count[j]
	})
	counts, time, maxFreq, maxFreqKey := 0, make(map[int]int), 0, 0
	for i := 0; i < len(count); i++ {
		counts += count[i]
		time[count[i]]++
		if time[count[i]] > maxFreq {
			maxFreq = time[count[i]]
			maxFreqKey = count[i]
		}
	}

	result[0] = float64(count[0])
	result[1] = float64(count[1])
	result[2] = float64(counts) / float64(len(count))
	if len(count)%2 == 0 {
		result[3] = (float64(count[len(count)/2]) + float64(count[len(count)/2]-1)) / float64(2)
	} else {
		result[3] = float64(count[len(count)/2])
	}
	result[4] = float64(maxFreqKey)
	return result
}
