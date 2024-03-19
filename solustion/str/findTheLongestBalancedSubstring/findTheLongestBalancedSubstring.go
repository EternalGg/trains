package findTheLongestBalancedSubstring

func findTheLongestBalancedSubstring(s string) int {
	result := 0
	zero, one, add := false, false, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			// 如果当前为0，并开启了1
			if one {
				zero = true
				one = false
				add = 1
			} else {
				// 如果当前为0，并且1没有开启
				add++
				zero = true
			}
		} else {
			// 如果有过0
			if zero != false {
				add++
				one = true
				if add > result {
					result = add
				}
			}
		}
	}
	return result
}
