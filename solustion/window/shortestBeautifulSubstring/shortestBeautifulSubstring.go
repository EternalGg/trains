package shortestBeautifulSubstring

import (
	"sort"
)

func ShortestBeautifulSubstring(s string, k int) string {

	for s[0] != '1' {

		s = s[1:]
		if len(s) == 0 {
			return ""
		}
	}

	for s[len(s)-1] != '1' {

		s = s[:len(s)-1]
		if len(s) == 0 {
			return ""
		}
	}
	if s == "" {
		return ""
	}

	leftpointer, rightpointer := 0, 0
	minlen, minlist := len(s), []string{}
	kcount := 0
	for leftpointer != len(s) {

		if rightpointer == len(s) {

			if kcount == k {
				if rightpointer-leftpointer <= minlen {
					if rightpointer-leftpointer < minlen {
						minlen = rightpointer - leftpointer
						minlist = []string{}
					}
					minlist = append(minlist, s[leftpointer:rightpointer])
				}
			}
			break
		}
		if kcount == k {
			if rightpointer-leftpointer <= minlen {
				if rightpointer-leftpointer < minlen {
					minlen = rightpointer - leftpointer
					minlist = []string{}
				}
				minlist = append(minlist, s[leftpointer:rightpointer])
			}

			for i := leftpointer; i < rightpointer; i++ {
				if s[leftpointer] == '1' {
					for j := i + 1; j < rightpointer; j++ {
						if s[j] == '1' {
							leftpointer = j
							goto here
						}
					}
				}
			}
		here:
			kcount--

		} else {
			if s[rightpointer] == '1' {
				kcount++
			}
			rightpointer++
		}
	}
	sort.Slice(minlist, func(i, j int) bool {
		return minlist[i] > minlist[j]
	})
	if len(minlist) == 0 {
		return ""
	} else {
		return minlist[0]
	}
}
