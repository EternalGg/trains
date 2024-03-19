package numDifferentIntegers

import "strconv"

func NumDifferentIntegers(word string) int {
	numMap, last, count := make(map[string]bool), "", 0

	for _, value := range word {
		if value >= 48 && value <= 57 {
			num, _ := strconv.ParseInt(last, 10, 64)
			if num == 0 && value != 48 {
				last = string(value)
				continue
			}
			last += string(value)
		} else {
			if !(last == "") {
				if !numMap[last] {
					count++
					numMap[last] = true
				}
				last = ""
			}
		}
	}

	if !(last == "") {
		if !numMap[last] {
			count++
			numMap[last] = true
		}
		last = ""
	}
	return count
}
