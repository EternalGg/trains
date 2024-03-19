package numberOfBeams

import "strconv"

func numberOfBeams(bank []string) int {
	result, nlist := 0, make([]int, 0)
	for i := 0; i < len(bank); i++ {
		str, _ := strconv.ParseInt(bank[i], 10, 32)
		if str == 0 {
			continue
		} else {
			nlist = append(nlist, 0)
			length := len(nlist) - 1
			for _, i3 := range bank[i] {
				if i3 == '1' {
					nlist[length]++
				}
			}
		}
	}

	if len(nlist) >= 1 {
		for i := 0; i < len(nlist)-1; i++ {
			result += nlist[i] * nlist[i+1]
		}
	}
	return result
}
