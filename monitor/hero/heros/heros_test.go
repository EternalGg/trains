package heros

import (
	"fmt"
	"testing"
)

func TestChoseHeroByName(t *testing.T) {
	//fmt.Println(ChoseHeroByName("狂战士"))
	heros := SelectAllHero()
	for i := 0; i < len(heros); i++ {
		fmt.Println(heros[i])
	}
}
