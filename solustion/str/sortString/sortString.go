package sortString

func PrefixCount(words []string, pref string) int {

	count := 0

	for _, word := range words {
		if len(word) < len(pref) {
			continue
		}
		for i2, s2 := range pref {
			if len(word) == i2 {
				break
			}
			if s2 != int32(word[i2]) {
				goto here
			}
		}
		count++
		continue
	here:
	}
	return count
}
