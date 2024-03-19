package canPermutePalindrome

import "fmt"

func canPermutePalindrome(s string) bool {
	m := map[uint8]int{}
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}
	if len(s)%2 == 0 {
		for _, i := range m {
			if i%2 != 0 {
				return false
			}
		}
	} else {
		odd := 0
		for _, i := range m {
			if i%2 != 0 {
				odd++
				if odd == 2 {
					fmt.Println(i)
					return false
				}
			}
		}
	}
	return true
}
