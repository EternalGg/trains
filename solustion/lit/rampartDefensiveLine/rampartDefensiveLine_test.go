package rampartDefensiveLine

import (
	"fmt"
	"testing"
)

func TestRampartDefensiveLine(t *testing.T) {
	list := make([][]int, 4)
	//list[0] = []int{0, 3}
	//list[1] = []int{4, 5}
	//list[2] = []int{7, 9}
	list[0] = []int{1, 2}
	list[1] = []int{5, 8}
	list[2] = []int{11, 15}
	list[3] = []int{18, 25}
	fmt.Println(RampartDefensiveLine(list))
}
