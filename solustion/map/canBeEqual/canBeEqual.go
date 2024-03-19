package canBeEqual

func canBeEqual(s1 string, s2 string) bool {
	s1Map, s2Map := make(map[uint8]int), make(map[uint8]int)
	for i := 0; i < 4; i++ {
		s1Map[s1[i]]++
		s2Map[s2[i]]++
	}
	s12, s13 := []byte(s1), []byte(s1)
	s12[0] = s1[2]
	s12[2] = s1[0]
	s13[1] = s1[3]
	s13[3] = s1[1]

	s22, s23 := []byte(s2), []byte(s2)
	s22[0] = s2[2]
	s22[2] = s2[0]
	s23[1] = s2[3]
	s23[3] = s2[1]

	if s1 == s2 || s1 == string(s22) || s1 == string(s23) || string(s12) == s2 || string(s12) == string(s22) || string(s12) == string(s23) || string(s13) == s2 || string(s13) == string(s22) || string(s13) == string(s23) {
		return true
	}
	return false
}
