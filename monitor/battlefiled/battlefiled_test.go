package battlefiled

import (
	"fmt"
	"testing"
)

func TestBattleFiled_NormalGame(t *testing.T) {
	bf := NormalGame()
	list := []*Position{}
	for i := 0; i < 45; i++ {
		if bf.Positions[20].Distance[i] <= 1 && bf.Positions[i] != nil {
			list = append(list, bf.Positions[i])
			fmt.Println(i)
		}
	}

}

func BenchmarkName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NormalGame()

	}
}
