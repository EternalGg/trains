package checkDistances

import "fmt"

type distances struct {
	first int
}

func CheckDistances(s string, distance []int) bool {
	byteMap := make(map[byte]*distances)
	for i := 0; i < len(s); i++ {
		if byteMap[s[i]] == nil {
			d := new(distances)
			d.first = i
			byteMap[s[i]] = d
		} else {
			fmt.Println(i-byteMap[s[i]].first-1, distance[(s[i])-97])
			if i-byteMap[s[i]].first-1 != distance[(s[i])-97] {
				return false
			}
		}
	}
	return true
}
