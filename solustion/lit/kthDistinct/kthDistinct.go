package kthDistinct

func kthDistinct(arr []string, k int) string {
	strList, strMap := make([]string, 0), make(map[string]int)
	for _, str := range arr {
		strMap[str]++
	}
	for _, str := range arr {
		if strMap[str] == 1 {
			strList = append(strList, str)
		}
	}
	if len(strList) >= k {
		return strList[k]
	}
	return ""
}
