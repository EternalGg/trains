package adventureCamp

import (
	"fmt"
	"testing"
)

func TestAdventureCamp(t *testing.T) {
	list := []string{"", "Gryffindor->Slytherin->Gryffindor", "Hogwarts->Hufflepuff->Ravenclaw"}
	fmt.Println(AdventureCamp(list))
}
