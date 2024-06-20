package battlefiled

import (
	"fmt"
	"testing"
)

func TestBattleFiled_NormalGame(t *testing.T) {
	//bf := BattleFiled{Positions: map[uint]*Position{}}
	bf := NormalGame()
	fmt.Println(bf.Positions[20])
}

func BenchmarkName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NormalGame()

	}
}
