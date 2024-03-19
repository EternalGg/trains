package areAlmostEqual

func areAlmostEqual(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}
	if len(s2) != len(s1) {
		return false
	}
	s1Map, s2Map, s1List, diff := make(map[byte]int), make(map[byte]int), make([]byte, 0), 0
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			diff++
		}
		if s1Map[s1[i]] == 0 {
			s1List = append(s1List, s1[i])
		}
		s1Map[s1[i]]++
		s2Map[s2[i]]++
	}

	equal := true
	for _, b := range s1List {
		if s1Map[b] != s2Map[b] {
			equal = false
			break
		}
	}
	if diff == 2 && equal {
		return true
	} else {
		return false
	}
}
