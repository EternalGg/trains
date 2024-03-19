package getGoodIndices

import (
	"fmt"
	"math/big"
	"testing"
)

func TestGetGoodIndices(t *testing.T) {
	ab := new(big.Int).Exp(big.NewInt(2), big.NewInt(4), nil)
	ten := new(big.Int).SetInt64(10)
	//c := new(big.Int).SetInt64()

	quotient := ab.Mod(ab, ten)
	fmt.Println(quotient)
}
