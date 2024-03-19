package truncateSentence

import "fmt"

func TruncateSentence(s string, k int) string {
	strList, result, last := make([]string, 0), "", ""
	for _, value := range s {
		if value != 32 {
			last += string(value)
		} else {
			strList = append(strList, last)
			last = ""
		}
	}
	if last != "" {
		strList = append(strList, last)
	}
	fmt.Println(strList)
	for i := 0; i < k-2; i++ {
		result += strList[i] + " "
	}
	result += strList[k-1]
	return result
}
