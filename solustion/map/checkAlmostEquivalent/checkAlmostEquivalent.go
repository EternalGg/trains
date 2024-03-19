package checkAlmostEquivalent

import "math"

func checkAlmostEquivalent(word1 string, word2 string) bool {
	w1Map, w2Map := make(map[uint8]int), make(map[uint8]int)
	for i := 0; i < len(word1); i++ {
		w1Map[word1[i]]++
	}
	for i := 0; i < len(word2); i++ {
		w2Map[word2[i]]++
	}
	for i := 97; i < 26; i++ {
		if math.Abs(float64(w1Map[uint8(i)]-w2Map[uint8(i)])) >= 3 {
			return false
		}
	}
	return true
}
