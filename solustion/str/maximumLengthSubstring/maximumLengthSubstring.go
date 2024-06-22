package maximumLengthSubstring

func MaximumLengthSubstring(s string) int {
	max, left := 0, 0
	for !(max >= len(s)-left) {
		m := make(map[uint8]int)
		for i := left; i < len(s); i++ {
			if m[s[i]] == 0 {
				m[s[i]] = i + 1
			} else {
				if max < i-left {
					max = i - left
					left = m[s[i]]
					break
				}
			}
		}
	}

	return max
}
