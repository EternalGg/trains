package uniqueOccurrences

import "fmt"

func UniqueOccurrences(arr []int) bool {
	numList, tempoMap, uniqueMap := make([]int, 0), make(map[int]int), make(map[int]bool)
	for _, value := range arr {
		if tempoMap[value] > 0 {
			numList = append(numList, value)
		}
		tempoMap[value]++
	}
	for _, value := range numList {
		if !uniqueMap[tempoMap[value]] {
			fmt.Println(tempoMap[value])
			uniqueMap[tempoMap[value]] = true
		} else {
			fmt.Println(tempoMap[value])
			return false
		}
	}
	return true
}
