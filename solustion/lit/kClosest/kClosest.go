package kClosest

import (
	"fmt"
	"math"
	"sort"
)

func KKClosest(points [][]int, k int) [][]int {
	distance, dMap := make([]int, len(points)), make(map[int][][]int)
	for i := 0; i < len(points); i++ {
		distance[i] = int(math.Pow(float64(points[i][0]), 2)) + int(math.Pow(float64(points[i][1]), 2))
		if len(dMap[distance[i]]) == 0 {
			dMap[distance[i]] = make([][]int, 0)
		}
		dMap[distance[i]] = append(dMap[distance[i]], points[i])
	}
	sort.Slice(distance, func(i, j int) bool {
		return distance[i] > distance[j]
	})
	result := [][]int{}
	for i := len(distance) - 1; i >= 0 && k != 0; i-- {
		length := func(a, b int) int {
			if a < b {
				return a
			} else {
				return b
			}
		}(len(dMap[distance[i]]), k)
		for j := 0; j < length; j++ {
			result = append(result, dMap[distance[i]][j])
		}
		k -= length
	}
	fmt.Println(result)
	return result
}

func Zzz() {

}
