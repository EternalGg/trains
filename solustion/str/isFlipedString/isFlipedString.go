package isFlipedString

import "strings"

func isFlipedString(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	s3 := s2 + s2
	if strings.Contains(s3, s1) {
		return true
	}
	return false
}
