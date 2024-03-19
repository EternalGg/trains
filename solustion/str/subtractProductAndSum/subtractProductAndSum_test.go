package subtractProductAndSum

import (
	"strconv"
	"testing"
)

func subtractProductAndSum(n int) int {
	nToByte, mult, plus := []byte(strconv.Itoa(n)), 1, 0
	for i := 0; i < len(nToByte); i++ {
		now := int(nToByte[i] - 48)
		mult *= now
		plus += now
	}
	if plus > mult {
		return plus - mult
	} else {
		return mult - plus
	}
}

func TestSubtractProductAndSum(t *testing.T) {

}
