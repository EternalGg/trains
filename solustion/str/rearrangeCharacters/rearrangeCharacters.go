package rearrangeCharacters

import "fmt"

func RearrangeCharacters(s string, target string) int {
	result, checkMap, iMap := 0, make(map[int32]int), make(map[int32]int)
	for _, value := range target {
		checkMap[value]++
	}

	for _, value := range s {
		if checkMap[value] > 0 {
			iMap[value]++
		}
	}
	result = iMap[int32(target[0])]
	for _, value := range target {
		fmt.Println(checkMap[value], iMap[value])
		if (checkMap[value] / iMap[value]) <= result {
			result = iMap[value]
		}
	}
	return result
}
