package subdomainVisits

import (
	"strconv"
)

func SubdomainVisits(cpdomains []string) []string {
	strMap, strList := make(map[string]int), make([]string, 0)
	for i := 0; i < len(cpdomains); i++ {
		time, last, end := int64(0), "", 0
		for j := 0; j < len(cpdomains[i]); j++ {
			switch cpdomains[i][j] {
			case ' ':
				time, _ = strconv.ParseInt(last, 10, 64)
				last = ""
				end = j
				goto here
			default:
				last += string(cpdomains[i][j])
			}
		}
	here:
		ad := 0
		for j := len(cpdomains[i]) - 1; j > end; j-- {
			switch cpdomains[i][j] {
			case '.':
				if strMap[last] == 0 {
					strList = append(strList, last)
				}
				strMap[last] += int(time)
				last = string(cpdomains[i][j]) + last
				ad++

			default:
				last = string(cpdomains[i][j]) + last
			}
		}
		if last != "" {
			if strMap[last] == 0 {
				strList = append(strList, last)
			}
			strMap[last] += int(time)
			last = ""
		}
	}
	for i := 0; i < len(strList); i++ {
		strList[i] = strconv.Itoa(strMap[strList[i]]) + " " + strList[i]
	}
	return strList
}
