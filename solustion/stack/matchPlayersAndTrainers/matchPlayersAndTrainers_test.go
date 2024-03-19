package matchPlayersAndTrainers

import (
	"testing"
)

func TestMatchPlayersAndTrainers(t *testing.T) {
	p := []int{4, 7, 9}
	tr := []int{8, 2, 5, 9}
	MatchPlayersAndTrainers(p, tr)
}
