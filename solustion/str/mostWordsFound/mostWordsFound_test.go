package mostWordsFound

import "testing"

func mostWordsFound(sentences []string) int {
	max := 0

	for _, sentence := range sentences {
		now := 0
		for _, str := range sentence {
			if str == 32 {
				now++
			}
		}
		if now > 0 {
			now += 2
		}
		if max < now {
			max = now
		}
	}
	return max
}

func TestMostWordsFound(t *testing.T) {

}
