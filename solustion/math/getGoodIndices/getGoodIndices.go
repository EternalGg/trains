package getGoodIndices

import (
	"fmt"
	"math/big"
)

func GetGoodIndices(variables [][]int, target int) []int {
	result := []int{}
	for i := 0; i < len(variables); i++ {

		ab := new(big.Int).Exp(big.NewInt(int64(variables[i][0])), big.NewInt(int64(variables[i][1])), nil)
		ten := new(big.Int).SetInt64(10)
		remainder := ab.Mod(ab, ten)
		c := big.NewInt(int64(variables[i][2]))
		abc := new(big.Int).Exp(remainder, c, nil)
		abcd := abc.Mod(abc, big.NewInt(int64(variables[i][3])))
		fmt.Println(abcd.Int64() == int64(target))
		if abcd.Int64() == int64(target) {
			result = append(result, i)
		}
	}
	return result
}
