package compressString

import "strconv"

func compressString(S string) string {
	result, last, times := "", S[0], 0
	for i := 0; i < len(S); i++ {
		if S[i] == last {
			times++
		} else {
			result += string(last) + strconv.Itoa(times)
			last = S[i]
			times = 1
		}
	}
	result += string(last) + strconv.Itoa(times)
	if len(result) > len(S) {
		return S
	} else {
		return result
	}

}
