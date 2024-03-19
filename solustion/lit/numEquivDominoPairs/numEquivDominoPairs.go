package numEquivDominoPairs

import (
	"fmt"
	"strconv"
)

func NumEquivDominoPairs(dominoes [][]int) int {
	sMap, result := make(map[string]int), 0
	for i := 0; i < len(dominoes); i++ {
		aTostr, Tostr := fmt.Sprintln(strconv.Itoa(dominoes[i][1]), ",", strconv.Itoa(dominoes[i][0])), fmt.Sprintln(strconv.Itoa(dominoes[i][0]), ",", strconv.Itoa(dominoes[i][1]))
		if aTostr == Tostr {
			if sMap[Tostr] >= 1 {
				result += sMap[Tostr]
				fmt.Println(Tostr, dominoes[i], sMap[Tostr], i)
			}
			sMap[Tostr]++
		} else {
			if sMap[Tostr] >= 1 {
				result += sMap[Tostr]
				fmt.Println(Tostr, dominoes[i], sMap[Tostr], i)
			}
			sMap[Tostr]++
			if sMap[aTostr] >= 1 {
				fmt.Println(aTostr, dominoes[i], sMap[aTostr], i)
				result += sMap[aTostr]
			}
		}

	}
	return result
}
