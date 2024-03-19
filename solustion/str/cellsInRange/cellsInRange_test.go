package cellsInRange

import (
	"fmt"
	"testing"
)

func cellsInRange(s string) []string {
	qa, nHead, nLast, sHead, sLast := make([]string, 0), s[1]-48, s[4]-48, s[0]-65, s[3]-65

	for i := int(sHead); i <= int(sLast); i++ {
		for j := nHead; j <= (nLast); j++ {
			now := fmt.Sprint(string(uint8(i) + uint8(65)))
			now = fmt.Sprint(now, j)
			qa = append(qa, now)
		}
	}
	return qa
}

func TestCellsInRange(t *testing.T) {
	fmt.Println(
		cellsInRange("K1:L2"))
}
